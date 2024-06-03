package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type ShareDeleteService struct {
}

func (service *ShareDeleteService) ShareDelete(shareId string) serializer.Response {
	// get shares from database
	if err := model.DB.Where("uuid = ?", shareId).Delete(&model.Share{}).Error; err != nil {
		loglog.Log().Error("[ShareDeleteService.ShareDelete] Fail to get share info: ", err)
		return serializer.DBErr("", err)
	}

	// delete share info that store in redis
	share := model.Share{
		Uuid: shareId,
	}
	share.DeleteShareInfoInRedis()

	return serializer.Success(nil)
}
