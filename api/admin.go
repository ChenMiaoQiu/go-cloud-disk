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

// UserFileStoreUpdate update user filestore
func UserFileStoreUpdate(c *gin.Context) {
	var service admin.UserFilestoreUpdateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.UserFilestoreUpdate()
	c.JSON(200, res)
}

// SearchUser search user by uuid or title or owner
func SearchShare(c *gin.Context) {
	var service admin.ShareSearchService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.ShareSearch()
	c.JSON(200, res)
}

// AdminDeleteShare delete share by uuid
func AdminDeleteShare(c *gin.Context) {
	var service admin.ShareDeleteService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	shareId := c.Param("shareId")
	res := service.ShareDelete(shareId)
	c.JSON(200, res)
}

// AdminDeleteFile delete all file that have same md5 code in database, don't delete file on cloud
func AdminDeleteFile(c *gin.Context) {
	var service admin.FileDeleteService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userStatus := c.MustGet("Status").(string)
	fileId := c.Param("fileId")
	res := service.FileDelete(userStatus, fileId)
	c.JSON(200, res)
}

// AdminGetFileStoreInfo get filestore info by userid
func AdminGetFileStoreInfo(c *gin.Context) {
	var service admin.FileStoreGetInfoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.Param("userId")
	res := service.FileStoreGetInfo(userId)
	c.JSON(200, res)
}
