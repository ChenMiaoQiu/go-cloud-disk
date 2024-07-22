package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/auth"
	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ"
	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ/script"
	"github.com/ChenMiaoQiu/go-cloud-disk/server"
	"github.com/ChenMiaoQiu/go-cloud-disk/task"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
	"github.com/gin-gonic/gin"
)

// initServer init server that server needed
func initServer() {
	// set cloud disk
	disk.SetBaseCloudDisk()

	// set log
	logger.BuildLogger()

	// connect database
	model.Database()

	// connect redis
	cache.Redis()

	// start regular task
	task.CronJob()

	// start casbin
	auth.InitCasbin()

	// start rabbitmq
	rabbitMQ.InitRabbitMq()
}

func loadingScript() {
	ctx := context.Background()
	go script.SendConfirmEmailSync(ctx)
}

func main() {
	// conf init
	conf.Init()
	initServer()
	loadingScript()

	// set router
	gin.SetMode(conf.GinMode)
	r := server.NewRouter()

	// gin gracefully shuts down the server
	srv := &http.Server{
		Addr:    ":" + conf.ServerPort,
		Handler: r,
	}

	go func() {
		log.Println("go-cloud-disk server start")
		// connect serve
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait system exit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// set exit time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		log.Println("Server exiting")
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
