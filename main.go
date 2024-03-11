package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// gin gracefully shuts down the server
	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}
	go func() {
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
