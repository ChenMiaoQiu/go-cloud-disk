package conf

import (
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/joho/godotenv"
)

func Init() {
	// get env
	godotenv.Load()

	//connect database
	model.Database(os.Getenv("MYSQL_DSN"))
}
