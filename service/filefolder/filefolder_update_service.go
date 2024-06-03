package filefolder

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type FileFolderUpdateService struct {
	FileFolderId      string `json:"filefolder" form:"filefolder" binding:"required"`
	NewFileFolderName string `json:"name" form:"name"`
	NewParentId       string `json:"parent" form:"parent" binding:"required"`
}

func (service *FileFolderUpdateService) UpdateFileFolderInfo(userid string) serializer.Response {
	var filefolder model.FileFolder
	var err error
	if err := model.DB.Where("uuid = ? and owner_id = ?", service.FileFolderId, userid).Find(&filefolder).Error; err != nil {
		loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to find filefolder info: ", err)
		return serializer.DBErr("", err)
	}

	// check target filefoler owner
	var targetFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ? and owner_id = ?", service.NewParentId, userid).Find(&targetFilefolder).Error; err != nil {
		loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to find new parent filefolder info: ", err)
		return serializer.DBErr("", err)
	}

	var parentFilefolder model.FileFolder
	if err := model.DB.Where("uuid = ?", filefolder.ParentFolderID).Find(&parentFilefolder).Error; err != nil {
		loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to find old parent filefolder info: ", err)
		return serializer.DBErr("", err)
	}

	// get new filefolder info
	newFileFolderName := filefolder.FileFolderName
	if service.NewFileFolderName != "" {
		newFileFolderName = service.NewFileFolderName
	}
	filefolder.FileFolderName = newFileFolderName
	filefolder.ParentFolderID = service.NewParentId

	// update filefolder info to database
	t := model.DB.Begin()
	defer func() {
		if err != nil {
			t.Rollback()
		} else {
			t.Commit()
		}
	}()

	if err := t.Save(&filefolder).Error; err != nil {
		loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to update filefolder info: ", err)
		return serializer.DBErr("", err)
	}

	// change filefolder size
	if targetFilefolder.Uuid != parentFilefolder.Uuid {
		err = parentFilefolder.SubFileFolderSize(t, filefolder.Size)
		if err != nil {
			loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to update old filefolder info: ", err)
			return serializer.DBErr("", err)
		}
		targetFilefolder.AddFileFolderSize(t, filefolder.Size)
		if err != nil {
			loglog.Log().Error("[FileFolderUpdateService.UpdateFileFolderInfo] Fail to update new filefolder info: ", err)
			return serializer.DBErr("", err)
		}
	}

	return serializer.Success(nil)
}
