package service

import (
	"fmt"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileDeleteService struct {
	FileId string `json:"id" form:"id" binding:"required"`
}

// FileDelete delete file and updata user store
func (service *FileDeleteService) FileDelete(userId string) serializer.Response {
	var userFile model.File
	var userStore model.FileStore
	if err := model.DB.Where("uuid = ?", service.FileId).Find(&userFile).Error; err != nil {
		return serializer.DBErr("get file err when delete file", err)
	}
	fmt.Println(userFile.Owner, userId)
	if userFile.Owner != userId {
		return serializer.NotAuthErr("Unable to delete file that do not match the user")
	}

	if err := model.DB.Where("owner_id = ?", userId).First(&userStore).Error; err != nil {
		return serializer.DBErr("get file store err when delete file", err)
	}

	// sub delete file size
	deleteFileSize := userFile.Size
	userStore.AddCurrentSize(-deleteFileSize)

	// update database
	if err := model.DB.Delete(&userFile).Error; err != nil {
		return serializer.DBErr("delete file err when delete file", err)
	}
	if err := model.DB.Save(&userStore).Error; err != nil {
		return serializer.DBErr("update user store err when delete file", err)
	}

	return serializer.Success(nil)
}
