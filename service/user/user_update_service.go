package user

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserUpdateService struct {
	NickName string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
}

// UpdateUserInfo update user nickname
func (service *UserUpdateService) UpdateUserInfo(userId string) serializer.Response {
	// check if user match userid
	var user model.User
	if err := model.DB.Where("uuid = ?", userId).Find(user).Error; err != nil {
		return serializer.DBErr("can't find user when update user info", err)
	}
	// update user to database
	user.NickName = service.NickName
	if err := model.DB.Save(user).Error; err != nil {
		return serializer.DBErr("can't save user when update user info", err)
	}
	return serializer.Success(nil)
}
