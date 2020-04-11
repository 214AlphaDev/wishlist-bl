package entities

import "errors"

type WishName struct {
	initialized bool
	name        *string
}

func (wn WishName) Initialized() bool {
	return wn.initialized
}

func (wn WishName) String() string {
	return *wn.name
}

func NewWishName(name string) (WishName, error) {

	if len(name) < 5 {
		return WishName{}, errors.New("wish name is too short")
	}

	return WishName{
		initialized: true,
		name:        &name,
	}, nil
}
