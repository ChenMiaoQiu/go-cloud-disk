package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileFolder struct {
	Uuid           string `gorm:"primarykey"`
	FileFolderName string
	ParentFolderID string
	FileStoreID    string
	OwnerID        string
}

// BeforeCreate create uuid before insert database
func (fileFolder *FileFolder) BeforeCreate(tx *gorm.DB) (err error) {
	if fileFolder.Uuid == "" {
		fileFolder.Uuid = uuid.New().String()
	}
	return
}

// CreateBaseFileFolder create a user fileFolder with fileStoreId and ownerId,
// and return it uuid or err
func CreateBaseFileFolder(ownerId string, fileStoreId string) (string, error) {
	fileStore := FileFolder{
		FileFolderName: "main",
		ParentFolderID: "root",
		FileStoreID:    fileStoreId,
		OwnerID:        ownerId,
	}
	if err := DB.Create(&fileStore).Error; err != nil {
		return "", err
	}
	return fileStore.Uuid, nil
}
