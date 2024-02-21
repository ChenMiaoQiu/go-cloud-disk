package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	Uuid           string `gorm:"primarykey"`
	Owner          string // file owner
	FileName       string // real filename
	FilePostfix    string
	FileUuid       string // filename in cloud
	FilePath       string
	ParentFolderId string
	UploadTime     string
	Size           int64 // file size
}

// BeforeCreate create uuid before insert database
func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	file.Uuid = uuid.New().String()
	return
}
