package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service"
	"github.com/gin-gonic/gin"
)

// GetFilefolderAllFile return a file from the filefolder
func GetFilefolderAllFile(c *gin.Context) {
	var service service.FileFolderGetAllFileService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	jwtUser := c.MustGet("UserId").(string)
	res := service.GetAllFile(jwtUser)
	c.JSON(200, res)
}

// GetFilefolderAllFilefolder return a filefolder from the filefolder
func GetFilefolderAllFilefolder(c *gin.Context) {
	var service service.FileFolderGetAllFileFolderService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	jwtUser := c.MustGet("UserId").(string)
	res := service.GetAllFileFolder(jwtUser)
	c.JSON(200, res)
}
