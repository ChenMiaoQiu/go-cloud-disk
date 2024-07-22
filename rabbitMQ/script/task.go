package script

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ/task"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

func SendConfirmEmailSync(ctx context.Context) {
	err := task.RunSendConfirmEmail(ctx)
	if err != nil {
		logger.Log().Error("SendConfirmEmailSync error: ", err)
	}
}
