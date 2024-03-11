package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/user"
	"github.com/gin-gonic/gin"
)

// UserLogin user login api
func UserLogin(c *gin.Context) {
	var service user.UserLoginService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.Login(c)
	c.JSON(200, res)
}

// UserRegiser user register api
func UserRegiser(c *gin.Context) {
	var service user.UserRegisterService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.Register()
	c.JSON(200, res)
}

// UserInfo get user info
func UserInfo(c *gin.Context) {
	var service user.UserInfoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.GetUserInfo(c.Param("id"))
	c.JSON(200, res)
}

// UserMyInfo get user info form jwt info
func UserMyInfo(c *gin.Context) {
	var service user.UserInfoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userIdString := c.MustGet("UserId").(string)
	res := service.GetUserInfo(userIdString)
	c.JSON(200, res)
}

// UpdateUserInfo update user nickname
func UpdateUserInfo(c *gin.Context) {
	var service user.UserUpdateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.MustGet("UserId").(string)
	res := service.UpdateUserInfo(userId)
	c.JSON(200, res)
}

// ConfirmUserEmail send confirm email
func ConfirmUserEmail(c *gin.Context) {
	var service user.UserSendConfirmEmailService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	res := service.SendConfirmEmail()
	c.JSON(200, res)
}
