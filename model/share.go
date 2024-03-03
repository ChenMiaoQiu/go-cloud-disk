package model

import (
	"context"
	"fmt"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Share struct {
	Uuid        string `gorm:"primarykey"`
	Owner       string
	FileId      string
	Title       string
	SharingTime string
}

// BeforeCreate create uuid before insert database
func (file *Share) BeforeCreate(tx *gorm.DB) (err error) {
	if file.Uuid == "" {
		file.Uuid = uuid.New().String()
	}
	return
}

// DownloadURL get share download url
func (share *Share) DownloadURL() (string, error) {
	var file File
	if err := DB.Where("uuid = ?", share.FileId).Find(&file).Error; err != nil {
		return "", fmt.Errorf("find user file err when build download url %v", err)
	}
	url, err := disk.BaseCloudDisk.GetObjectURL(file.FilePath, "", file.FileUuid+"."+file.FilePostfix)
	if err != nil {
		fmt.Println(file.FilePath, " ", file.FileUuid+"."+file.FilePostfix)
		return "", fmt.Errorf("get object url err when get share download url %v", err)
	}
	return url, nil
}

// ViewCount get share view from redis
func (share *Share) ViewCount() string {
	countStr, _ := cache.RedisClient.Get(context.Background(), cache.ShareKey(share.Uuid)).Result()
	return countStr
}

// AddViewCount add share view in redis
func (share *Share) AddViewCount() {
	ctx := context.Background()
	cache.RedisClient.Incr(ctx, cache.ShareKey(share.Uuid))
	cache.RedisClient.ZIncrBy(ctx, cache.DailyRankKey, 1, share.Uuid)
}
