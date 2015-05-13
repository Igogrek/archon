/*
* Archon Patch Server
* Copyright (C) 2014 Andrew Rodman
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
* ---------------------------------------------------------------------
 */
package patch_server

import (
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"libarchon/encryption"
	"libarchon/logger"
	"libarchon/util"
	"net"
	"os"
	"runtime/debug"
	"strings"
	"sync"
)

var log *logger.Logger

type pktHandler func(p *PatchClient) error

// Struct for holding client-specific data.
type PatchClient struct {
	conn   *net.TCPConn
	ipAddr string

	clientCrypt *encryption.PSOCrypt
	serverCrypt *encryption.PSOCrypt
	recvData    []byte
	recvSize    int
	packetSize  uint16
}

func (lc PatchClient) Connection() *net.TCPConn { return lc.conn }
func (lc PatchClient) IPAddr() string           { return lc.ipAddr }

var patchConnections *util.ConnectionList = util.NewClientList()

// Data for one patch file.
type PatchEntry struct {
	filename string
	checksum uint32
	fileSize uint32
	index    uint32
	dirSteps byte
	pathDirs []string
}

// Basic tree structure for holding patch data that more closely represents
// a file hierarchy and makes it easier to handle the client working dir.
// Patch files and subdirectories are represented as lists in order to make
// a breadth-first search easier and the order predictable.
type PatchDir struct {
	dirname string
	patches []*PatchEntry
	subdirs []*PatchDir
}

// File names that should be ignored when searching for patch files. This
// could also be an array but the map makes it quicker to compare.
var SkipPaths = map[string]byte{".": 1, "..": 1, ".DS_Store": 1, ".rid": 1}

// Each index corresponds to a patch file. This is constructed in the order
// that the patch tree will be traversed and makes it faster to locate a
// patch entry when the client sends us an index in the FileStatusPacket.
var patchTree PatchDir
var patchIndex []*PatchEntry

// Create and initialize a new struct to hold client information.
func newClient(conn *net.TCPConn) (*PatchClient, error) {
	client := new(PatchClient)
	client.conn = conn
	client.ipAddr = conn.RemoteAddr().String()

	client.clientCrypt = encryption.NewCrypt()
	client.serverCrypt = encryption.NewCrypt()
	client.clientCrypt.CreateKeys()
	client.serverCrypt.CreateKeys()
	client.recvData = make([]byte, 2048)

	var err error = nil
	if SendWelcome(client) != 0 {
		err = errors.New("Error sending welcome packet to: " + client.ipAddr)
		client = nil
	}
	return client, err
}

// Traverse the patch tree depth-first and send the check file requests.
func sendFileList(client *PatchClient, node *PatchDir) {
	// Step into the next directory.
	SendChangeDir(client, node.dirname)
	for _, subdir := range node.subdirs {
		sendFileList(client, subdir)
	}
	fmt.Printf("Dir: %s\n", node.dirname)
	for _, patch := range node.patches {
		fmt.Printf("File: %s\n", patch.filename)
		SendCheckFile(client, patch.index, patch.filename)
	}
	// Move them back up each time we leave a directory.
	SendDirAbove(client)
}

// The client sent us a checksum for one of the patch files. Compare it
// to what we have and add it to the list of files to update if there
// is any discrepancy.
func handleFileStatus(client *PatchClient) {
	var fileStatus FileStatusPacket
	util.StructFromBytes(client.recvData[:], &fileStatus)

	patch := patchIndex[fileStatus.PatchId]
	fmt.Printf("Checking file %s\n", patch.filename)
	if fileStatus.Checksum != patch.checksum || fileStatus.FileSize != patch.fileSize {
		fmt.Printf("Needs update\n")
	}
}

// Handle a packet sent to the PATCH server.
func processPatchPacket(client *PatchClient) error {
	var pktHeader PCPktHeader
	util.StructFromBytes(client.recvData[:BBHeaderSize], &pktHeader)

	if GetConfig().DebugMode {
		fmt.Printf("Got %v bytes from client:\n", pktHeader.Size)
		util.PrintPayload(client.recvData, int(pktHeader.Size))
		fmt.Println()
	}
	var err error = nil
	switch pktHeader.Type {
	case WelcomeType:
		SendWelcomeAck(client)
	case LoginType:
		cfg := GetConfig()
		if SendWelcomeMessage(client) == 0 {
			SendRedirect(client, cfg.RedirectPort(), cfg.HostnameBytes())
		}
	default:
		msg := fmt.Sprintf("Received unknown packet %2x from %s", pktHeader.Type, client.ipAddr)
		log.Info(msg, logger.LogPriorityMedium)
	}
	return err
}

