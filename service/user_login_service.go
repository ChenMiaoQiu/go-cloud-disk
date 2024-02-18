package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamsErr("账号或密码错误", nil)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamsErr("账号或密码错误", nil)
	}

	return serializer.Success(serializer.BuildUser(user))
}
