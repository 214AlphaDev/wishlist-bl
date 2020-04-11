package services

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	. "github.com/214alphadev/wishlist-bl/entities"
	. "github.com/214alphadev/wishlist-bl/repository"
	"github.com/214alphadev/wishlist-bl/utils"
	. "github.com/214alphadev/wishlist-bl/value_objects"
)

type IWishService interface {
	Create(memberID MemberID, name WishName, description WishDescription, story WishStory, category Category) (WishID, error)
	Vote(wishID WishID, memberID MemberID) error
	WithdrawVote(wishID WishID, memberID MemberID) error
	Wishes(start *WishID, next uint32) ([]Wish, error)
	GetByID(wishID WishID) (*Wish, error)
	VotesOfWish(wishID WishID) (uint32, error)
}

type wishService struct {
	wishRepository IWishRepository
	voteRepository IVoteRepository
	voteService    IVoteService
}

func (ws *wishService) Create(memberID MemberID, name WishName, description WishDescription, story WishStory, category Category) (WishID, error) {

	if err := utils.Initialized(memberID, name, description, story); err != nil {
		return WishID{}, err
	}

	id, err := NewWishID(uuid.NewV4().String())
	if err != nil {
		return WishID{}, err
	}

	wish, err := NewWish(memberID, id, name, description, story, category)
	if err != nil {
		return WishID{}, err
	}

	if err := ws.wishRepository.Save(wish); err != nil {
		return WishID{}, err
	}

	return id, nil

}

func (ws *wishService) Vote(wishID WishID, memberID MemberID) error {

	if err := utils.Initialized(wishID, memberID); err != nil {
		return err
	}

	votesLeft, err := ws.voteService.VotesLeft(memberID)
	if err != nil {
		return err
	}
	if votesLeft <= 0 {
		return errors.New("NoVotesLeft")
	}

	fetchedWish, err := ws.wishRepository.Get(wishID)
	if err != nil {
		return err
	}
	if fetchedWish == nil {
		return fmt.Errorf("WishDoesNotExist")
	}

	vote, err := NewVote(memberID, wishID)
	if err != nil {
		return err
	}

	exists, err := ws.voteRepository.DoesExist(vote)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("AlreadyVoted")
	}

	return ws.voteRepository.Save(vote)

}

func (ws *wishService) WithdrawVote(wishID WishID, memberID MemberID) error {

	if err := utils.Initialized(wishID, memberID); err != nil {
		return err
	}

	fetchedWish, err := ws.wishRepository.Get(wishID)
	if err != nil {
		return err
	}
	if fetchedWish == nil {
		return errors.New("WishDoesNotExist")
	}

	vote, err := NewVote(memberID, wishID)
	if err != nil {
		return err
	}

	exists, err := ws.voteRepository.DoesExist(vote)
	if err != nil {
		return err
	}

	switch exists {
	case true:
		return ws.voteRepository.Delete(vote)
	default:
		return errors.New("NeverVotedOnWish")
	}

}

func (ws *wishService) Wishes(start *WishID, next uint32) ([]Wish, error) {

	if err := utils.Initialized(start); err != nil {
		return nil, err
	}

	return ws.wishRepository.Query(start, next)

}

func (ws *wishService) GetByID(wishID WishID) (*Wish, error) {

	w, err := ws.wishRepository.Get(wishID)

	if err != nil {
		return nil, err
	}

	return w, nil

}

func (ws *wishService) VotesOfWish(wishID WishID) (uint32, error) {

	if err := utils.Initialized(wishID); err != nil {
		return 0, err
	}

	return ws.wishRepository.VotesOf(wishID)

}

func NewWishService(wishRepository IWishRepository, voteRepository IVoteRepository, voteService IVoteService) IWishService {
	return &wishService{
		wishRepository: wishRepository,
		voteRepository: voteRepository,
		voteService:    voteService,
	}
}
