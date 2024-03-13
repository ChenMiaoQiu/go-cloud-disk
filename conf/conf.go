package conf

import (
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/auth"
	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/disk"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/task"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
	"github.com/joho/godotenv"
)

func Init() {
	// get env
	godotenv.Load()
	// set cloud disk
	disk.SetBaseCloudDisk(os.Getenv("CLOUD_DISK_VERSION"))

	// set log
	loglog.BuildLogger(os.Getenv("LOG_LEVEL"))

	// connect database
	model.Database(os.Getenv("MYSQL_DSN"))

	// connect redis
	cache.Redis()

	// start regular task
	task.CronJob()

	// start casbin
	auth.InitCasbin()
}
