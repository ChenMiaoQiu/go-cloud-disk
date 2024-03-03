package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileFolderUpdateService struct {
	FileFolderId      string `json:"filefolder" form:"filefolder" binding:"required"`
	NewFileFolderName string `json:"name" form:"name"`
	NewParentId       string `json:"parent" form:"parent" binding:"required"`
}

func (service *FileFolderUpdateService) UpdateFileFolderInfo(userid string) serializer.Response {
	var filefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.FileFolderId).Error; err != nil {
		return serializer.DBErr("find filefolder err when update filefolde info", err)
	}

	// check if user match filefolder owner
	if userid != filefolder.OwnerID {
		return serializer.NotAuthErr("can't match filefolder owner")
	}

	// check target filefoler owner
	var parentFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.NewParentId).Error; err != nil {
		return serializer.DBErr("find parent filefolder err when update filefolder info", err)
	}
	if userid != parentFilefolder.OwnerID {
		return serializer.NotAuthErr("can't match filefolder owner")
	}

	// get new filefolder info
	newFileFolderName := filefolder.FileFolderName
	if service.NewFileFolderName != "" {
		newFileFolderName = service.NewFileFolderName
	}
	filefolder.FileFolderName = newFileFolderName
	filefolder.ParentFolderID = service.NewParentId

	// update filefolder info to database
	if err := model.DB.Save(&filefolder).Error; err != nil {
		return serializer.DBErr("save filefolder info err when update filefolder info", err)
	}
	return serializer.Success(nil)
}
