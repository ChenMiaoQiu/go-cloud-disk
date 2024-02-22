package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// File_store serializer
type FileStore struct {
	Uuid        string `json:"id"`
	OwnerId     string `json:"owner"`
	MaxSize     uint   `json:"maxsize"`
	CurrentSize uint   `json:"currentsize"`
}

// BuildFileStore return fileStore serializer
func BuildFileStore(fileStore model.FileStore) FileStore {
	return FileStore{
		Uuid:        fileStore.Uuid,
		OwnerId:     fileStore.OwnerID,
		MaxSize:     fileStore.MaxSize,
		CurrentSize: fileStore.CurrentSize,
	}
}
