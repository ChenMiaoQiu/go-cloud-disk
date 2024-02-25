package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

func FastBuildFileName(fileName string, filePostfix string) string {
	var res strings.Builder
	res.Write([]byte(fileName))
	res.Write([]byte("."))
	res.Write([]byte(filePostfix))
	return res.String()
}

func FastBuildString(str ...string) string {
	var res strings.Builder
	for _, s := range str {
		res.Write([]byte(s))
	}
	return res.String()
}

func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
