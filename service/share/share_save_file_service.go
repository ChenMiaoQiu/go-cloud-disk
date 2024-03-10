package share

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type ShareSaveFileService struct {
	FileId         string `json:"fileid" form:"fileid" binding:"required"`
	SaveFilefolder string `json:"filefolder" form:"filefolder" binding:"required"`
}

func (service *ShareSaveFileService) ShareSaveFile(userId string) serializer.Response {
	// get save file from database
	var saveFile model.File
	if err := model.DB.Where("uuid = ?", service.FileId).Find(&saveFile).Error; err != nil {
		return serializer.DBErr("can't find file when share save file", err)
	}

	// get save Filefolder from database and check owner
	var targetFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.SaveFilefolder).Find(&targetFilefolder).Error; err != nil {
		return serializer.DBErr("can't find filefolder when save file", err)
	}
	if targetFilefolder.OwnerID != userId {
		return serializer.NotAuthErr("can't match user when share save file")
	}

	// get user filefolder from database
	var targetFileStore model.FileStore
	if err := model.DB.Where("uuid = ?", targetFilefolder.FileStoreID).Find(&targetFileStore).Error; err != nil {
		return serializer.DBErr("can't find filestore when save file", err)
	}

	// check if current size exceed when add file size
	if targetFileStore.CurrentSize+saveFile.Size > targetFileStore.MaxSize {
		return serializer.NotAuthErr("file size beyond filestore max size")
	}
	// change filefolder size
	targetFileStore.AddCurrentSize(saveFile.Size)
	if err := model.DB.Save(&targetFileStore).Error; err != nil {
		return serializer.DBErr("updata userstore size err when save file", err)
	}
	if err := targetFilefolder.AddFileFolderSize(saveFile.Size); err != nil {
		return serializer.DBErr("add filefolder size err when save file", err)
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
		return serializer.DBErr("create file err when save file", err)
	}

	return serializer.Success(nil)
}
