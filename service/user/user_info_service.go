package user

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	logger "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type UserInfoService struct {
}

// GetUserInfo get user info by userid
func (service *UserInfoService) GetUserInfo(userid string) serializer.Response {
	var user model.User

	err := model.DB.Model(&model.User{}).Where("uuid = ?", userid).First(&user).Error
	if err != nil {
		logger.Log().Error("[UserInfoService.GetUserInfo] Fail to find user")
		return serializer.ParamsErr("NotFound", err)
	}

	return serializer.Success(serializer.BuildUser(user))
}
