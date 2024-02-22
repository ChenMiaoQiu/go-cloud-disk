package service

type GetUploadURLService struct {
	Fileowner string `form:"owner" json:"owner" binding:"require"`
	FileType  string `form:"filetype" json:"filetype" binding:"require, min=2"`
	FilePath  string `form:"filepath" json:"filepath" binding:"require"`
}
