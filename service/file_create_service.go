package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

type FileCreateService struct {
	FileName       string `json:"filename" form:"filename" binding:"required"`
	FilePostfix    string `json:"file_postfix" form:"file_postfix" binding:"required"`
	FileUuid       string `json:"file_uuid" form:"file_uuid" binding:"required"`
	ParentFolderId string `json:"folder" form:"folder" binding:"required"`
	Size           int64  `json:"size" form:"size" binding:"required"`
}

func (service *FileCreateService) CreateFile(owner string) serializer.Response {
	// check if the file was successfully uploaded to the cloud
	uploadFileNameInCloud := utils.FastBuildFileName(service.FileUuid, service.FilePostfix)
	successUpload, err := utils.BaseCloudDisk.IsObjectExist(owner, "", uploadFileNameInCloud)
	if err != nil {
		return serializer.ErrorResponse(err)
	}
	if !successUpload {
		return serializer.DBErr("upload faild", nil)
	}

	// check filefolder auth
	var fileFolder model.FileFolder
	if err = model.DB.Find(&fileFolder).Where("uuid = ?", service.FileUuid).Error; err != nil {
		return serializer.DBErr("can't find filefolder", err)
	}

	if fileFolder.OwnerID != owner {
		return serializer.NotAuthErr("")
	}

	// create file in the database
	file := model.File{
		Owner:          owner,
		FileName:       service.FileName,
		FilePostfix:    service.FilePostfix,
		FileUuid:       service.FileUuid,
		ParentFolderId: service.ParentFolderId,
		Size:           service.Size,
	}

	if err = model.DB.Create(&file).Error; err != nil {
		return serializer.DBErr("create file error", err)
	}
	return serializer.Success(nil)
}
