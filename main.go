package main

import (
	"log"

	rabbitmqclient "dev.azure.com/slb1-swt/wireline-automation/_git/lib-go.git/rabbitmqclient"
)

type MySubscriber struct {
	Topic string //Todo: This property is not needed, to be removed.
}

/*
Called the received message for the topic
*/
func (sub MySubscriber) HandleMessage(body []byte, topic string, replyToTopic string) {
	log.Printf(" [x] %s", topic)
}

func main() {

	log.Printf("Hello PDP Express")
	rabbitmqclient.Greeting()

	connMngr := rabbitmqclient.CreateConnectionManager()

	// sub := MySubscriber{Topic: "maxwell.job.status.state.v1"}
	// connMngr.SubscribeExchange("pubsub", "maxwell.job.status.state.v1", sub)

	testPubSub := true
	if testPubSub {
		// Subscribe to Maxwell pub/sub topic, then replish to PDP express pub/sub topic
		repubSub := rabbitmqclient.RepublishSubscriber{RepublishTopic: "pdp.express.depth.tension.state.v1", ConnManager: connMngr}
		connMngr.Subscribe("maxwell.job.status.state.v1", repubSub)
	}

	testRpc := false
	if testRpc {
		// Subscribe to Maxwell RPC topic, then replish to PDP express RPC topic
		repubRpc := rabbitmqclient.RepublishSubscriber{RepublishTopic: "pdp.express.job_summary.get.v1", ConnManager: connMngr}
		msg := rabbitmqclient.BuildPayload()
		connMngr.PublishRequest("maxwell.job_summary.get.v1", msg, repubRpc)
	}

	var forever chan struct{}
	log.Printf(" [*] Waiting ")
	<-forever
}
