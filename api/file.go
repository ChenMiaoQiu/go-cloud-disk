package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/file"
	"github.com/gin-gonic/gin"
)

// GetUploadURL return uploadurl
func GetUploadURL(c *gin.Context) {
	var service file.GetUploadURLService
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
	var service file.FileCreateService
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
	var service file.FileGetDownloadURLService
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
	var service file.FileUploadService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.UploadFile(c)
	c.JSON(200, res)
}

// DeleteFile delete file in database, don't delete file on cloud
func DeleteFile(c *gin.Context) {
	var service file.FileDeleteService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	fileid := c.Param("fileid")
	userId := c.MustGet("UserId").(string)
	res := service.FileDelete(userId, fileid)
	c.JSON(200, res)
}

// UpdateFile update file info, such as remove file, update filename
func UpdateFile(c *gin.Context) {
	var service file.FileUpdateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId := c.MustGet("UserId").(string)
	res := service.UpdateFileInfo(userId)
	c.JSON(200, res)
}
