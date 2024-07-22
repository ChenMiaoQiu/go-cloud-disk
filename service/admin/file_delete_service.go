package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type FileDeleteService struct {
}

// FileDelete delete all file that have same MD5 code
func (service *FileDeleteService) FileDelete(operStatus string, fileId string) serializer.Response {
	// get delete file from database
	var err error
	var deleteFile model.File
	if err = model.DB.Where("uuid = ?", fileId).Find(&deleteFile).Error; err != nil {
		logger.Log().Error("[FileDeleteService.FileDelete] Fail find delete file info: ", err)
		return serializer.DBErr("", err)
	}

	// get file owner
	var fileOwner model.User
	if err = model.DB.Where("uuid = ?", deleteFile.Owner).Find(&fileOwner).Error; err != nil {
		logger.Log().Error("[FileDeleteService.FileDelete] Fail find delete file onwer: ", err)
		return serializer.DBErr("", err)
	}

	// can't delete admin file
	if operStatus == model.StatusAdmin {
		if fileOwner.Status == model.StatusAdmin || fileOwner.Status == model.StatusSuperAdmin {
			return serializer.NotAuthErr("")
		}
	}

	// delete all file
	var files []model.File
	if err = model.DB.Where("file_uuid = ?", deleteFile.FileUuid).Find(&files).Error; err != nil {
		logger.Log().Error("[FileDeleteService.FileDelete] Fail find delete files: ", err)
		return serializer.DBErr("", err)
	}

	if err = model.DB.Delete(&files).Error; err != nil {
		logger.Log().Error("[FileDeleteService.FileDelete] Fail to delete files: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
