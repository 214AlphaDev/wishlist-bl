package entities

import "github.com/satori/go.uuid"

type WishID struct {
	initialized bool
	id          *uuid.UUID
}

func (wi WishID) Initialized() bool {
	return wi.initialized
}

func (wi WishID) String() string {
	return wi.id.String()
}

func NewWishID(wishID string) (WishID, error) {

	id, err := uuid.FromString(wishID)
	if err != nil {
		return WishID{}, err
	}

	return WishID{
		id:          &id,
		initialized: true,
	}, nil

}
