package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// FileInfo represents information about a file or directory
type FileInfo struct {
	Name    string
	Size    int64
	IsDir   bool
	ModTime int64
}

// FileChunk represents a chunk of a file
type FileChunk struct {
	ChunkIndex int
	Data       []byte
}

// FileSyncNode represents a node in the file synchronization system
type FileSyncNode struct {
	ID       int
	Address  string
	RootPath string
	mutex    sync.Mutex
}

// FileSyncSystem represents the file synchronization system
type FileSyncSystem struct {
	Nodes []*FileSyncNode
}

// NewFileSyncNode creates a new instance of a file synchronization node
func NewFileSyncNode(id int, address, rootPath string) *FileSyncNode {
	return &FileSyncNode{
		ID:       id,
		Address:  address,
		RootPath: rootPath,
	}
}

// NewFileSyncSystem creates a new instance of the file synchronization system
func NewFileSyncSystem() *FileSyncSystem {
	return &FileSyncSystem{
		Nodes: []*FileSyncNode{},
	}
}

// AddNode adds a new file synchronization node to the system
func (fs *FileSyncSystem) AddNode(address, rootPath string) {
	node := NewFileSyncNode(len(fs.Nodes), address, rootPath)
	fs.Nodes = append(fs.Nodes, node)
}

// StartServer starts a file synchronization server to handle incoming requests
func StartServer(node *FileSyncNode, fs *FileSyncSystem) {
	listener, err := net.Listen("tcp", node.Address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("FileSync Server %d started. Listening on %s\n", node.ID, node.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go HandleFileSyncRequest(conn, node, fs)
	}
}

// HandleFileSyncRequest handles incoming file synchronization requests
func HandleFileSyncRequest(conn net.Conn, node *FileSyncNode, fs *FileSyncSystem) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)

	var request string
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println("Error decoding request:", err)
		return
	}

	switch request {
	case "LIST":
		fileList := GetFileList(node.RootPath)
		encoder := gob.NewEncoder(conn)
		encoder.Encode(fileList)
	case "SYNC":
		var fileInfo FileInfo
		err := decoder.Decode(&fileInfo)
		if err != nil {
			fmt.Println("Error decoding file info:", err)
			return
		}

		node.mutex.Lock()
		defer node.mutex.Unlock()

		filePath := filepath.Join(node.RootPath, fileInfo.Name)

		if fileInfo.IsDir {
			// Create directory if it doesn't exist
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				err := os.MkdirAll(filePath, os.ModePerm)
				if err != nil {
					fmt.Println("Error creating directory:", err)
					return
				}
			}
		} else {
			// Receive file chunks and write to file
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			for {
				var fileChunk FileChunk
				err := decoder.Decode(&fileChunk)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Error decoding file chunk:", err)
					return
				}

				file.WriteAt(fileChunk.Data, int64(fileChunk.ChunkIndex*len(fileChunk.Data)))
			}
		}
	}
}

// GetFileList returns a list of files and directories in the given path
func GetFileList(path string) []FileInfo {
	var fileList []FileInfo

	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return err
		}

		relPath, _ := filepath.Rel(path, filePath)
		fileList = append(fileList, FileInfo{
			Name:    relPath,
			Size:    info.Size(),
			IsDir:   info.IsDir(),
			ModTime: info.ModTime().Unix(),
		})

		return nil
	})

	return fileList
}

// SyncFiles synchronizes files between nodes
func SyncFiles(sourceNode, targetNode *FileSyncNode, fileInfo FileInfo) {
	sourcePath := filepath.Join(sourceNode.RootPath, fileInfo.Name)
	targetPath := filepath.Join(targetNode.RootPath, fileInfo.Name)

	if fileInfo.IsDir {
		// Create directory if it doesn't exist
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			err := os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}
	} else {
		// Read file and send chunks to target node
		file, err := os.Open(sourcePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		conn, err := net.Dial("tcp", targetNode.Address)
		if err != nil {
			fmt.Println("Error connecting to target node:", err)
			return
		}
		defer conn.Close()

		encoder := gob.NewEncoder(conn)
		encoder.Encode("SYNC")
		encoder.Encode(fileInfo)

		reader := bufio.NewReader(file)
		buffer := make([]byte, 1024)
		chunkIndex := 0

		for {
			bytesRead, err := reader.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			chunk := FileChunk{
				ChunkIndex: chunkIndex,
				Data:       buffer[:bytesRead],
			}

			encoder.Encode(chunk)
			chunkIndex++
		}
	}
}

// WriteSampleFile writes a sample file for testing
func WriteSampleFile(filePath string) {
	fileContent := []byte("This is a sample file content.")
	err := ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing sample file:", err)
	}
}

func main() {
	fs := NewFileSyncSystem()

	// Add nodes to the file synchronization system
	fs.AddNode("localhost:6000", "node1")
	fs.AddNode("localhost:6001", "node2")

	// Start file synchronization servers
	go StartServer(fs.Nodes[0], fs)
	go StartServer(fs.Nodes[1], fs)

	// Simulate a file change on one node
	filePath := filepath.Join(fs.Nodes[0].RootPath, "sample.txt")
	WriteSampleFile(filePath)

	fmt.Println("Waiting for servers to start...")
	time.Sleep(2 * time.Second)

	// Simulate a file change on one node
	filePath := filepath.Join(fs.Nodes[0].RootPath, "sample.txt")
	WriteSampleFile(filePath)

	// Wait for a while to allow synchronization
	time.Sleep(2 * time.Second)

	// Print file lists from each node
	fmt.Printf("Node 0 File List: %v\n", GetFileList(fs.Nodes[0].RootPath))
	fmt.Printf("Node 1 File List: %v\n", GetFileList(fs.Nodes[1].RootPath))
}
