package file

import (
	"mime/multipart"

	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	logger "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
	"gorm.io/gorm"
)

type FileUploadService struct {
	FolderId string `form:"filefolder" json:"filefolder" binding:"required"`
}

// checkIfFileSizeExceedsVolum check if upload file size exceed user filestore size
func checkIfFileSizeExceedsVolum(userStore *model.FileStore, userId string, size int64) (bool, error) {
	if err := model.DB.Where("owner_id = ?", userId).Find(userStore).Error; err != nil {
		return false, err
	}
	ans := userStore.CurrentSize+size > userStore.MaxSize
	return ans, nil
}

// createFile use transaction to save user file info for user store safe
func createFile(t *gorm.DB, file model.File, userStore model.FileStore) error {
	// save file info to database
	var err error
	if err = t.Save(&file).Error; err != nil {
		return err
	}
	// add user file store volum
	userStore.AddCurrentSize(file.Size)
	if err = t.Save(&userStore).Error; err != nil {
		return err
	}
	return nil
}

func (service *FileUploadService) UploadFile(userId string, file *multipart.FileHeader, dst string) serializer.Response {
	// get user upload file and save it to local
	var userStore model.FileStore
	var err error

	// check if the currentSize exceeds maxsize after adding the file size
	var isExceed bool
	if isExceed, err = checkIfFileSizeExceedsVolum(&userStore, userId, file.Size); err != nil {
		logger.Log().Error("[FileUploadService.UploadFile] Fail to check user volum: ", err)
		return serializer.DBErr("", err)
	}
	if isExceed {
		return serializer.ParamsErr("ExceedStoreLimit", nil)
	}

	// upload file to cloud
	md5String, err := utils.GetFileMD5(dst)
	if err != nil {
		logger.Log().Error("[FileUploadService.UploadFile] Fail to get file md5 Code: ", err)
		return serializer.ParamsErr("", err)
	}
	// if the file has been recently uploaded, do not upload it to
	// the cloud and get file info from redis
	filePath := model.GetFileInfoFromRedis(md5String)
	if filePath == "" {
		err = disk.BaseCloudDisk.UploadSimpleFile(dst, userId, md5String, file.Size)
		if err != nil {
			logger.Log().Error("[FileUploadService.UploadFile] Fail to upload file to Cloud: ", err)
			return serializer.InternalErr("", err)
		}
		filePath = userId
	}

	// insert file to database
	filename, extend := utils.SplitFilename(file.Filename)
	fileModel := model.File{
		Owner:          userId,
		FileName:       filename,
		FilePostfix:    extend,
		FileUuid:       md5String,
		FilePath:       filePath,
		ParentFolderId: service.FolderId,
		Size:           file.Size,
	}

	t := model.DB.Begin()
	// insert user file info to database
	if err := createFile(t, fileModel, userStore); err != nil {
		logger.Log().Error("[FileUploadService.UploadFile] Fail to create file info: ", err)
		t.Rollback()
		return serializer.DBErr("", err)
	}

	// add deleted file size to filefolder and parent filefolder
	var userFileFolder model.FileFolder
	if err := t.Where("uuid = ?", service.FolderId).Find(&userFileFolder).Error; err != nil {
		logger.Log().Error("[FileUploadService.UploadFile] Fail to get filefolder info: ", err)
		t.Rollback()
		return serializer.DBErr("", err)
	}
	if err := userFileFolder.AddFileFolderSize(t, file.Size); err != nil {
		logger.Log().Error("[FileUploadService.UploadFile] Fail to update filefolder volum: ", err)
		t.Rollback()
		return serializer.DBErr("", err)
	}

	t.Commit()

	// save file info to database
	fileModel.SaveFileUploadInfoToRedis()
	return serializer.Success(serializer.BuildFile(fileModel))
}
