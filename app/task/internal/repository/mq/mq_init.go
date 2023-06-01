package mq

import (
	"strings"

	"github.com/streadway/amqp"

	"github.com/CocaineCong/micro-todoList/config"
)

var MQ *amqp.Connection

func RabbitMQ() {
	connString := strings.Join([]string{config.RabbitMQ, "://", config.RabbitMQUser, ":", config.RabbitMQPassWord, "@", config.RabbitMQHost, ":", config.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
