package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	logger "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type ShareSearchService struct {
	Uuid  string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
	Owner string `json:"owner" form:"owner"`
}

// ShareSearch search share by uuid or title or owner
func (service *ShareSearchService) ShareSearch() serializer.Response {
	var shares []model.Share

	// build search condition
	searchInfo := model.DB.Model(&model.Share{})
	if service.Uuid != "" {
		searchInfo.Where("uuid = ?", service.Uuid)
	}
	if service.Title != "" {
		searchInfo.Where("title like ?", "%"+service.Title+"%")
	}
	if service.Owner != "" {
		searchInfo.Where("status = ?", service.Owner)
	}

	// search share from database
	if err := searchInfo.Find(&shares).Error; err != nil {
		logger.Log().Error("[ShareSearchService.ShareSearch] Fail to find share: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(serializer.BuildShares(shares))
}
