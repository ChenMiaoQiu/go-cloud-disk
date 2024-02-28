package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderGetAllFileFolderService struct {
	FileFolderID string `json:"id" form:"id" binding:"required"`
}

func (service *FileFolderGetAllFileFolderService) GetAllFileFolder(userId string) serializer.Response {
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.FileFolderID).Find(&fileFolder).Error; err != nil {
		return serializer.DBErr("search filefolder err when get all filefolder", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("can't matched user when get all filefolder")
	}

	var filefolder []model.FileFolder
	if err := model.DB.Where("parent_folder_id = ?", service.FileFolderID).Find(&filefolder).Error; err != nil {
		return serializer.DBErr("get all file db err", err)
	}
	return serializer.Success(serializer.BuildFileFolders(filefolder))
}
