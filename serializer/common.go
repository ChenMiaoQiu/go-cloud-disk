package serializer

import (
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

const (
	// CodeNotLogin Not Login Code 401
	CodeNotLogin = http.StatusUnauthorized
	// CodeNotRightErr Unauthorized Code 401
	CodeNotRightError = http.StatusUnauthorized
	// CodeDBError database error Code 500
	CodeDBError = http.StatusInternalServerError
	// CodeError common error code 404
	CodeError = http.StatusNotFound
	// CodeParamsError params error Code 50001
	CodeParamsError = 50001
)

// NotLogin return an unlogin response
func NotLogin() Response {
	return Response{
		Code: CodeNotLogin,
		Msg:  "Not Login",
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
		msg = "database error"
	}
	return Err(CodeDBError, msg, err)
}

// DBErr return a params error response
func ParamsErr(msg string, err error) Response {
	if msg == "" {
		msg = "parmars error"
	}
	return Err(CodeParamsError, msg, err)
}
