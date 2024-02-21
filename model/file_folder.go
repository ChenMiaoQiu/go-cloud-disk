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
}

// BeforeCreate create uuid before insert database
func (fileFolder *FileFolder) BeforeCreate(tx *gorm.DB) (err error) {
	fileFolder.Uuid = uuid.New().String()
	return
}

func CreateBaseFileFolder(fileStoreId string) error {
	fileStore := FileFolder{
		FileFolderName: "main",
		ParentFolderID: "root",
		FileStoreID:    fileStoreId,
	}
	if err := DB.Create(&fileStore).Error; err != nil {
		return err
	}
	return nil
}
