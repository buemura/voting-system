package event

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"

	"github.com/buemura/voting-system/internal/entity"
	"github.com/buemura/voting-system/pkg/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TransactionEventHandler(ch *amqp.Channel, msg amqp.Delivery) {
	switch msg.RoutingKey {
	case queue.VOTE_REQUESTED_QUEUE:
		// Parse message body
		var in *entity.CreateVote
		err := json.Unmarshal([]byte(msg.Body), &in)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// call Usecase

		if err != nil {
			slog.Error(err.Error())

			// TODO: adds retry stategy before sending it to DLQ
			err = queue.PublishToQueue(ch, msg.Body, queue.VOTE_REQUESTED_DLQ)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to send message to DLQ queue: %s", err))
			}
		}
	}
}
