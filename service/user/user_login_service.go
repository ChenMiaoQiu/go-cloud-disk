package user

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=3,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=40"`
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

	// admin token only given 1 hour
	if user.Status == model.StatusAdmin || user.Status == model.StatusSuperAdmin {
		token, err = utils.GenToken("miaoqiu", 1, &user)
	}

	if err != nil {
		return serializer.Err(serializer.CodeError, "token generate error", err)
	}
	return serializer.Success(returnUser{
		Token: token,
		User:  serializer.BuildUser(user),
	})
}
