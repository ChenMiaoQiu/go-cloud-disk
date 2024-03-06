package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

type Share struct {
	Uuid        string `json:"shareid"`
	FileId      string `json:"sharefileid"`
	Owner       string `json:"owner"`
	Title       string `json:"title"`
	Filename    string `json:"filename"`
	SharingTime string `json:"sharetime"`
	View        string `json:"view"`
	DownloadURL string `json:"downloadurl,omitempty"`
	Size        int64  `json:"filesize"`
}

func BuildShare(share model.Share) Share {
	return Share{
		Uuid:        share.Uuid,
		FileId:      share.FileId,
		Owner:       share.Owner,
		Title:       share.Title,
		Filename:    share.FileName,
		View:        share.ViewCount(),
		SharingTime: share.SharingTime,
		Size:        share.Size,
	}
}

func BuildShareWithDownloadUrl(share model.Share, url string) Share {
	return Share{
		Uuid:        share.Uuid,
		FileId:      share.FileId,
		Owner:       share.Owner,
		Title:       share.Title,
		Filename:    share.FileName,
		View:        share.ViewCount(),
		SharingTime: share.SharingTime,
		DownloadURL: url,
		Size:        share.Size,
	}
}

func BuildShares(Shares []model.Share) (shareSerializer []Share) {
	for _, share := range Shares {
		shareSerializer = append(shareSerializer, BuildShare(share))
	}
	return
}
