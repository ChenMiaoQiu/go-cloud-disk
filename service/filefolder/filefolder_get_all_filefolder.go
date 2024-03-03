package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderGetAllFileFolderService struct {
}

// GetAllFileFolder get user all filefolder form filefolder
func (service *FileFolderGetAllFileFolderService) GetAllFileFolder(userId string, fileFolderID string) serializer.Response {
	// check if user match
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", fileFolderID).Find(&fileFolder).Error; err != nil {
		return serializer.DBErr("search filefolder err when get all filefolder", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("can't matched user when get all filefolder")
	}

	var filefolder []model.FileFolder
	if err := model.DB.Where("parent_folder_id = ?", fileFolderID).Find(&filefolder).Error; err != nil {
		return serializer.DBErr("get all file db err", err)
	}
	return serializer.Success(serializer.BuildFileFolders(filefolder))
}
