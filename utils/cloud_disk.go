package utils

type CloudDisk interface {
	// GetUploadPresignedURL generate presigned URL.
	// user can use presigned url to upload file by put method
	GetUploadPresignedURL(userId string, filePath string, fileName string) (string, error)
	// GetDownloadPresignedURL generate presigned URL.
	// user can use presigned url to download file by get method
	GetDownloadPresignedURL(userId string, filePath string, fileName string) (string, error)
	// GetObjectURL generate a object URL. User can use URL to view the file.
	GetObjectURL(userId string, filePath string, fileName string) (string, error)
	// DeleteObject delete user object
	DeleteObject(userId string, filePath string, items []string) error
	// DeleteObjectFilefold delete user object folder
	DeleteObjectFilefolder(userId string, filePath string) error
	// IsObjectExist check file is exist
	IsObjectExist(userId string, filePath string, fileName string) (bool, error)
	// UploadSimpleFile upload file smaller than 1GB to the cloud
	UploadSimpleFile(localFilePath string, userId string, md5 string, fileSize int64) error
}

var _ CloudDisk = (*TencentCloudDisk)(nil)

type NewCloudDisk func() CloudDisk

var BaseCloudDisk CloudDisk

var NewCloudDiskMap map[string]NewCloudDisk

func init() {
	NewCloudDiskMap = make(map[string]NewCloudDisk)
	NewCloudDiskMap["TENCENT"] = NewTencentCloudDisk
}

func SetBaseCloudDisk(version string) {
	ver, ok := NewCloudDiskMap[version]
	if !ok {
		panic("unaccept this cloud disk version")
	}
	BaseCloudDisk = ver()
}
