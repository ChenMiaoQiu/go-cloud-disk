package model

import (
	"context"
	"math/rand"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	Uuid           string `gorm:"primarykey"`
	Owner          string // file owner if File deleted the owner is null,
	FileName       string // real filename
	FilePostfix    string
	FileUuid       string // file used md5 as name in cloud
	FilePath       string // file's filefolder in cloud, used for save share file
	ParentFolderId string
	Size           int64 // file size
}

// BeforeCreate create uuid before insert database
func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	if file.Uuid == "" {
		file.Uuid = uuid.New().String()
	}
	return
}

// GetFileInfoFromRedis get file upload path form redis
func GetFileInfoFromRedis(md5 string) string {
	filePath := cache.RedisClient.Get(context.Background(), cache.FileInfoStoreKey(md5)).Val()
	return filePath
}

// SaveFileUploadInfoToRedis save file path to redis
func (file *File) SaveFileUploadInfoToRedis() {
	randTime := time.Hour*12 + time.Minute*time.Duration(rand.Intn(60))
	cache.RedisClient.Set(context.Background(), cache.FileInfoStoreKey(file.FileUuid), file.FilePath, randTime)
}
