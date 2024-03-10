package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserSearchService struct {
	Uuid     string `json:"uuid" form:"uuid"`
	NickName string `json:"nickname" form:"nickname"`
	Status   string `json:"status" form:"status"`
}

// UserSearch search user by uuid or nickname or status
func (service *UserSearchService) UserSearch() serializer.Response {
	var users []model.User

	// build search gorm.DB
	searchInfo := model.DB.Model(&model.User{})
	if service.Uuid != "" {
		searchInfo.Where("uuid = ?", service.Uuid)
	}
	if service.NickName != "" {
		searchInfo.Where("nick_name like ?", "%"+service.NickName+"%")
	}
	if service.Status != "" {
		searchInfo.Where("status = ?", service.Status)
	}

	// search user in database
	if err := searchInfo.Find(&users).Error; err != nil {
		return serializer.DBErr("get user info err when search user", err)
	}

	return serializer.Success(serializer.BuildUsers(users))
}
