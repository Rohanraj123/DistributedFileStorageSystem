package models

// Metadata represents the metadata for a file
type Metadata struct {
	Filename  string   `json:"file_name"`
	Filesize  int64    `json:"file_size"`
	ChunkIDs  []string `json:"chunk_ids"`
	Timestamp int64    `json:"timestamp"`
}
