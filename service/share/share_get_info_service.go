package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type ShareGetInfoService struct {
}

func (service *ShareGetInfoService) GetShareInfo(shareid string) serializer.Response {
	// get share form database
	// will add redis to rise search speed in next version
	var share model.Share
	if err := model.DB.Where("uuid = ?", shareid).Find(&share).Error; err != nil {
		return serializer.DBErr("get share err when get share info", err)
	}
	// add view of share
	share.AddViewCount()

	// build downloadurl
	downloadUrl, err := share.DownloadURL()
	if err != nil {
		return serializer.ErrorResponse(err)
	}
	return serializer.Success(serializer.BuildShareWithDownloadUrl(share, downloadUrl))
}
