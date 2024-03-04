package rank

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type GetDailyRankService struct {
}

func (service *GetDailyRankService) GetDailyRank() serializer.Response {
	var shares []model.Share

	// get share rank in cache
	shareRank, err := cache.RedisClient.ZRevRange(context.Background(), cache.DailyRankKey, 0, 9).Result()
	if err != nil {
		return serializer.DBErr("get daily rank err in cache", err)
	}

	// get share info from database by share id
	for _, shareId := range shareRank {
		var share model.Share
		if err := model.DB.Where("uuid = ?", shareId).Find(&share).Error; err != nil {
			return serializer.DBErr("get share err when get daily rank", err)
		}
		shares = append(shares, share)
	}

	// fill empty share to share
	emptyShare := model.Share{
		Uuid:        "",
		Owner:       "",
		FileId:      "",
		Title:       "虚位以待",
		SharingTime: "",
	}
	for len(shares) < 10 {
		shares = append(shares, emptyShare)
	}

	return serializer.Success(serializer.BuildShares(shares))
}
