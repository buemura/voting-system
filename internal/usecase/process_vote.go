package usecase

import (
	"github.com/buemura/voting-system/internal/entity"
	"github.com/buemura/voting-system/internal/repository"
)

type ProcessVote struct {
	candidateRepo repository.CandidateRepository
	voteRepo      repository.VoteRepository
}

func NewProcessVote(
	candidateRepo repository.CandidateRepository,
	voteRepo repository.VoteRepository,
) *ProcessVote {
	return &ProcessVote{
		candidateRepo: candidateRepo,
		voteRepo:      voteRepo,
	}
}

func (u *ProcessVote) Execute(in *entity.CreateVote) (*entity.Vote, error) {
	_, err := u.candidateRepo.FindByID(in.CandidateID)
	if err != nil {
		return nil, err
	}

	vote, err := entity.NewVote(in)
	if err != nil {
		return nil, err
	}

	_, err = u.voteRepo.Create(vote)
	if err != nil {
		return nil, err
	}

	return vote, nil
}
