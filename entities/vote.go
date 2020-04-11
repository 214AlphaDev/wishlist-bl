package entities

import (
	"github.com/214alphadev/wishlist-bl/utils"
	. "github.com/214alphadev/wishlist-bl/value_objects"
)

type Vote struct {
	initialized bool
	wishID      WishID
	memberID    MemberID
}

func (v Vote) Initialized() bool {
	return v.initialized
}

func (v Vote) MemberID() MemberID {
	return v.memberID
}

func (v Vote) WishID() WishID {
	return v.wishID
}

func NewVote(memberID MemberID, wishID WishID) (Vote, error) {

	if err := utils.Initialized(memberID, wishID); err != nil {
		return Vote{}, err
	}

	return Vote{
		initialized: true,
		wishID:      wishID,
		memberID:    memberID,
	}, nil

}
