package repository

import "github.com/buemura/voting-system/internal/entity"

type CandidateRepository interface {
	FindByID(id string) (*entity.Candidate, error)
}
