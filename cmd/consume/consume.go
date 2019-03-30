package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/tkp-junnotantra/tc-nsq/messaging"
)

const (
	defaultConsumerMaxAttempts = 10
	defaultConsumerMaxInFlight = 100
)

func main() {
	// initiate consumer
	cfg := messaging.ConsumerConfig{
		Channel:       "ocky",               // TODO: update to given channel name
		LookupAddress: "devel-go.tkpd:4161", // TODO: update to nsqlookups adress
		Topic:         "nsq_0319",           // TODO: update to given topic name
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       handleMessage,
	}
	consumer := messaging.NewConsumer(cfg)

	// run consumer
	consumer.Run()

	// keep app alive until terminated
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}

func handleMessage(message *nsq.Message) error {
	// TODO: print and finish message
	log.Println(string(message.Body))
	message.Finish()
	return nil
}

func requeueMessage(message *nsq.Message) error {
	// TODO: requeue message
	log.Println(string(message.Body))
	message.Requeue(3 * time.Second)
	return errors.New("Fail")
}
