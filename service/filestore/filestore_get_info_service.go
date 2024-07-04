package filestore

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	logger "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type FileStoreGetInfoService struct {
}

func (service *FileStoreGetInfoService) FileStoreGetInfo(userId string, storeId string) serializer.Response {
	// check store owner
	var store model.FileStore
	if err := model.DB.Where("uuid = ? and owner_id = ?", storeId, userId).Find(&store).Error; err != nil {
		logger.Log().Error("[FileStoreGetInfoService.FileStoreGetInfo] Fail to find user store: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildFileStore(store))
}
