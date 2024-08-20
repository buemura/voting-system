package repository

import "github.com/buemura/voting-system/internal/entity"

type VoteRepository interface {
	Create(*entity.Vote) (*entity.Vote, error)
}
