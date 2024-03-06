package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileFolder struct {
	Uuid           string `gorm:"primarykey"`
	FileFolderName string
	ParentFolderID string
	FileStoreID    string
	OwnerID        string
	Size           int64
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

// SubSize sub filefolder size
func (fileFolder *FileFolder) SubSize(size int64) error {
	fileFolder.Size = max(fileFolder.Size-size, 0)
	return nil
}

// AddFileFolderSize add filefolder size and add size for parent filefolder
func (fileFolder *FileFolder) AddFileFolderSize(appendSize int64) (err error) {
	// add size for filefolder
	fileFolder.Size += appendSize
	parentId := fileFolder.ParentFolderID
	if err := DB.Save(fileFolder).Error; err != nil {
		return fmt.Errorf("save filefolder err when sub filesize err %v", err)
	}

	// add size for parent filefolder
	for parentId != "root" && parentId != "" {
		var nowFileFolder FileFolder
		if err := DB.Where("uuid = ?", parentId).Find(&nowFileFolder).Error; err != nil {
			return fmt.Errorf("find filefolder err when add filesize %v", err)
		}
		if nowFileFolder.Uuid == "" {
			break
		}
		nowFileFolder.Size += appendSize
		if err := DB.Save(&nowFileFolder).Error; err != nil {
			return fmt.Errorf("save filefolder err when add filesize err %v", err)
		}
		parentId = nowFileFolder.ParentFolderID
	}

	return err
}

// SubFileFolderSize sub filefolder size and sub size for parent filefolder
func (fileFolder *FileFolder) SubFileFolderSize(size int64) (err error) {
	// sub size for filefolder
	fileFolder.SubSize(size)
	parentId := fileFolder.ParentFolderID
	if err := DB.Save(fileFolder).Error; err != nil {
		return fmt.Errorf("save filefolder err when sub filesize err %v", err)
	}

	// sub size for parent filefolder
	for parentId != "root" && parentId != "" {
		var nowFileFolder FileFolder
		if err := DB.Where("uuid = ?", parentId).Find(&nowFileFolder).Error; err != nil {
			return fmt.Errorf("find filefolder err when sub filesize %v", err)
		}
		if nowFileFolder.Uuid == "" {
			break
		}
		nowFileFolder.SubSize(size)
		if err := DB.Save(&nowFileFolder).Error; err != nil {
			return fmt.Errorf("save filefolder err when sub filesize err %v", err)
		}
		parentId = nowFileFolder.ParentFolderID
	}

	return err
}