// Handle a packet sent to the DATA server.
func processDataPacket(client *PatchClient) error {
	var pktHeader PCPktHeader
	util.StructFromBytes(client.recvData[:BBHeaderSize], &pktHeader)

	if GetConfig().DebugMode {
		fmt.Printf("Got %v bytes from client:\n", pktHeader.Size)
		util.PrintPayload(client.recvData, int(pktHeader.Size))
		fmt.Println()
	}
	var err error = nil
	switch pktHeader.Type {
	case WelcomeType:
		SendWelcomeAck(client)
	case LoginType:
		SendDataAck(client)
		sendFileList(client, &patchTree)
		SendFileListDone(client)
	case FileStatusType:
		handleFileStatus(client)
	default:
		msg := fmt.Sprintf("Received unknown packet %02x from %s", pktHeader.Type, client.ipAddr)
		log.Info(msg, logger.LogPriorityMedium)
	}
	return err
}

// Handle communication with a particular client until the connection is
// closed or an error is encountered.
func handleClient(client *PatchClient, desc string, handler pktHandler) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := fmt.Sprintf("Error in client communication: %s: %s\n%s\n",
				client.ipAddr, err, debug.Stack())
			log.Error(errMsg, logger.LogPriorityCritical)
		}
		client.conn.Close()
		patchConnections.RemoveClient(client)
		log.Info("Disconnected "+desc+" client "+client.ipAddr, logger.LogPriorityMedium)
	}()

	log.Info("Accepted "+desc+" connection from "+client.ipAddr, logger.LogPriorityMedium)
	// We're running inside a goroutine at this point, so we can block on this connection
	// and not interfere with any other clients.
	for {
		// Wait for the packet header.
		for client.recvSize < BBHeaderSize {
			bytes, err := client.conn.Read(client.recvData[client.recvSize:])
			if bytes == 0 || err == io.EOF {
				// The client disconnected, we're done.
				client.conn.Close()
				return
			} else if err != nil {
				// Socket error, nothing we can do now
				log.Warn("Socket Error ("+client.ipAddr+") "+err.Error(),
					logger.LogPriorityMedium)
				return
			}

			client.recvSize += bytes
			if client.recvSize >= BBHeaderSize {
				// We have our header; decrypt it.
				client.clientCrypt.Decrypt(client.recvData[:BBHeaderSize], BBHeaderSize)
				client.packetSize, err = util.GetPacketSize(client.recvData[:2])
				if err != nil {
					// Something is seriously wrong if this causes an error. Bail.
					panic(err.Error())
				}
			}
		}

		// Wait until we have the entire packet.
		for client.recvSize < int(client.packetSize) {
			bytes, err := client.conn.Read(client.recvData[client.recvSize:])
			if err != nil {
				log.Warn("Socket Error ("+client.ipAddr+") "+err.Error(),
					logger.LogPriorityMedium)
				return
			}
			client.recvSize += bytes
		}

		// We have the whole thing; decrypt the rest of it if needed and pass it along.
		if client.packetSize > BBHeaderSize {
			client.clientCrypt.Decrypt(
				client.recvData[BBHeaderSize:client.packetSize],
				uint32(client.packetSize-BBHeaderSize))
		}
		if err := handler(client); err != nil {
			log.Info(err.Error(), logger.LogPriorityLow)
			break
		}

		// Alternatively, we could set the slice to to nil here and make() a new one in order
		// to allow the garbage collector to handle cleanup, but I expect that would have a
		// noticable impact on performance. Instead, we're going to clear it manually.
		util.ZeroSlice(client.recvData, client.recvSize)
		client.recvSize = 0
		client.packetSize = 0
	}
}

