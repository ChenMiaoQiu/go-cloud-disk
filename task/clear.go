package task

import (
	"os"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

func DeleteLastDayFile() error {
	uploadDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	dst := utils.FastBuildString("./user/", uploadDay)
	err := os.Remove(dst)
	return err
}
