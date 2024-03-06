package file

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileUpdateService struct {
	FileId      string `json:"file" form:"file" binding:"required"`
	FileName    string `json:"name" form:"name"`
	NewParentId string `json:"parent" form:"parent" binding:"required"`
}

func (service *FileUpdateService) UpdateFileInfo(userId string) serializer.Response {
	var file model.File
	if err := model.DB.Where("uuid = ?", service.FileId).Find(&file).Error; err != nil {
		return serializer.DBErr("find user err when update file info", err)
	}
	if file.Owner != userId {
		return serializer.NotAuthErr("can't match user")
	}

	var nowFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", file.ParentFolderId).Find(&nowFilefolder).Error; err != nil {
		return serializer.DBErr("find now filefolder err when update file info", err)
	}
	// check target filefolder owner
	var parentFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.NewParentId).Find(&parentFilefolder).Error; err != nil {
		return serializer.DBErr("find parent filefolder err when update file info", err)
	}
	if userId != parentFilefolder.OwnerID {
		return serializer.NotAuthErr("can't match file owner")
	}

	// new file info
	file.ParentFolderId = service.NewParentId
	newFilename := file.FileName
	if service.FileName != "" {
		newFilename = service.FileName
	}
	file.FileName = newFilename
	// update file info to database
	if err := model.DB.Save(&file).Error; err != nil {
		return serializer.DBErr("save file info err when update file info", err)
	}

	// change filefolder size
	if nowFilefolder.Uuid != parentFilefolder.Uuid {
		nowFilefolder.SubFileFolderSize(file.Size)
		parentFilefolder.AddFileFolderSize(file.Size)
	}

	return serializer.Success(serializer.BuildFile(file))
}
