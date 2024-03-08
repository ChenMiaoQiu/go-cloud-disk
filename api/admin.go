package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/admin"
	"github.com/gin-gonic/gin"
)

// UpdateUserAuth change user auth
func UpdateUserAuth(c *gin.Context) {
	var service admin.UserChangeAuthService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userStatus := c.MustGet("Status").(string)
	res := service.UserChangeAuth(userStatus)
	c.JSON(200, res)
}

// SearchUser search user by uuid or username or status
func SearchUser(c *gin.Context) {
	var service admin.UserSearchService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.UserSearch()
	c.JSON(200, res)
}
