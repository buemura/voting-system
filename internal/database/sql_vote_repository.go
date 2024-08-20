package database

import (
	"database/sql"

	"github.com/buemura/voting-system/internal/entity"
)

type SqlVoteRepository struct {
	db *sql.DB
}

func NewSqlVoteRepository() *SqlVoteRepository {
	return &SqlVoteRepository{db: Conn}
}

func (r *SqlVoteRepository) Create(in *entity.Vote) (*entity.Vote, error) {
	_, err := r.db.Exec(
		`INSERT INTO "vote" (id, candidate_id, created_at) VALUES ($1, $2, $3)`,
		in.ID, in.CandidateID, in.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return in, nil
}
