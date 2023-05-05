package main

import (
	"log"

	rabbitmqclient "dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface.git/rabbitmqclient"
	yifeilogger "github.com/YifeWu/gologger"
)

func main() {
	log.Println("hello")
	yifeilogger.LogInfo("github module test")
	// loggershell.LogInfo("This is very good news!")
	rabbitmqclient.LogRabbitMQ("This is fantascis!")
}
