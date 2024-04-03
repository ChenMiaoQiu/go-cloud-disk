package file

import (
	"fmt"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"gorm.io/gorm"
)

type FileDeleteService struct {
}

// deleteFile delete file by transation that gorm offerd
func deleteFile(userFile *model.File, userStore *model.FileStore, userFileFolder *model.FileFolder) error {
	deleteFileFunc := func(tx *gorm.DB) error {
		// sub deleted file size to filefolder and parent filefolder
		if err := userFileFolder.SubFileFolderSize(userFile.Size); err != nil {
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

	// exce
	if err := model.DB.Transaction(deleteFileFunc); err != nil {
		return err
	}
	return nil
}

// FileDelete delete file and updata user store
func (service *FileDeleteService) FileDelete(userId string, fileid string) serializer.Response {
	var userFile model.File
	var userStore model.FileStore

	// check file owner
	if err := model.DB.Where("uuid = ?", fileid).Find(&userFile).Error; err != nil {
		return serializer.DBErr("get file err when delete file", err)
	}
	if userFile.Owner != userId {
		return serializer.NotAuthErr("Unable to delete file that do not match the user")
	}
	if err := model.DB.Where("owner_id = ?", userId).First(&userStore).Error; err != nil {
		return serializer.DBErr("get file store err when delete file", err)
	}
	var userFileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", userFile.ParentFolderId).Find(&userFileFolder).Error; err != nil {
		return serializer.DBErr("get filefolder err when delete file", err)
	}

	// use transaction to delete file
	if err := deleteFile(&userFile, &userStore, &userFileFolder); err != nil {
		return serializer.DBErr("update info to database err", err)
	}

	return serializer.Success(nil)
}
