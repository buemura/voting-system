package main

import (
	"log"

	"github.com/buemura/voting-system/internal/config"
	"github.com/buemura/voting-system/internal/database"
	"github.com/buemura/voting-system/internal/event"
	"github.com/buemura/voting-system/pkg/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	conn, ch := queue.Connect(config.BROKER_URL)
	defer func() {
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close channel: %v", err)
		}
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}()

	if err := queue.DeclareQueue(ch, queue.VOTE_REQUESTED_QUEUE); err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
	if err := queue.DeclareQueue(ch, queue.VOTE_REQUESTED_DLQ); err != nil {
		log.Fatalf("Failed to declare DLQ: %v", err)
	}

	msgs := make(chan amqp.Delivery)

	go queue.Consume(ch, msgs, queue.VOTE_REQUESTED_QUEUE)

	for msg := range msgs {
		event.EventHandler(ch, msg)
	}

	select {}
}
