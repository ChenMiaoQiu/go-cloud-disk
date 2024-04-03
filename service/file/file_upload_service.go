package file

import (
	"mime/multipart"

	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
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
func createFile(file *model.File, userStore *model.FileStore) error {
	createFileFunc := func(tx *gorm.DB) error {
		// save file info to database
		if e := model.DB.Save(file).Error; e != nil {
			return e
		}
		// add user file store volum
		userStore.AddCurrentSize(file.Size)
		if e := model.DB.Save(userStore).Error; e != nil {
			return e
		}
		return nil
	}

	if err := model.DB.Transaction(createFileFunc); err != nil {
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
		return serializer.DBErr("get user store err when upload file", err)
	}
	if isExceed {
		return serializer.ParamsErr("upload file size exceed user maxsize", nil)
	}

	// upload file to cloud
	md5String, err := utils.GetFileMD5(dst)
	if err != nil {
		return serializer.ParamsErr("file err", err)
	}
	// if the file has been recently uploaded, do not upload it to
	// the cloud and get file info from redis
	filePath := model.GetFileInfoFromRedis(md5String)
	if filePath == "" {
		err = disk.BaseCloudDisk.UploadSimpleFile(dst, userId, md5String, file.Size)
		if err != nil {
			return serializer.DBErr("can't upload to cloud", err)
		}
		filePath = userId
	}

	// insert file to database
	filename, extend := utils.SplitFilename(file.Filename)
	fileModel := &model.File{
		Owner:          userId,
		FileName:       filename,
		FilePostfix:    extend,
		FileUuid:       md5String,
		FilePath:       filePath,
		ParentFolderId: service.FolderId,
		Size:           file.Size,
	}

	// insert user file info to database
	if err := createFile(fileModel, &userStore); err != nil {
		return serializer.DBErr("create file err when upload file", err)
	}

	// save file info to database
	fileModel.SaveFileUploadInfoToRedis()

	// add deleted file size to filefolder and parent filefolder
	var userFileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.FolderId).Find(&userFileFolder).Error; err != nil {
		return serializer.DBErr("get filefolder err when upload file", err)
	}
	if err := userFileFolder.AddFileFolderSize(file.Size); err != nil {
		return serializer.DBErr("sub filefolder size err when upload file %v", err)
	}

	return serializer.Success(serializer.BuildFile(*fileModel))
}
