package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// file_store serializer
type FileFolder struct {
	Uuid             string `json:"filefolder_id"`
	FileFolderName   string `json:"name"`
	Filetype         string `json:"filetype"`
	ParentFileFolder string `json:"parent"`
	Size             int64  `json:"size"`
}

func BuildFileFolder(fileFolder model.FileFolder) FileFolder {
	return FileFolder{
		Uuid:             fileFolder.Uuid,
		FileFolderName:   fileFolder.FileFolderName,
		ParentFileFolder: fileFolder.ParentFolderID,
		Filetype:         "file_folder",
		Size:             fileFolder.Size,
	}
}

func BuildFileFolders(fileFolder []model.FileFolder) (FileFolderSerializers []FileFolder) {
	for _, f := range fileFolder {
		FileFolderSerializers = append(FileFolderSerializers, BuildFileFolder(f))
	}
	return
}
