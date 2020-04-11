package services

import (
	. "github.com/214alphadev/wishlist-bl/entities"
	. "github.com/214alphadev/wishlist-bl/repository"
	"github.com/214alphadev/wishlist-bl/utils"
	. "github.com/214alphadev/wishlist-bl/value_objects"
)

type IVoteService interface {
	Votes(member MemberID) ([]Vote, error)
	VotesLeft(member MemberID) (uint32, error)
}

type voteService struct {
	voteRepository IVoteRepository
}

func (vs *voteService) Votes(member MemberID) ([]Vote, error) {

	if err := utils.Initialized(member); err != nil {
		return nil, err
	}

	return vs.voteRepository.GetVotesOfMember(member)

}

func (vs *voteService) VotesLeft(member MemberID) (uint32, error) {

	maxVotes := 3

	votes, err := vs.voteRepository.GetVotesOfMember(member)
	if err != nil {
		return 0, err
	}

	if len(votes) > maxVotes {
		return 0, nil
	}

	return uint32(maxVotes - len(votes)), nil

}

func NewVoteService(voteRepository IVoteRepository) IVoteService {
	return &voteService{
		voteRepository: voteRepository,
	}
}
