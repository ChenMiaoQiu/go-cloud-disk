package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/gin-gonic/gin"
)

type FileUploadService struct {
	FolderId string `form:"filefolder" json:"filefolder" binding:"required"`
}

// splitFilename split file.filename to filename and extend name
func splitFilename(str string) (filename string, extend string) {
	for i := len(str) - 1; i >= 0 && str[i] != '/'; i-- {
		if str[i] == '.' {
			if i != 0 {
				filename = str[:i-1]
			}
			if i != len(str)-1 {
				extend = str[i+1:]
			}
			return
		}
	}
	return str, ""
}

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

	// check if the currentSize exceeds maxsize after adding1
	// the file size when save file to local
	var isExceed bool
	if isExceed, err = checkIfFileSizeExceedsVolum(&userStore, userId, file.Size); err != nil {
		return serializer.DBErr("get user store err when upload file", err)
	}
	if isExceed {
		return serializer.ParamsErr("upload file size exceed user maxsize", nil)
	}

	if file == nil {
		return serializer.ParamsErr("not file", err)
	}
	dst := utils.FastBuildString("./user/", userId, "/", file.Filename)
	c.SaveUploadedFile(file, dst)

	// upload file to cloud
	md5String, err := utils.GetFileMD5(dst)
	if err != nil {
		return serializer.ParamsErr("file err", err)
	}
	err = utils.BaseCloudDisk.UploadSimpleFile(dst, userId, md5String, file.Size)
	if err != nil {
		return serializer.DBErr("can't upload to cloud", err)
	}

	// insert file to database
	filename, extend := splitFilename(file.Filename)
	fileModel := &model.File{
		Owner:          userId,
		FileName:       filename,
		FilePostfix:    extend,
		FileUuid:       md5String,
		ParentFolderId: service.FolderId,
		Size:           file.Size,
	}
	if err = model.DB.Create(&fileModel).Error; err != nil {
		return serializer.DBErr("insert file to database error when upload file", err)
	}

	// updata user store now size to database
	userStore.AddCurrentSize(file.Size)
	if err = model.DB.Save(userStore).Error; err != nil {
		return serializer.DBErr("updata userstore size err when upload file", err)
	}

	return serializer.Success(nil)
}
