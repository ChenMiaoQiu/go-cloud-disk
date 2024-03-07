package file

import (
	"fmt"

	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type FileGetDownloadURLService struct {
}

type fileGetDownloadURLResponse struct {
	Url string `json:"dowload_url"`
}

func (service *FileGetDownloadURLService) GetDownloadURL(userId string, fileid string) serializer.Response {
	var file model.File
	if err := model.DB.Where("uuid = ?", fileid).Find(&file).Error; err != nil {
		return serializer.DBErr("can't find file", err)
	}

	if userId != file.Owner {
		return serializer.NotLogin("unauth")
	}

	fileName := file.FileUuid + "." + file.FilePostfix
	url, err := disk.BaseCloudDisk.GetObjectURL(file.FilePath, "", fileName)
	if err != nil {
		return serializer.ErrorResponse(fmt.Errorf("can't get object url %v", err))
	}
	return serializer.Success(fileGetDownloadURLResponse{
		Url: url,
	})
}
