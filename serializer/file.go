package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// file serializer
type File struct {
	Uuid     string `json:"file_id"`
	FileName string `json:"filename"`
	Size     int64  `json:"size"`
}

func BuildFile(file model.File) File {
	return File{
		Uuid:     file.Uuid,
		FileName: file.FileName,
		Size:     file.Size,
	}
}

func BuildFiles(files []model.File) (FileSerializers []File) {
	for _, f := range files {
		FileSerializers = append(FileSerializers, BuildFile(f))
	}
	return
}
