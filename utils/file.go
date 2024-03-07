package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path"
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

// GetFileMD5 get file md5 code
func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	// add extend name to md5 calculate because md5 caculate will get same
	// md5 code in file have same content and diff extented name
	ext := path.Ext(file.Name())
	hash.Write([]byte(ext))
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// splitFilename split file.filename to filename and extend name
func SplitFilename(str string) (filename string, extend string) {
	for i := len(str) - 1; i >= 0 && str[i] != '/'; i-- {
		if str[i] == '.' {
			if i != 0 {
				filename = str[:i]
			}
			if i != len(str)-1 {
				extend = str[i+1:]
			}
			return
		}
	}
	return str, ""
}
