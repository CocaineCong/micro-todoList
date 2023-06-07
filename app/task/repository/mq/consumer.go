package mq

import (
	"context"

	"github.com/streadway/amqp"
)

// ConsumeMessage MQ到mysql
func ConsumeMessage(ctx context.Context, queueName string) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
