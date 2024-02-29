package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service"
	"github.com/gin-gonic/gin"
)

// GetUploadURL return uploadurl
func GetUploadURL(c *gin.Context) {
	var service service.GetUploadURLService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userIdInJWT := c.MustGet("UserId").(string)
	res := service.GetUploadURL(userIdInJWT)
	c.JSON(200, res)
}

// CreateFile create file in the database
func CreateFile(c *gin.Context) {
	var service service.FileCreateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userIdInJWT := c.MustGet("UserId").(string)
	res := service.CreateFile(userIdInJWT)
	c.JSON(200, res)
}

// GetDownloadURL return a url to download file
func GetDownloadURL(c *gin.Context) {
	var service service.FileGetDownloadURLService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	fileId := c.Param("fileid")
	userIdInJWT := c.MustGet("UserId").(string)
	res := service.GetDownloadURL(userIdInJWT, fileId)
	c.JSON(200, res)
}

// UploadFile upload file to cloud
func UploadFile(c *gin.Context) {
	var service service.FileUploadService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.UploadFile(c)
	c.JSON(200, res)
}

// DeleteFile delete file in database, don't delete file on cloud
func DeleteFile(c *gin.Context) {
	var service service.FileDeleteService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	fileid := c.Param("fileid")
	userId := c.MustGet("UserId").(string)
	res := service.FileDelete(userId, fileid)
	c.JSON(200, res)
}
