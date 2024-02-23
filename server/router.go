package server

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/api"
	"github.com/ChenMiaoQiu/go-cloud-disk/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("user/login", api.UserLogin)
		v1.POST("user/register", api.UserRegiser)

		auth := v1.Group("")
		auth.Use(middleware.JWTAuth())
		{
			auth.GET("user/:id", api.UserInfo)
			auth.GET("user", api.UserMyInfo)

			auth.POST("file", api.CreateFile)
			auth.GET("files", api.GetFilefolderAllFile)

			auth.GET("downloadpath", api.GetDownloadURL)
			auth.GET("uploadpath", api.GetUploadURL)
		}
	}

	return r
}
