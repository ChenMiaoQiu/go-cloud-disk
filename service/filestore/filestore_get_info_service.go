package filestore

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileStoreGetInfoService struct {
}

func (service *FileStoreGetInfoService) FileStoreGetInfo(userId string, storeId string) serializer.Response {
	// check store owner
	var store model.FileStore
	if err := model.DB.Where("uuid = ?", storeId).Find(&store).Error; err != nil {
		return serializer.DBErr("get filestore err when get filestore info", err)
	}
	if store.OwnerID != userId {
		return serializer.NotAuthErr("can't match user when get filestore info")
	}
	return serializer.Success(serializer.BuildFileStore(store))
}
