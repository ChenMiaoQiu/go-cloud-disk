package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}
