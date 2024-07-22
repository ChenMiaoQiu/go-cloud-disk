package rabbitMQ

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessageToMQ(ctx context.Context, queueName string, body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logger.Log().Error("[SendMessageToMQ] Failed to open a channel: %s", err)
		return
	}

	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		logger.Log().Error("[SendMessageToMQ] Failed to publish a message: %s", err)
		return
	}
	return
}
