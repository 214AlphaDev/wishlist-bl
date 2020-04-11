package entities

import "fmt"

type WishStory struct {
	initialized bool
	story       *string
}

func (ws WishStory) Initialized() bool {
	return ws.initialized
}

func (ws WishStory) String() string {
	return *ws.story
}

func (ws WishStory) IsNil() bool {
	return ws.story == nil
}

func NewWishStory(story *string) (WishStory, error) {

	switch story {
	case nil:
		return WishStory{initialized: true}, nil
	default:

		if len(*story) < 20 {
			return WishStory{}, fmt.Errorf("wish story is too short")
		}

		return WishStory{initialized: true, story: &*story}, nil

	}

}
