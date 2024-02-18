package service

import (
	"strconv"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserInfoService struct {
}

// GetUserInfo get user info by userid
func (service *UserInfoService) GetUserInfo(userid string) serializer.Response {
	var user model.User

	//Prevent SQL injection
	id, err := strconv.Atoi(userid)
	if err != nil {
		return serializer.ParamsErr("url sql insert", err)
	}

	result := model.DB.Model(&model.User{}).Where("id = ?", id).First(&user)

	if result.Error != nil {
		return serializer.ParamsErr("can't find user", result.Error)
	}

	return serializer.Success(serializer.BuildUser(user))
}