// Main worker for the patch server. Creates the socket and starts listening for connections,
// spawning off client threads to handle communications for each client.
func startPatch(wg *sync.WaitGroup) {
	patchConfig := GetConfig()
	socket, err := util.OpenSocket(patchConfig.Hostname, patchConfig.PatchPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Waiting for PATCH connections on %s:%s...\n",
		patchConfig.Hostname, patchConfig.PatchPort)
	for {
		connection, err := socket.AcceptTCP()
		if err != nil {
			log.Error("Failed to accept connection: "+err.Error(), logger.LogPriorityHigh)
			continue
		}
		client, err := newClient(connection)
		if err != nil {
			continue
		}
		patchConnections.AddClient(client)
		go handleClient(client, "PATCH", processPatchPacket)
	}
	wg.Done()
}

// Main worker for the data server. Creates the socket and starts listening for connections,
// spawning off client threads to handle communications for each client.
func startData(wg *sync.WaitGroup) {
	patchConfig := GetConfig()
	socket, err := util.OpenSocket(patchConfig.Hostname, patchConfig.DataPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Waiting for DATA connections on %s:%s...\n",
		patchConfig.Hostname, patchConfig.DataPort)
	for {
		connection, err := socket.AcceptTCP()
		if err != nil {
			log.Error("Failed to accept connection: "+err.Error(), logger.LogPriorityHigh)
			continue
		}
		client, err := newClient(connection)
		if err != nil {
			continue
		}
		patchConnections.AddClient(client)
		go handleClient(client, "DATA", processDataPacket)
	}
	wg.Done()
}

// Recursively build the list of patch files present in the patch directory
// to sync with the client. Files are represented in a tree, directories act
// as nodes (PatchDir) and each keeps a list of patches/subdirectories.
func loadPatches(node *PatchDir, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Couldn't parse %s\n", path)
		return err
	}
	dirs := strings.Split(path, "/")
	node.dirname = dirs[len(dirs)-1]

	for _, file := range files {
		filename := file.Name()
		if _, ignore := SkipPaths[filename]; ignore {
			continue
		} else if file.IsDir() {
			subdir := new(PatchDir)
			node.subdirs = append(node.subdirs, subdir)
			loadPatches(subdir, path+"/"+filename)
		} else {
			data, err := ioutil.ReadFile(path + "/" + filename)
			if err != nil {
				return err
			}
			patch := new(PatchEntry)
			patch.filename = filename
			patch.fileSize = uint32(file.Size())
			patch.checksum = crc32.ChecksumIEEE(data)

			node.patches = append(node.patches, patch)
			fmt.Printf("%s (%d bytes, checksum: %v)\n",
				path+"/"+filename, patch.fileSize, patch.checksum)
		}
	}
	return nil
}

// Build the patch index, performing a depth-first search and mapping
// each patch entry to an array so that they're quickly indexable when
// we need to look up the patch data.
func buildPatchIndex(node *PatchDir) {
	for _, dir := range node.subdirs {
		buildPatchIndex(dir)
	}
	for _, patch := range node.patches {
		patchIndex = append(patchIndex, patch)
		patch.index = uint32(len(patchIndex) - 1)
	}
}

func StartServer() {
	fmt.Println("Initializing Archon PATCH and DATA servers...")

	// Initialize our config singleton from one of two expected file locations.
	config := GetConfig()
	fmt.Printf("Loading config file %v...", patchConfigFile)
	err := config.InitFromFile(patchConfigFile)
	if err != nil {
		os.Chdir(ServerConfigDir)
		fmt.Printf("Failed.\nLoading config from %v...", ServerConfigDir+"/"+patchConfigFile)
		err = config.InitFromFile(patchConfigFile)
		if err != nil {
			fmt.Println("Failed.\nPlease check that one of these files exists and restart the server.")
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
	}
	fmt.Printf("Done.\n\n--Configuration Parameters--\n%v\n\n", config.String())

	// Construct our patch tree from the specified directory.
	fmt.Printf("Loading patches from %s...\n", config.PatchDir)
	os.Chdir(config.PatchDir)
	if err := loadPatches(&patchTree, "."); err != nil {
		fmt.Printf("Failed to load patches: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println()
	buildPatchIndex(&patchTree)

	// Initialize the logger.
	log = logger.New(config.logWriter, config.LogLevel)
	log.Info("Server Initialized", logger.LogPriorityCritical)

	// Create a WaitGroup so that main won't exit until the server threads have exited.
	var wg sync.WaitGroup
	wg.Add(2)
	go startPatch(&wg)
	go startData(&wg)
	wg.Wait()
}
