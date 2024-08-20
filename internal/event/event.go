package event

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"

	"github.com/buemura/voting-system/internal/database"
	"github.com/buemura/voting-system/internal/entity"
	"github.com/buemura/voting-system/internal/usecase"
	"github.com/buemura/voting-system/pkg/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func EventHandler(ch *amqp.Channel, msg amqp.Delivery) {
	switch msg.RoutingKey {
	case queue.VOTE_REQUESTED_QUEUE:
		slog.Info("[Event][EventHandler] - Incoming event:")
		// Parse message body
		var in *entity.CreateVote
		err := json.Unmarshal([]byte(msg.Body), &in)
		if err != nil {
			log.Fatalf(err.Error())
		}

		slog.Info(fmt.Sprintf("[Event][EventHandler] - Payload: %s", string(msg.Body)))

		// call Usecase
		candidateRepo := database.NewSqlCandidateRepository()
		voteRepo := database.NewSqlVoteRepository()
		uc := usecase.NewProcessVote(candidateRepo, voteRepo)

		_, err = uc.Execute(in)
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
