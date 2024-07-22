package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type UserFilestoreUpdateService struct {
	UserId        string `json:"userid" form:"userid" required:"binding"`
	NewStoreVolum int64  `json:"volum" form:"volum" required:"binding"`
}

func (service *UserFilestoreUpdateService) UserFilestoreUpdate() serializer.Response {
	// search filestore from database
	var userFilestore model.FileStore
	if err := model.DB.Where("owner_id = ?", service.UserId).First(&userFilestore).Error; err != nil {
		logger.Log().Error("[UserFilestoreUpdateService.UserFilestoreUpdate] Fail to find filestore info: ", err)
		return serializer.DBErr("", err)
	}

	// Maximum capacity of 1GB
	userFilestore.MaxSize = min(service.NewStoreVolum, int64(1024*1024*1024))
	userFilestore.MaxSize = max(0, userFilestore.MaxSize)

	if err := model.DB.Save(&userFilestore).Error; err != nil {
		logger.Log().Error("[UserFilestoreUpdateService.UserFilestoreUpdate] Fail to update filestore info: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
