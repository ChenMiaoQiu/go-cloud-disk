package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderGetAllFileService struct {
	FileFolderID string `json:"id" form:"id" binding:"required"`
}

func (service *FileFolderGetAllFileService) GetAllFile(userId string) serializer.Response {
	var fileFolder model.FileFolder
	if err := model.DB.Find(&fileFolder).Where("uuid = ?", service.FileFolderID).Error; err != nil {
		return serializer.DBErr("dberr", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("can't matched user")
	}

	var files []model.File
	if err := model.DB.Find(&files).Where("parent_folder_id = ?", service.FileFolderID).Error; err != nil {
		return serializer.DBErr("get all file db err", err)
	}
	return serializer.Success(serializer.BuildFiles(files))
}
