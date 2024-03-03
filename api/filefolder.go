package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/filefolder"
	"github.com/gin-gonic/gin"
)

// GetFilefolderAllFile return a file from the filefolder
func GetFilefolderAllFile(c *gin.Context) {
	var service filefolder.FileFolderGetAllFileService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	fileFolderId := c.Param("filefolderid")
	jwtUser := c.MustGet("UserId").(string)
	res := service.GetAllFile(jwtUser, fileFolderId)
	c.JSON(200, res)
}

// GetFilefolderAllFilefolder return a filefolder from the filefolder
func GetFilefolderAllFilefolder(c *gin.Context) {
	var service filefolder.FileFolderGetAllFileFolderService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	fileFolderId := c.Param("filefolderid")
	jwtUser := c.MustGet("UserId").(string)
	res := service.GetAllFileFolder(jwtUser, fileFolderId)
	c.JSON(200, res)
}

// CreateFileFolder create filefolder in parent filefolder that user send
// and return a filefolder info
func CreateFileFolder(c *gin.Context) {
	var service filefolder.FileFolderCreateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	jwtUser := c.MustGet("UserId").(string)
	res := service.CreateFileFolder(jwtUser)
	c.JSON(200, res)
}

// DeleteFileFolder delete filefolder by filefolderid
func DeleteFileFolder(c *gin.Context) {
	var service filefolder.DeleteFileFolderService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	jwtUser := c.MustGet("UserId").(string)
	fileFolderId := c.Param("filefolderid")
	res := service.DeleteFileFolder(jwtUser, fileFolderId)
	c.JSON(200, res)
}

// UpdateFileFolder update filefolder name or filefolder position
func UpdateFileFolder(c *gin.Context) {
	var service filefolder.FileFolderUpdateService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	jwtUser := c.MustGet("UserId").(string)
	res := service.UpdateFileFolderInfo(jwtUser)
	c.JSON(200, res)
}
