package file

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileUpdateService struct {
	FileId      string `json:"file" form:"file" binding:"required"`
	NewParentId string `json:"parent" form:"parent" binding:"required"`
}

func (service *FileUpdateService) UpdateFileInfo(userId string) serializer.Response {
	var file model.File
	if err := model.DB.Where("uuid = ?", userId).Find(&file).Error; err != nil {
		return serializer.DBErr("find user err when update file info", err)
	}
	if file.Owner != userId {
		return serializer.NotAuthErr("can't match user")
	}

	// check target filefolder owner
	var parentFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.NewParentId).Error; err != nil {
		return serializer.DBErr("find parent filefolder err when update file info", err)
	}
	if userId != parentFilefolder.OwnerID {
		return serializer.NotAuthErr("can't match file owner")
	}

	file.ParentFolderId = service.NewParentId
	// update file info to database
	if err := model.DB.Save(&file).Error; err != nil {
		return serializer.DBErr("save file info err when update file info", err)
	}

	return serializer.Success(serializer.BuildFile(file))
}
