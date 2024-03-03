package share

import (
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

type ShareCreateService struct {
	FileId string `json:"fileid" form:"fileid" binding:"required"`
	Title  string `json:"title" form:"title"`
}

type createShareSuccessResponse struct {
	ShareId string `json:"shareid"`
}

func (service *ShareCreateService) CreateShare(userId string) serializer.Response {
	// check file owner
	var shareFile model.File
	if err := model.DB.Where("uuid = ?", service.FileId).Find(&shareFile).Error; err != nil {
		return serializer.DBErr("find file err when create share", err)
	}
	if shareFile.Owner != userId {
		return serializer.NotAuthErr("can't match user when create share")
	}

	// create share and save to database
	newShare := model.Share{
		Owner:       userId,
		FileId:      service.FileId,
		Title:       service.Title,
		SharingTime: time.Unix(time.Now().Unix(), 0).Format(utils.DefaultTimeTemplate),
	}
	if err := model.DB.Create(&newShare).Error; err != nil {
		return serializer.DBErr("create share err when create share", err)
	}

	return serializer.Success(createShareSuccessResponse{
		ShareId: newShare.Uuid,
	})
}
