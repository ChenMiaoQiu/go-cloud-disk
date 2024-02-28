package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderCreateService struct {
	ParentFolderID string `json:"parent" form:"parent" binding:"required"`
	FileFolderName string `json:"name" form:"name" binding:"required"`
	FileStoreId    string `json:"store" form:"store" binding:"required"`
}

func (service *FileFolderCreateService) CreateFileFolder(userId string) serializer.Response {
	FilerFolder := model.FileFolder{
		FileFolderName: service.FileFolderName,
		ParentFolderID: service.ParentFolderID,
		FileStoreID:    service.FileStoreId,
		OwnerID:        userId,
	}
	if err := model.DB.Create(&FilerFolder).Error; err != nil {
		return serializer.DBErr("create filefolder err when filefolercreate", err)
	}
	return serializer.Success(nil)
}
