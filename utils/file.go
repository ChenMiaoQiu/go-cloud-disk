package utils

import "strings"

func FastBuildFileName(fileName string, filePostfix string) string {
	var res strings.Builder
	res.Write([]byte(fileName))
	res.Write([]byte("."))
	res.Write([]byte(filePostfix))
	return res.String()
}
