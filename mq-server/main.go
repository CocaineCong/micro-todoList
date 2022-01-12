package main

import (
	"mq-server/conf"
	"mq-server/service"
)

func main() {
	conf.Init()

	forever := make(chan bool)
	service.CreateTask()
	<-forever
}
