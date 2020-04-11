package repository

import (
	. "github.com/214alphadev/wishlist-bl/entities"
	. "github.com/214alphadev/wishlist-bl/value_objects"
)

type IVoteRepository interface {
	Save(v Vote) error
	DoesExist(v Vote) (bool, error)
	Delete(v Vote) error
	GetVotesOfMember(memberID MemberID) ([]Vote, error)
}
