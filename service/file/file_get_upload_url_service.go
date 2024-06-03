package file

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
	"github.com/google/uuid"
)

type GetUploadURLService struct {
	FileType string `form:"filetype" json:"filetype" binding:"required,min=1"`
}

type getUploadURLResponse struct {
	Url      string `json:"url"`
	FileUuid string `json:"file_uuid"`
}

func (service *GetUploadURLService) GetUploadURL(fileowner string) serializer.Response {
	fileID := uuid.New().String()
	fileName := fileID + "." + service.FileType
	url, err := disk.BaseCloudDisk.GetUploadPresignedURL(fileowner, "", fileName)
	if err != nil {
		loglog.Log().Error("[GetUploadURLService.GetUploadURL] Fail to get upload url: ", err)
		return serializer.InternalErr("", err)
	}

	return serializer.Success(getUploadURLResponse{
		Url:      url,
		FileUuid: fileID,
	})
}
