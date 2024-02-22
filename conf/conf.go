package conf

import (
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/joho/godotenv"
)

func Init() {
	// get env
	godotenv.Load()
	// set cloud disk
	utils.SetBaseCloudDisk(os.Getenv("CLOUD_DISK_VERSION"))

	//connect database
	model.Database(os.Getenv("MYSQL_DSN"))
}
