package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/service"
	"github.com/gin-gonic/gin"
)

// UserLogin user login api
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserRegiser user register api
func UserRegiser(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserInfo get user info
func UserInfo(c *gin.Context) {
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetUserInfo(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMyInfo get user info form jwt info
func UserMyInfo(c *gin.Context) {
	var service service.UserInfoService

	if err := c.ShouldBind(&service); err == nil {
		userIdString := c.MustGet("UserId").(string)
		res := service.GetUserInfo(userIdString)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
