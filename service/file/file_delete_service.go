package file

import (
	"fmt"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
	"gorm.io/gorm"
)

type FileDeleteService struct {
}

// deleteFile delete file by transation that gorm offerd
func deleteFile(t *gorm.DB, userFile model.File, userStore model.FileStore, userFileFolder model.FileFolder) error {
	// sub deleted file size to filefolder and parent filefolder
	if err := userFileFolder.SubFileFolderSize(t, userFile.Size); err != nil {
		return fmt.Errorf("sub filefolder size err when delete file %v", err)
	}

	// sub deleted file size to userstore
	userStore.SubCurrentSize(userFile.Size)
	if err := model.DB.Delete(&userFile).Error; err != nil {
		return fmt.Errorf("delete file err when delete file %v", err)
	}
	if err := model.DB.Save(&userStore).Error; err != nil {
		return fmt.Errorf("update user store err when delete file %v", err)
	}
	return nil
}

// FileDelete delete file and updata user store
func (service *FileDeleteService) FileDelete(userId string, fileid string) serializer.Response {
	var userFile model.File
	var userStore model.FileStore
	var err error
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()

	// check file owner
	if err = t.Where("uuid = ?", fileid).Find(&userFile).Error; err != nil {
		loglog.Log().Error("[FileDeleteService.FileDelete] Fail to find user file")
		return serializer.DBErr("", err)
	}
	if userFile.Owner != userId {
		return serializer.NotAuthErr("")
	}
	if err = t.Where("owner_id = ?", userId).First(&userStore).Error; err != nil {
		loglog.Log().Error("[FileDeleteService.FileDelete] Fail to find user filestore")
		return serializer.DBErr("", err)
	}
	var userFileFolder model.FileFolder
	if err = t.Where("uuid = ?", userFile.ParentFolderId).Find(&userFileFolder).Error; err != nil {
		loglog.Log().Error("[FileDeleteService.FileDelete] Fail to find user filefolder")
		return serializer.DBErr("", err)
	}

	// use transaction to delete file
	if err = deleteFile(t, userFile, userStore, userFileFolder); err != nil {
		loglog.Log().Error("[FileDeleteService.FileDelete] Fail to update user filestore volum: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
