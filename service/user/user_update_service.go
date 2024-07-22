package user

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type UserUpdateService struct {
	NickName string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
}

// UpdateUserInfo update user nickname
func (service *UserUpdateService) UpdateUserInfo(userId string) serializer.Response {
	// check if user match userid
	var user model.User
	if err := model.DB.Where("uuid = ?", userId).Find(&user).Error; err != nil {
		logger.Log().Error("[UserUpdateService.UpdateUserInfo] Fail to find user: ", err)
		return serializer.DBErr("", err)
	}
	// update user to database
	user.NickName = service.NickName
	if err := model.DB.Save(&user).Error; err != nil {
		logger.Log().Error("[UserUpdateService.UpdateUserInfo] Fail to save user info: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildUser(user))
}
