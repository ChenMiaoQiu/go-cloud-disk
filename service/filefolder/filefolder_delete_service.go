package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type DeleteFileFolderService struct {
}

// DeleteFileFolder tmp delete filefolder, this func will update when add size model
func (service *DeleteFileFolderService) DeleteFileFolder(userId string, fileFolderId string) serializer.Response {
	// check if user auth match this filefolder
	var fileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", fileFolderId).Find(&fileFolder).Error; err != nil {
		return serializer.DBErr("find filefolder err when delete filefolder", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("user can't match filefolder when delete filefolder")
	}

	// delete filefolder form list and protect filefolder from duplicate delete
	if err := model.DB.Delete(&fileFolder).Error; err != nil {
		return serializer.DBErr("delete filefolder err when delete filefolder func", err)
	}

	// delete filefolder size from parent filefolder
	if fileFolder.ParentFolderID != "root" {
		var parentFileFolder model.FileFolder
		if err := model.DB.Where("uuid = ?", fileFolder.ParentFolderID).Find(&parentFileFolder).Error; err != nil {
			return serializer.DBErr("find parent filefolder err when delete filefolder", err)
		}
		if err := parentFileFolder.SubFileFolderSize(fileFolder.Size); err != nil {
			return serializer.DBErr("sub parent filefodler size err when delete filefolder", err)
		}
	}

	// delete filefolder size from userSotre
	var userStore model.FileStore
	if err := model.DB.Where("uuid = ?", fileFolder.FileStoreID).Error; err != nil {
		return serializer.DBErr("can't find userStore when delete filefolder", err)
	}
	if userStore.OwnerID != userId {
		return serializer.NotAuthErr("user can't match store when delete filefolder")
	}
	userStore.SubCurrentSize(fileFolder.Size)

	return serializer.Success(nil)
}
