package server

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/api"
	"github.com/ChenMiaoQiu/go-cloud-disk/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // set upload speed
	r.Use(middleware.Cors())
	r.GET("ping", api.Ping)

	v1 := r.Group("/api/v1")
	{
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/register", api.UserRegiser)

		auth := v1.Group("")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("user/:id", api.UserInfo)
			auth.GET("user", api.UserMyInfo)

			auth.GET("file/:fileid", api.GetDownloadURL)
			auth.POST("file", api.UploadFile)
			auth.DELETE("file/:fileid", api.DeleteFile)

			auth.POST("filefolder", api.CreateFileFolder)
			auth.GET("filefolder/:filefolderid/file", api.GetFilefolderAllFile)
			auth.GET("filefolder/:filefolderid/filefolder", api.GetFilefolderAllFilefolder)
		}
	}

	return r
}
