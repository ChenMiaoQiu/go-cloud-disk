package task

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

type jobFunc func() error

// Run runing job and print result that job executed
func Run(jobName string, job jobFunc) {
	// caculate job executed time
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	if err != nil {
		log.Println(err)
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

// CronJob start cron job
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	// every day restart dailyrank in 0:0:0
	Cron.AddFunc("0 0 0 * * *", func() { Run("restart daily rank", RestartDailyRank) })
	// every day delete last day file in 1:0:0
	Cron.AddFunc("0 0 1 * * *", func() { Run("restart delete last day file", DeleteLastDayFile) })
	Cron.Start()

	fmt.Println("cron job start")
}
