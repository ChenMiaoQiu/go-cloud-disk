package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type ShareDeleteService struct {
}

func (service *ShareDeleteService) DeleteShare(shareId string, userId string) serializer.Response {
	// get shares from database
	var share model.Share
	if err := model.DB.Where("uuid = ?", shareId).Find(&share).Error; err != nil {
		return serializer.DBErr("get share err when get all user's share", err)
	}

	// check share owner
	if share.Owner != userId {
		return serializer.NotAuthErr("can't match user when delete share")
	}

	// delay double delete, ensure the safe of info
	share.DeleteShareInfoInRedis()
	if err := model.DB.Delete(&share).Error; err != nil {
		return serializer.DBErr("delete share err when user delete share", err)
	}
	share.DeleteShareInfoInRedis()

	return serializer.Success(nil)
}
