package main

import (
	"decentralized-storage/internal/models"
	"decentralized-storage/internal/storage"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	// Input file path
	filePath := "exp.txt"
	chunkSize := 1024 * 1024 // 1 MB

	// chunk the file
	chunkIDs, err := storage.ChunkFile(filePath, chunkSize)
	if err != nil {
		fmt.Printf("Error chunking file: %v\n", err)
		return
	}

	// Get file info for metadata
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error generating file info: %v\n", err)
	}

	// create metadata
	metadata := models.Metadata{
		Filename:  fileInfo.Name(),
		Filesize:  fileInfo.Size(),
		ChunkIDs:  chunkIDs,
		Timestamp: time.Now().Unix(),
	}

	// Save the metadata to file
	metadataFile := fmt.Sprintf("storage/%s.metadata.json", metadata.Filename)
	file, err := os.Create(metadataFile)
	if err != nil {
		fmt.Printf("Error creating metadata file: %v\n", err)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(metadata)
	if err != nil {
		fmt.Printf("Error waiting metadata to file: %v\n", err)
		return
	}

	fmt.Printf("File uploaded successfully!\n")
	fmt.Printf("Chunks: %v\n", chunkIDs)
}
