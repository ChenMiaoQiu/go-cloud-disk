package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/google/uuid"
)

type GetUploadURLService struct {
	FileType string `form:"filetype" json:"filetype" binding:"required,min=2"`
}

type getUploadURLResponse struct {
	Url      string `json:"url"`
	FileUuid string `json:"file_uuid"`
}

func (service *GetUploadURLService) GetUploadURL(fileowner string) serializer.Response {
	fileID := uuid.New().String()
	fileName := fileID + "." + service.FileType
	url, err := utils.BaseCloudDisk.GetUploadPresignedURL(fileowner, "", fileName)
	if err != nil {
		return serializer.Err(serializer.CodeError, "get object error", err)
	}

	return serializer.Success(getUploadURLResponse{
		Url:      url,
		FileUuid: fileID,
	})
}
