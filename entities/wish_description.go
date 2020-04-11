package entities

import "errors"

type WishDescription struct {
	initialized bool
	description *string
}

func (wd WishDescription) Initialized() bool {
	return wd.initialized
}

func (wd WishDescription) String() string {
	return *wd.description
}

func NewWishDescription(description string) (WishDescription, error) {

	if len(description) < 30 {
		return WishDescription{}, errors.New("wish description is too short")
	}

	return WishDescription{
		initialized: true,
		description: &description,
	}, nil

}
