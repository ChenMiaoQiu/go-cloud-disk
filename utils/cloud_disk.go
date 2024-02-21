package utils

type CloudDisk interface {
	// GetUploadPresignedURL generate presigned URL.
	// user can use presigned url to download file
	GetUploadPresignedURL(user string, filePath string) (string, error)
	// GetDownloadPresignedURL generate presigned URL.
	// user can use presigned url to download file
	GetDownloadPresignedURL(user string, filePath string, fileName string) (string, error)
	// GetObjectURL generate a object URL. User can use URL to view the file.
	GetObjectURL(user string, filePath string, fileName string) (string, error)
	// DeleteObject delete user object
	DeleteObject(user string, items []string)
	// DeleteObjectFilefold delete user object folder
	DeleteObjectFilefolder()
}

type NewCloudDisk func() CloudDisk

var DefaultCloudDisk CloudDisk

var NewCloudDiskMap map[string]NewCloudDisk

func init() {
	NewCloudDiskMap = make(map[string]NewCloudDisk)
}
