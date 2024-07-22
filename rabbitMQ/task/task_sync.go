package task

import (
	"context"
	"encoding/json"

	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type SendConfirmEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func RunSendConfirmEmail(ctx context.Context) error {
	msgs, err := rabbitMQ.ConsumerMessage(ctx, rabbitMQ.RabbitMqSendEmailQueue)
	if err != nil {
		return err
	}
	var forever chan struct{}

	go func() {
		for msg := range msgs {
			logger.Log().Info("[RunSendConfirmEmail] Received message: ", string(msg.Body))

			sendConirmEmailReq := SendConfirmEmailRequest{}
			err = json.Unmarshal(msg.Body, &sendConirmEmailReq)
			if err != nil {
				logger.Log().Error("[RunSendConfirmEmail] Unmarshal message error: ", err)
			}

			err = utils.SendConfirmMessage(sendConirmEmailReq.Email, sendConirmEmailReq.Code)
			if err != nil {
				logger.Log().Error("[RunSendConfirmEmail] Send confirm message error: ", err)
			}

			msg.Ack(false)
		}
	}()

	logger.Log().Info("RunSendConfirmEmail service started")
	<-forever
	return nil
}
