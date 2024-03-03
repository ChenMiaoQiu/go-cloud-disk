package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderGetAllFileService struct {
}

// GetAllFile get all file in user filefolder
func (service *FileFolderGetAllFileService) GetAllFile(userId string, fileFolderID string) serializer.Response {
	// check if user match
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", fileFolderID).Find(&fileFolder).Error; err != nil {
		return serializer.DBErr("dberr", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("can't matched user")
	}

	var files []model.File
	if err := model.DB.Where("parent_folder_id = ?", fileFolderID).Find(&files).Error; err != nil {
		return serializer.DBErr("get all file db err", err)
	}
	return serializer.Success(serializer.BuildFiles(files))
}
