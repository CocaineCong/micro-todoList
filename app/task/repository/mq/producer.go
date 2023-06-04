package mq

import (
	"github.com/streadway/amqp"
)

// SendMessage2MQ 发送消息到mq
func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}

	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}

	return
}
