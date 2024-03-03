package file

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileDeleteService struct {
}

// FileDelete delete file and updata user store
func (service *FileDeleteService) FileDelete(userId string, fileid string) serializer.Response {
	var userFile model.File
	var userStore model.FileStore
	if err := model.DB.Where("uuid = ?", fileid).Find(&userFile).Error; err != nil {
		return serializer.DBErr("get file err when delete file", err)
	}
	if userFile.Owner != userId {
		return serializer.NotAuthErr("Unable to delete file that do not match the user")
	}

	if err := model.DB.Where("owner_id = ?", userId).First(&userStore).Error; err != nil {
		return serializer.DBErr("get file store err when delete file", err)
	}

	// sub deleted file size to filefolder and parent filefolder
	// will add rabbitMQ or kafka for enhance speed
	var userFileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", userFile.ParentFolderId).Find(&userFileFolder).Error; err != nil {
		return serializer.DBErr("get filefolder err when delete file", err)
	}
	if err := userFileFolder.SubFileFolderSize(userFile.Size); err != nil {
		return serializer.DBErr("sub filefolder size err when delete file %v", err)
	}

	// sub deleted file size to userstore
	userStore.SubCurrentSize(userFile.Size)
	if err := model.DB.Delete(&userFile).Error; err != nil {
		return serializer.DBErr("delete file err when delete file", err)
	}
	if err := model.DB.Save(&userStore).Error; err != nil {
		return serializer.DBErr("update user store err when delete file", err)
	}

	return serializer.Success(nil)
}
