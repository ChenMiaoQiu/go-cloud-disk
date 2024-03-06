package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// File_store serializer
type FileStore struct {
	MaxSize     int64 `json:"maxsize"`
	CurrentSize int64 `json:"currentsize"`
}

// BuildFileStore return fileStore serializer
func BuildFileStore(fileStore model.FileStore) FileStore {
	return FileStore{
		MaxSize:     fileStore.MaxSize,
		CurrentSize: fileStore.CurrentSize,
	}
}
