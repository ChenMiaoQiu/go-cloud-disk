package task

import (
	"time"

	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
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
		loglog.Log().Error("%s error: %dms\n err:%v", jobName, (to-from)/int64(time.Millisecond), err)
	} else {
		loglog.Log().Info("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

// CronJob start cron job
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	// every day restart dailyrank in 0:0:0
	if _, err := Cron.AddFunc("@daily", func() { Run("restart daily rank", RestartDailyRank) }); err != nil {
		loglog.Log().Error("set restart daily rank func err", err)
	}
	// every day delete last day file in 1:0:0
	if _, err := Cron.AddFunc("0 1 * * *", func() { Run("delete last day file", DeleteLastDayFile) }); err != nil {
		loglog.Log().Error("set delete last day file func err", err)
	}
	Cron.Start()
}
