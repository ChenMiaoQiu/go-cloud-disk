package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserInfoService struct {
}

// GetUserInfo get user info by userid
func (service *UserInfoService) GetUserInfo(userid string) serializer.Response {
	var user model.User

	result := model.DB.Model(&model.User{}).Where("uuid = ?", userid).First(&user)

	if result.Error != nil {
		return serializer.ParamsErr("can't find user", result.Error)
	}

	return serializer.Success(serializer.BuildUser(user))
}
