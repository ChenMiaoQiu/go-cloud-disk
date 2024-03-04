package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/filestore"
	"github.com/gin-gonic/gin"
)

func GetFileStoreInfo(c *gin.Context) {
	var service filestore.FileStoreGetInfoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	storeId := c.Param("filestoreId")
	userId := c.MustGet("UserId").(string)
	res := service.FileStoreGetInfo(userId, storeId)
	c.JSON(200, res)
}
