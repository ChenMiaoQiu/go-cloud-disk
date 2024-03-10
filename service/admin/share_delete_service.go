package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type ShareDeleteService struct {
}

func (service *ShareDeleteService) ShareDelete(shareId string) serializer.Response {
	// get shares from database
	var share model.Share
	if err := model.DB.Where("uuid = ?", shareId).Find(&share).Error; err != nil {
		return serializer.DBErr("get share err when get all user's share", err)
	}

	if err := model.DB.Delete(&share).Error; err != nil {
		return serializer.DBErr("delete share err when user delete share", err)
	}
	// delete share info that store in redis
	share.DeleteShareInfoInRedis()

	return serializer.Success(nil)
}
