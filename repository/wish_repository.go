package repository

import (
	. "github.com/214alphadev/wishlist-bl/entities"
)

type IWishRepository interface {
	Save(wish Wish) error
	Get(wishID WishID) (*Wish, error)
	Query(from *WishID, next uint32) ([]Wish, error)
	VotesOf(wishID WishID) (uint32, error)
}
