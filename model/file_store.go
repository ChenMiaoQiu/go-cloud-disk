package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileStore struct {
	Uuid        string `gorm:"primarykey"`
	OwnerID     string
	CurrentSize uint
	MaxSize     uint
}

// BeforeCreate create uuid before insert database
func (fileStore *FileStore) BeforeCreate(tx *gorm.DB) (err error) {
	if fileStore.Uuid != "" {
		fileStore.Uuid = uuid.NewString()
	}
	return
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
