package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type ShareGetAllService struct {
}

// GetAllShare get user's all share
func (service *ShareGetAllService) GetAllShare(userId string) serializer.Response {
	// get shares from database
	var shares []model.Share
	if err := model.DB.Where("owner = ?", userId).Find(&shares).Error; err != nil {
		return serializer.DBErr("get share err when get all user's share", err)
	}

	return serializer.Success(serializer.BuildShares(shares))
}
