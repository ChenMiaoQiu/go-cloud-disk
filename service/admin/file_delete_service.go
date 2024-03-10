package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileDeleteService struct {
}

// FileDelete delete all file that have same MD5 code
func (service *FileDeleteService) FileDelete(operStatus string, fileId string) serializer.Response {
	// get delete file from database
	var deleteFile model.File
	if err := model.DB.Where("uuid = ?", fileId).Find(&deleteFile).Error; err != nil {
		return serializer.DBErr("find deletefile err when admin delete file", err)
	}

	// get file owner
	var fileOwner model.User
	if err := model.DB.Where("uuid = ?", deleteFile.Owner).Find(&fileOwner).Error; err != nil {
		return serializer.DBErr("find user err when admin delete file", err)
	}

	// if file owner is admin, cancel delete file
	if operStatus == model.StatusAdmin {
		if fileOwner.Status == model.StatusAdmin || fileOwner.Status == model.StatusSuperAdmin {
			return serializer.NotAuthErr("can't delete admin file")
		}
	}

	// delete all file
	var files []model.File
	if err := model.DB.Where("file_uuid = ?", deleteFile.FileUuid).Find(&files).Error; err != nil {
		return serializer.DBErr("get file err when admin delete file", err)
	}

	if err := model.DB.Delete(&files).Error; err != nil {
		return serializer.DBErr("delete file err when admin delete file", err)
	}

	return serializer.Success(nil)
}
