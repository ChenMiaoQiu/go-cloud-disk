package rabbitmq

import (
	"fmt"
	"strings"

	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMq *amqp.Connection

func InitRabbitMq() {
	connString := strings.Join([]string{conf.RabbitMQ, "://", conf.RabbitMQUser, ":", conf.RabbitMQPassword, "@", conf.RabbitMQHost, ":", conf.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
	}
	RabbitMq = conn
}
