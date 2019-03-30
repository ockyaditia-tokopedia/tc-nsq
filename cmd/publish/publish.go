package main

import (
	"github.com/tkp-junnotantra/tc-nsq/messaging"
)

func main() {
	// initiate producer
	prodConf := messaging.ProducerConfig{
		NsqdAddress: "devel-go.tkpd:4150", // TODO: update to nsqd address
	}
	prod := messaging.NewProducer(prodConf)

	// publish message
	topic := "nsq_0319" // TODO: update to given topic name
	msg := "Hajar bre"  // TODO: write your message here
	prod.Publish(topic, msg)
}
