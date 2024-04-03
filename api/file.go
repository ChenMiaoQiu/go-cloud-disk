package api

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/file"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
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

// getUploadFile get param from request
func getUploadFileParam(c *gin.Context) (userId string, file *multipart.FileHeader, dst string, err error) {
	userId = c.MustGet("UserId").(string)
	file, err = c.FormFile("file")
	if err != nil {
		err = fmt.Errorf("get upload file err %v", err)
		return
	}
	// save file to local
	if file == nil {
		err = fmt.Errorf("not file in parmas")
		return
	}

	// Simple check if file size can be upload. Should use userstore to check
	// file can be upload when allow user upload bigger file.
	// Example, use file.checkIfFileSizeExceedsVolum() to check file can be upload
	// In this situation, use simple check to enhance api speed
	if file.Size > 1024*1024*10 {
		err = fmt.Errorf("file size too bigger")
		return
	}
	// save file to the specified folder for easy delete file in the future
	uploadDay := time.Now().Format("2006-01-02")
	dst = utils.FastBuildString("./user/", uploadDay, "/", userId, "/", file.Filename)
	c.SaveUploadedFile(file, dst)
	return
}

// UploadFile upload file to cloud
func UploadFile(c *gin.Context) {
	var service file.FileUploadService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	userId, file, dst, err := getUploadFileParam(c)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	res := service.UploadFile(userId, file, dst)
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
