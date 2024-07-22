package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type FileStoreGetInfoService struct {
}

// FileStoreGetInfo get user store info by userid
func (service *FileStoreGetInfoService) FileStoreGetInfo(userId string) serializer.Response {
	// get store from database
	var store model.FileStore
	if err := model.DB.Where("onwer_id = ?", userId).First(&store).Error; err != nil {
		logger.Log().Error("[FileStoreGetInfoService.FileStoreGetInfo] Fail to get user filestore info: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildFileStore(store))
}
