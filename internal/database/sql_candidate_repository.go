package database

import (
	"database/sql"

	"github.com/buemura/voting-system/internal/entity"
)

type SqlCandidateRepository struct {
	db *sql.DB
}

func NewSqlCandidateRepository() *SqlCandidateRepository {
	return &SqlCandidateRepository{db: Conn}
}

func (r *SqlCandidateRepository) FindByID(id string) (*entity.Candidate, error) {
	var candidate entity.Candidate
	err := r.db.QueryRow(`SELECT id, name FROM "candidate" WHERE id = $1`, id).Scan(&candidate.ID, &candidate.Name)
	if err != nil {
		return nil, err
	}

	return &candidate, nil
}
