package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserFilestoreUpdateService struct {
	UserId        string `json:"userid" form:"userid" required:"binding"`
	NewStoreVolum int64  `json:"volum" form:"volum" required:"binding"`
}

func (service *UserFilestoreUpdateService) UserFilestoreUpdate() serializer.Response {
	var user model.User
	if err := model.DB.Where("uuid = ?", service.UserId).Find(&user).Error; err != nil {
		return serializer.DBErr("get user err when admin get filestore info", err)
	}
	// search filestore from database
	var userFilestore model.FileStore
	if err := model.DB.Where("uuid = ?", user.UserFileStoreID).Find(&userFilestore).Error; err != nil {
		return serializer.DBErr("search filestore err when update filestore volum", err)
	}

	if userFilestore.Uuid == "" {
		return serializer.DBErr("search empty filestore when update filestore volum", nil)
	}

	// Maximum capacity of 1GB
	userFilestore.MaxSize = min(service.NewStoreVolum, int64(1024*1024*1024))
	userFilestore.MaxSize = max(0, userFilestore.MaxSize)

	if err := model.DB.Save(&userFilestore).Error; err != nil {
		return serializer.DBErr("save filestore err when update user filestore", err)
	}

	return serializer.Success(nil)
}
