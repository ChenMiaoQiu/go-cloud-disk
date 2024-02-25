package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
)

type FileUploadService struct {
}

func (service *FileUploadService) UploadFile(c *gin.Context) serializer.Response {
	// get user info form jwt
	userId := c.MustGet("UserId").(string)

	// get user upload file and save it to local
	file, err := c.FormFile("file")
	if file == nil {
		return serializer.ParamsErr("not file", err)
	}
	dst := utils.FastBuildString("./user/", userId, "/", file.Filename)
	c.SaveUploadedFile(file, dst)

	// upload file to cloud
	md5String, err := utils.GetFileMD5(dst)
	if err != nil {
		return serializer.ParamsErr("file err", err)
	}
	err = utils.BaseCloudDisk.UploadSimpleFile(dst, userId, md5String, file.Size)
	if err != nil {
		return serializer.DBErr("can't upload to cloud", err)
	}

	return serializer.Success(nil)
}
