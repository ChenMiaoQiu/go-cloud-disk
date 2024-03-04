package api

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/service/rank"
	"github.com/gin-gonic/gin"
)

func GetDailyRank(c *gin.Context) {
	var service rank.GetDailyRankService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}

	res := service.GetDailyRank()
	c.JSON(200, res)
}
