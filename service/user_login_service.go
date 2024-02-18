package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

type returnUser struct {
	Token string `json:"token"`
	serializer.User
}

// Login check username and password can matched
// and return user info and jwt token
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamsErr("账号或密码错误", nil)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamsErr("账号或密码错误", nil)
	}
	token, err := utils.GenToken("miaoqiu", 24, &user)
	if err != nil {
		return serializer.Err(serializer.CodeError, "token generate error", err)
	}
	return serializer.Response{
		Code: serializer.CodeSuccess,
		Msg:  "success",
		Data: returnUser{
			Token: token,
			User:  serializer.BuildUser(user),
		},
	}
}
