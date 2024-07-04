package file

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	logger "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type FileUpdateService struct {
	FileId      string `json:"file" form:"file" binding:"required"`
	FileName    string `json:"name" form:"name"`
	NewParentId string `json:"parent" form:"parent" binding:"required"`
}

func (service *FileUpdateService) UpdateFileInfo(userId string) serializer.Response {
	var file model.File
	var err error
	if err := model.DB.Where("uuid = ?", service.FileId).Find(&file).Error; err != nil {
		logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to find file: ", err)
		return serializer.DBErr("", err)
	}
	if file.Owner != userId {
		return serializer.NotAuthErr("")
	}

	var nowFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", file.ParentFolderId).Find(&nowFilefolder).Error; err != nil {
		logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to find filefolder: ", err)
		return serializer.DBErr("", err)
	}
	// check target filefolder owner
	var parentFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.NewParentId).Find(&parentFilefolder).Error; err != nil {
		logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to find filefolder: ", err)
		return serializer.DBErr("find parent filefolder err when update file info", err)
	}
	if userId != parentFilefolder.OwnerID {
		return serializer.NotAuthErr("")
	}

	// build new file info
	file.ParentFolderId = service.NewParentId
	newFilename := file.FileName
	if service.FileName != "" {
		newFilename = service.FileName
	}
	file.FileName = newFilename
	// update file info to database
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()
	if err := t.Save(&file).Error; err != nil {
		logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to update file: ", err)
		return serializer.DBErr("", err)
	}

	// change filefolder size
	if nowFilefolder.Uuid != parentFilefolder.Uuid {
		err = nowFilefolder.SubFileFolderSize(t, file.Size)
		if err != nil {
			logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to SubFileFolderSize: ", err)
			return serializer.DBErr("", err)
		}
		err = parentFilefolder.AddFileFolderSize(t, file.Size)
		if err != nil {
			logger.Log().Error("[FileUpdateService.UpdateFileInfo] Fail to AddFileFolderSize: ", err)
			return serializer.DBErr("", err)
		}
	}

	return serializer.Success(serializer.BuildFile(file))
}
