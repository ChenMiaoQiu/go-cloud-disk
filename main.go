package main

import (
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
	"github.com/ChenMiaoQiu/go-cloud-disk/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// conf init
	conf.Init()

	// set router
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	r.Run(":" + os.Getenv("SERVER_PORT"))
}
