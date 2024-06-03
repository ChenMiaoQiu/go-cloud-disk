package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type DeleteFileFolderService struct {
}

// DeleteFileFolder tmp delete filefolder, this func will update when add size model
func (service *DeleteFileFolderService) DeleteFileFolder(userId string, fileFolderId string) serializer.Response {
	// check if user auth match this filefolder
	var fileFolder model.FileFolder
	var err error
	if err := model.DB.Where("uuid = ?", fileFolderId).Find(&fileFolder).Error; err != nil {
		loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filefolder info: ", err)
		return serializer.DBErr("", err)
	}
	if fileFolder.OwnerID != userId {
		return serializer.NotAuthErr("")
	}

	// delete filefolder form list and protect filefolder from duplicate delete
	if fileFolder.ParentFolderID == "root" || fileFolder.ParentFolderID == "" {
		return serializer.ParamsErr("CanDeleteRoot", nil)
	}
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()

	if err := t.Delete(&fileFolder).Error; err != nil {
		loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filefolder info: ", err)
		return serializer.DBErr("delete filefolder err when delete filefolder func", err)
	}

	// delete filefolder size from parent filefolder
	if fileFolder.ParentFolderID != "root" {
		var parentFileFolder model.FileFolder
		if err := t.Where("uuid = ?", fileFolder.ParentFolderID).Find(&parentFileFolder).Error; err != nil {
			loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filefolder info: ", err)
			return serializer.DBErr("", err)
		}
		if err := parentFileFolder.SubFileFolderSize(t, fileFolder.Size); err != nil {
			loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to update parent filefolder info: ", err)
			return serializer.DBErr("", err)
		}
	}

	// delete filefolder size from userSotre
	var userStore model.FileStore
	if err := t.Where("uuid = ? and owner_id = ?", fileFolder.FileStoreID, userId).Find(&userStore).Error; err != nil {
		loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to find filestore: ", err)
		return serializer.DBErr("", err)
	}
	userStore.AddCurrentSize(fileFolder.Size)
	if err = t.Save(&userStore).Error; err != nil {
		loglog.Log().Error("[DeleteFileFolderService.DeleteFileFolder] Fail to update filestore: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
