package file

import (
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
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

func (service *FileUploadService) UploadFile(c *gin.Context) serializer.Response {
	// get user info form jwt
	userId := c.MustGet("UserId").(string)

	// get user upload file and save it to local
	var userStore model.FileStore
	file, err := c.FormFile("file")
	if err != nil {
		return serializer.ParamsErr("get upload file err", err)
	}

	// check if the currentSize exceeds maxsize after adding the file size
	var isExceed bool
	if isExceed, err = checkIfFileSizeExceedsVolum(&userStore, userId, file.Size); err != nil {
		return serializer.DBErr("get user store err when upload file", err)
	}
	if isExceed {
		return serializer.ParamsErr("upload file size exceed user maxsize", nil)
	}

	// save file to local
	if file == nil {
		return serializer.ParamsErr("not file", err)
	}
	// save file to the specified folder for easy delete file in the future
	uploadDay := time.Now().Format("2006-01-02")
	dst := utils.FastBuildString("./user/", uploadDay, "/", userId, "/", file.Filename)
	c.SaveUploadedFile(file, dst)

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
	if err = model.DB.Create(&fileModel).Error; err != nil {
		return serializer.DBErr("insert file to database error when upload file", err)
	}
	// save file info to database
	fileModel.SaveFileUploadInfoToRedis()

	// updata user store now size to database
	userStore.AddCurrentSize(file.Size)
	if err = model.DB.Save(userStore).Error; err != nil {
		return serializer.DBErr("updata userstore size err when upload file", err)
	}

	// add deleted file size to filefolder and parent filefolder
	var userFileFolder model.FileFolder
	if err := model.DB.Where("uuid = ?", service.FolderId).Find(&userFileFolder).Error; err != nil {
		return serializer.DBErr("get filefolder err when delete file", err)
	}
	if err := userFileFolder.AddFileFolderSize(file.Size); err != nil {
		return serializer.DBErr("sub filefolder size err when delete file %v", err)
	}

	return serializer.Success(serializer.BuildFile(*fileModel))
}
