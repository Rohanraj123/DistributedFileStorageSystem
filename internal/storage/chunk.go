package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// ChunkFile splits the file into chunks of specified size and returns their IDs.
func ChunkFile(filepath string, chunkSize int) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	chunkIDs := []string{}
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Generate a chunk ID using SHA-256 hash
		chunkID := fmt.Sprintf("%x", sha256.Sum256(buffer[:n]))
		chunkIDs = append(chunkIDs, chunkID)

		// Save the chunk to the disk
		err = os.WriteFile(fmt.Sprintf("storage/%s.chunk", chunkID), buffer[:n], 0644)
		if err != nil {
			return nil, err
		}
	}

	return chunkIDs, nil
}
