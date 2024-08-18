package entity

import (
	"crypto/rand"
	"time"

	"github.com/lucsky/cuid"
)

type CreateVote struct {
	CandidateID string `json:"candidate_id"`
}

type Vote struct {
	ID          string
	CandidateID string    `json:"candidate_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewVote(in *CreateVote) (*Vote, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Vote{
		ID:          cuid,
		CandidateID: in.CandidateID,
		CreatedAt:   time.Now(),
	}, nil
}
