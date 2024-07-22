package rank

import (
	"context"
	"sort"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type GetDailyRankService struct {
}

func (service *GetDailyRankService) GetDailyRank() serializer.Response {
	shares := make([]model.Share, 0, 16)

	// get share rank in cache
	shareRank, err := cache.RedisClient.ZRevRange(context.Background(), cache.DailyRankKey, 0, 9).Result()
	if err != nil {
		logger.Log().Error("[GetDailyRankService.GetDailyRank] Fail to get rank form cache: ", err)
		return serializer.DBErr("", err)
	}

	if len(shareRank) > 0 {
		err := model.DB.Model(&model.Share{}).Where("uuid in (?)", shareRank).Find(&shares).Error
		if err != nil {
			logger.Log().Error("[GetDailyRankService.GetDailyRank] Fail to get rank form database: ", err)
			return serializer.DBErr("", err)
		}
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

	// sort shares
	rspShare := serializer.BuildShares(shares)
	sort.Slice(rspShare, func(i, j int) bool {
		return rspShare[i].View > rspShare[j].View
	})

	return serializer.Success(rspShare)
}
