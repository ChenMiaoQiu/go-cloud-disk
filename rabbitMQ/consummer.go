package rabbitMQ

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumerMessage(ctx context.Context, queueName string) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logger.Log().Error("[ConsumerMessage] Failed to open a channel: ", err)
		return nil, err
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	// mq balance
	err = ch.Qos(1, 0, false)
	if err != nil {
		logger.Log().Error("[ConsumerMessage] Failed to set Qos: ", err)
		return nil, err
	}
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
