package serializer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response base serializer
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// Response Url link
type ResponseUrl struct {
	Url string `json:"url"`
}

const (
	// CodeNotLogin success code 200
	CodeSuccess = http.StatusOK
	// CodeNotLogin Not Login Code 1250
	CodeNotLogin = 1250
	// CodeNotRightErr Unauthorized Code 401
	CodeNotAuthError = http.StatusUnauthorized
	// CodeDBError database error Code 500
	CodeDBError = http.StatusInternalServerError
	// InternalError error Code 500
	CodeInternalError = http.StatusInternalServerError
	// CodeError common error code 404
	CodeError = http.StatusNotFound
	// CodeParamsError params error Code 50001
	CodeParamsError = 50001
)

// Success return a success response
func Success(data interface{}) Response {
	return Response{
		Code: CodeSuccess,
		Msg:  "Success",
		Data: data,
	}
}

// NotAuthErr use msg build a not auth err response, if msg is null
// msg info is "NotAuth"
func NotAuthErr(msg string) Response {
	if msg == "" {
		msg = "NotAuth"
	}
	return Response{
		Code: CodeNotAuthError,
		Msg:  msg,
	}
}

// NotLogin return an unlogin response
func NotLogin(msg string) Response {
	if msg == "" {
		msg = "NotLogin"
	}
	return Response{
		Code: CodeNotLogin,
		Msg:  msg,
	}
}

// Err return a common error response
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = fmt.Sprintf("%+v", err)
	}
	return res
}

// DBErr return a database error response
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "DBerr"
	}
	return Err(CodeDBError, msg, err)
}

// InternalErr return an Internal err response
func InternalErr(msg string, err error) Response {
	if msg == "" {
		msg = "Internal"
	}
	return Err(CodeInternalError, msg, err)
}

// DBErr return a params error response
func ParamsErr(msg string, err error) Response {
	if msg == "" {
		msg = "ParamErr"
	}
	return Err(CodeParamsError, msg, err)
}

// ErrorResponse return err msg
func ErrorResponse(err error) Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ParamsErr("JsonNotMatched", err)
	}

	return ParamsErr("", err)
}
