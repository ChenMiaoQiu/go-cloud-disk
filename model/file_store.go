package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileStore struct {
	Uuid        string `gorm:"primarykey"`
	OwnerID     string
	CurrentSize int64
	MaxSize     int64
}

// BeforeCreate create uuid before insert database
func (fileStore *FileStore) BeforeCreate(tx *gorm.DB) (err error) {
	if fileStore.Uuid == "" {
		fileStore.Uuid = uuid.NewString()
	}
	return
}

// AddCurrentSize add size to currentsize
func (fileStore *FileStore) AddCurrentSize(size int64) (err error) {
	if fileStore.CurrentSize+size > fileStore.MaxSize {
		return fmt.Errorf("add size exceed currentSize")
	}
	fileStore.CurrentSize += size
	return nil
}

// SubCurrentSize sub size to cueerentsize
func (fileStore *FileStore) SubCurrentSize(size int64) (err error) {
	fileStore.CurrentSize = max(fileStore.CurrentSize-size, 0)
	return nil
}

// CreateFileStore create new fileStore by userId, and return it uuid or err
func CreateFileStore(userId string) (string, error) {
	fileStore := FileStore{
		OwnerID:     userId,
		CurrentSize: 0,
		MaxSize:     1024 * 1024,
	}
	if err := DB.Create(&fileStore).Error; err != nil {
		return "", err
	}
	return fileStore.Uuid, nil
}
