package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderCreateService struct {
	ParentFolderID string `json:"parent" form:"parent" binding:"required"`
	FileFolderName string `json:"name" form:"name" binding:"required"`
}

// CreateFileFolder create filefolder to user database
func (service *FileFolderCreateService) CreateFileFolder(userId string) serializer.Response {
	// check if user match
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.ParentFolderID).Find(&fileFolder).Error; err != nil {
		return serializer.DBErr("dberr", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("can't matched user when create filefolder")
	}

	// insert filefolder to database
	createFilerFolder := model.FileFolder{
		FileFolderName: service.FileFolderName,
		ParentFolderID: service.ParentFolderID,
		FileStoreID:    fileFolder.FileStoreID,
		OwnerID:        userId,
		Size:           0,
	}

	if err := model.DB.Create(&createFilerFolder).Error; err != nil {
		return serializer.DBErr("create filefolder err when filefolercreate", err)
	}
	return serializer.Success(serializer.BuildFileFolder(createFilerFolder))
}
