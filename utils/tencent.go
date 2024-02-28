package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCloudDisk struct {
	bucketname string
	secretId   string
	secretKey  string
}

func (cloud *TencentCloudDisk) getDefaultClient() *cos.Client {
	u, _ := url.Parse(cloud.bucketname)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRET_ID"),
			SecretKey: os.Getenv("SECRET_KEY"),
		},
	})
	return c
}

// getUploadPresignedURLPresigned use name string to generate presigned url
// user can use persigned url to upload file
func (cloud *TencentCloudDisk) getUploadPresignedURLPresigned(key string) (string, error) {
	u, _ := url.Parse(cloud.bucketname)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})
	ctx := context.Background()

	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}
	opt.Query.Add("x-cos-security-token", "<token>")
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodPut, key, cloud.secretId, cloud.secretKey, time.Minute*15, opt)
	if err != nil {
		return "", fmt.Errorf("create getUploadPresignedURLPresigned error %v", err)
	}
	return presignedURL.String(), nil
}

// GetUploadPresignedURL use userId, filePath, fileName to generate cloud disk key
func (cloud *TencentCloudDisk) GetUploadPresignedURL(userId string, filePath string, fileName string) (string, error) {
	key := fastBuildKey(userId, filePath, fileName)
	presignedURL, err := cloud.getUploadPresignedURLPresigned(key)
	if err != nil {
		return "", err
	}
	return presignedURL, nil
}

// getObjectUrl use key to generate objecturl, user can user objectURL to
// download file or view photo
func (cloud *TencentCloudDisk) getObjectUrl(key string) (str string, err error) {
	var ok bool
	if ok, err = cloud.checkObjectIsExist(key); err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("this object don't exist on cloud")
	}
	client := cloud.getDefaultClient()
	ourl := client.Object.GetObjectURL(key)
	return ourl.String(), nil
}

// GetObjectURL use userId, filePath, fileName to generate cloud disk key
func (cloud *TencentCloudDisk) GetObjectURL(userId string, filePath string, fileName string) (string, error) {
	key := fastBuildKey(userId, filePath, fileName)
	objectURL, err := cloud.getObjectUrl(key)
	if err != nil {
		return "", err
	}
	return objectURL, nil
}

// deleteObject delte multi object in cloud
func (cloud *TencentCloudDisk) deleteObject(keys []string) error {
	client := cloud.getDefaultClient()
	obs := []cos.Object{}
	for _, v := range keys {
		obs = append(obs, cos.Object{Key: v})
	}
	opt := &cos.ObjectDeleteMultiOptions{
		Objects: obs,
	}

	_, _, err := client.Object.DeleteMulti(context.Background(), opt)
	if err != nil {
		return fmt.Errorf("delete object error %v", err)
	}
	return nil
}

// DeleteObject use items to build keys
func (cloud *TencentCloudDisk) DeleteObject(userId string, filePath string, items []string) error {
	var keys []string
	for _, file := range items {
		key := fastBuildKey(userId, filePath, file)
		keys = append(keys, key)
	}
	err := cloud.deleteObject(keys)
	return err
}

func (cloud *TencentCloudDisk) deleteFilefold(dir string) error {
	client := cloud.getDefaultClient()
	var marker string
	opt := &cos.BucketGetOptions{
		Prefix:  dir,
		MaxKeys: 1000,
	}

	isTruncated := true
	var errInTruncated error
	for isTruncated {
		opt.Marker = marker
		v, _, err := client.Bucket.Get(context.Background(), opt)
		if err != nil {
			errInTruncated = err
			break
		}
		for _, content := range v.Contents {
			_, err = client.Object.Delete(context.Background(), content.Key)
			if err != nil {
				errInTruncated = err
				break
			}
		}
		if errInTruncated != nil {
			break
		}
		isTruncated = v.IsTruncated
		marker = v.NextMarker
	}
	if errInTruncated != nil {
		return errInTruncated
	}
	return nil
}

func (cloud *TencentCloudDisk) DeleteObjectFilefolder(userId string, filePath string) error {
	key := fastBuildKey(userId, filePath, "")
	err := cloud.deleteFilefold(key)
	return err
}

func (cloud *TencentCloudDisk) checkObjectIsExist(key string) (bool, error) {
	client := cloud.getDefaultClient()
	ok, err := client.Object.IsExist(context.Background(), key)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// IsObjectExist check object is exist
func (cloud *TencentCloudDisk) IsObjectExist(userId string, filePath string, fileName string) (bool, error) {
	key := fastBuildKey(userId, filePath, fileName)
	ok, err := cloud.checkObjectIsExist(key)
	return ok, err
}

// uploadSimpleFile use PutFromFile upload local file to the cloud
func (cloud *TencentCloudDisk) uploadSimpleFile(localFilePath string, key string) error {
	client := cloud.getDefaultClient()
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
	}
	_, err := client.Object.PutFromFile(context.Background(), key, localFilePath, opt)
	if err != nil {
		return err
	}
	return nil
}

// UploadSimpleFile upload file smaller than 1GB to the cloud
func (cloud *TencentCloudDisk) UploadSimpleFile(localFilePath string, userId string, md5 string, fileSize int64) error {
	if fileSize/1024/1024/1024 > 1 {
		return fmt.Errorf("file to big, please use uploadfile")
	}

	// check if the file exists on cloud
	extend := path.Ext(localFilePath)
	ok, err := cloud.IsObjectExist(userId, "", md5+extend)
	if err != nil {
		return err
	}

	key := fastBuildKey(userId, "", md5+extend)
	// if not on the cloud, upload file
	if !ok {
		if err = cloud.uploadSimpleFile(localFilePath, key); err != nil {
			return err
		}
	}

	return nil
}

// fastBuildKey use userId, filePath, filename to generate key by Builder
func fastBuildKey(userId string, filePath string, file string) string {
	var key strings.Builder
	key.Write([]byte("user/"))
	if userId != "" {
		key.Write([]byte(userId))
		key.Write([]byte("/"))
	}
	if filePath != "" {
		key.Write([]byte(filePath))
		key.Write([]byte("/"))
	}
	key.Write([]byte(file))
	return key.String()
}

// create new tencent cloud disk
func NewTencentCloudDisk() CloudDisk {
	return &TencentCloudDisk{
		bucketname: os.Getenv("BUCKET_NAME"),
		secretId:   os.Getenv("SECRET_ID"),
		secretKey:  os.Getenv("SECRET_KEY"),
	}
}
