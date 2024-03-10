package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileStoreGetInfoService struct {
}

// FileStoreGetInfo get user store info by userid
func (service *FileStoreGetInfoService) FileStoreGetInfo(userId string) serializer.Response {
	var user model.User
	if err := model.DB.Where("uuid = ?", userId).Find(&user).Error; err != nil {
		return serializer.DBErr("get user err when admin get filestore info", err)
	}
	// get store from database
	var store model.FileStore
	if err := model.DB.Where("uuid = ?", user.UserFileStoreID).Find(&store).Error; err != nil {
		return serializer.DBErr("get filestore err admin when get filestore info", err)
	}
	return serializer.Success(serializer.BuildFileStore(store))
}
