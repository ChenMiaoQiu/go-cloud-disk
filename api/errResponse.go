package api

import (
	"encoding/json"

	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

// ErrorResponse return err msg
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamsErr("JSON类型不匹配", err)
	}

	return serializer.ParamsErr("参数错误", err)
}
