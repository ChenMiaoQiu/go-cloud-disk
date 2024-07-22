package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type ShareSaveFileService struct {
	FileId         string `json:"fileid" form:"fileid" binding:"required"`
	SaveFilefolder string `json:"filefolder" form:"filefolder" binding:"required"`
}

func (service *ShareSaveFileService) ShareSaveFile(userId string) serializer.Response {
	// get save file from database
	var saveFile model.File
	var err error
	if err = model.DB.Where("uuid = ?", service.FileId).Find(&saveFile).Error; err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to find file info: ", err)
		return serializer.DBErr("", err)
	}

	// get save Filefolder from database and check owner
	var targetFilefolder model.FileFolder
	if err = model.DB.Where("uuid = ? and owner_id = ?", service.SaveFilefolder, userId).Find(&targetFilefolder).Error; err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to find filefolder: ", err)
		return serializer.DBErr("", err)
	}

	// get user filefolder from database
	var targetFileStore model.FileStore
	if err := model.DB.Where("uuid = ?", targetFilefolder.FileStoreID).Find(&targetFileStore).Error; err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to find filestore: ", err)
		return serializer.DBErr("", err)
	}

	// check if current size exceed when add file size
	if targetFileStore.CurrentSize+saveFile.Size > targetFileStore.MaxSize {
		return serializer.ParamsErr("ExceedStoreLimit", nil)
	}
	// change filefolder size
	targetFileStore.AddCurrentSize(saveFile.Size)
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()

	if err := t.Save(&targetFileStore).Error; err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to updata userstore: ", err)
		return serializer.DBErr("", err)
	}
	if err := targetFilefolder.AddFileFolderSize(t, saveFile.Size); err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to add filefolder size: ", err)
		return serializer.DBErr("", err)
	}

	// save file to filefolder
	newFile := model.File{
		Owner:          targetFileStore.OwnerID,
		FileName:       saveFile.FileName,
		FilePostfix:    saveFile.FilePostfix,
		FileUuid:       saveFile.FileUuid,
		FilePath:       saveFile.FilePath,
		Size:           saveFile.Size,
		ParentFolderId: service.SaveFilefolder,
	}
	if err := model.DB.Create(&newFile).Error; err != nil {
		logger.Log().Error("[ShareSaveFileService.ShareSaveFile] Fail to create file: ", err)
		return serializer.DBErr("", err)
	}

	return serializer.Success(nil)
}
