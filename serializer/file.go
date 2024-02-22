package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// file serializer
type File struct {
	Uuid        string `json:"id"`
	Owner       string `json:"owner"`
	FileName    string `json:"filename"`
	FilePostfix string `json:"filetype"`
	Size        int64  `json:"size"`
}

func BuildFile(file model.File) File {
	return File{
		Uuid:        file.Uuid,
		Owner:       file.Owner,
		FileName:    file.FileName,
		FilePostfix: file.FilePostfix,
		Size:        file.Size,
	}
}

func BuildFiles(files []model.File) (FileSerializers []File) {
	for _, f := range files {
		FileSerializers = append(FileSerializers, BuildFile(f))
	}
	return
}
