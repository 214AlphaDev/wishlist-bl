package entities

import (
	"github.com/214alphadev/wishlist-bl/utils"
	. "github.com/214alphadev/wishlist-bl/value_objects"
)

type Wish struct {
	initialized bool
	id          *WishID
	name        *WishName
	description *WishDescription
	story       *WishStory
	creator     *MemberID
	category    *Category
}

func (w Wish) Initialized() bool {
	return w.initialized
}

func (w Wish) ID() WishID {
	return *w.id
}

func (w Wish) Name() WishName {
	return *w.name
}

func (w Wish) Description() WishDescription {
	return *w.description
}

func (w Wish) Story() WishStory {
	return *w.story
}

func (w Wish) Creator() MemberID {
	return *w.creator
}

func (w Wish) Category() Category {
	return *w.category
}

func NewWish(creator MemberID, id WishID, name WishName, description WishDescription, story WishStory, category Category) (Wish, error) {

	if err := utils.Initialized(creator, id, name, description, story); err != nil {
		return Wish{}, err
	}

	return Wish{
		id:          &id,
		name:        &name,
		description: &description,
		story:       &story,
		initialized: true,
		creator:     &creator,
		category:    &category,
	}, nil

}
