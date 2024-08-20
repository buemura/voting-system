package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/buemura/voting-system/internal/config"
	"github.com/buemura/voting-system/internal/entity"
	"github.com/buemura/voting-system/pkg/queue"
	"github.com/go-chi/chi/v5"
)

func handleRequestError(w http.ResponseWriter, status int, err error, detailedErr string) {
	if len(detailedErr) > 0 {
		slog.Error(detailedErr)
	}
	slog.Error("Error: " + err.Error())
	http.Error(w, err.Error(), status)
}

func RegisterRoutes(mux *chi.Mux) http.Handler {
	mux.Post("/vote", createVote)
	return mux
}

func createVote(w http.ResponseWriter, r *http.Request) {
	slog.Info("[Handler][CreateVote] - Incoming request")
	var input entity.CreateVote
	b, err := io.ReadAll(r.Body)
	if err != nil {
		handleRequestError(w, http.StatusBadRequest, err, "")
		return
	}

	if err := json.Unmarshal(b, &input); err != nil {
		handleRequestError(w, http.StatusBadRequest, err, "")
		return
	}

	_, ch := queue.Connect(config.BROKER_URL)
	err = queue.Publish(ch, string(b), queue.VOTE_REQUESTED_QUEUE)
	if err != nil {
		slog.Info(fmt.Sprintf("Failed to publish message to queue: %s", queue.VOTE_REQUESTED_QUEUE))
		handleRequestError(w, http.StatusInternalServerError, err, "")
		return
	}

	slog.Info(fmt.Sprintf("Published message to queue: %s", queue.VOTE_REQUESTED_QUEUE))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(input)
}
