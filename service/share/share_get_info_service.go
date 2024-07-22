package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type ShareGetInfoService struct {
}

func (service *ShareGetInfoService) GetShareInfo(shareid string) serializer.Response {
	share := model.Share{
		Uuid: shareid,
	}
	// try get share info from redis
	if share.CheckRedisExistsShare() {
		downloadUrl := share.GetShareInfoFromRedis()
		// check if empty share
		if downloadUrl != "" {
			share.AddViewCount()
		}
		return serializer.Success(serializer.BuildShareWithDownloadUrl(share, downloadUrl))
	}

	// can't get share info from redis search database
	if err := model.DB.Where("uuid = ?", shareid).Find(&share).Error; err != nil {
		logger.Log().Error("[ShareGetInfoService.GetShareInfo] Fail to get share info: ", err)
		return serializer.DBErr("", err)
	}

	// get downloadurl if can't get download url means share is deleted
	downloadUrl, err := share.DownloadURL()
	if err != nil {
		share.SetEmptyShare()
	}

	// if daily view count more than 20 views, add it to redis
	// for enhance search speed
	if share.DailyViewCount() > 20 {
		// if is empty share remove it from daily rank
		err := share.SaveShareInfoToRedis(downloadUrl)
		if err != nil {
			logger.Log().Error(err.Error())
		}
	}

	// add view of share
	if downloadUrl != "" {
		share.AddViewCount()
	}
	return serializer.Success(serializer.BuildShareWithDownloadUrl(share, downloadUrl))
}
