package lobby

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	lobbyStore core.LobbyStore
	userService core.UserService
}

func NewService(lobbyStore core.LobbyStore, userService core.UserService) core.LobbyService {
	return &service{lobbyStore: lobbyStore, userService: userService}
}

func (s *service) List(ctx context.Context) ([]*core.Lobby, error) {
	s.Clean(ctx)
	return s.lobbyStore.List(ctx)
}

func (s *service) Create(
	ctx context.Context,
	restaurantID int,
	userID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	return s.lobbyStore.Create(ctx, restaurantID, userID, expires, address)
}

func (s *service) Edit(
	ctx context.Context,
	lobbyID int,
	restaurantID int,
	ownerID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	return s.lobbyStore.Edit(ctx, lobbyID, restaurantID, ownerID, expires, address)
}

func (s *service) Join(ctx context.Context, lobbyID int, clientID int) (*core.Lobby, error) {
	belongs, err := s.BelongsToLobby(ctx, clientID, lobbyID)
	if err != nil{
		return nil, err
	}

	if belongs{
		return nil, errors.New("join lobby: user already belongs to this lobby")
	}

	err = s.lobbyStore.Join(ctx, lobbyID, clientID)
	if err != nil {
		return nil, err
	}

	return s.lobbyStore.Get(ctx, lobbyID)
}

func (s *service) Get(ctx context.Context, userID int, lobbyID int) (*core.Lobby, error) {
	belongs, err := s.BelongsToLobby(ctx, userID, lobbyID)
	if err != nil{
		return nil, err
	}

	if !belongs{
		return nil, errors.New("get lobby: user not belongs to this lobby")
	}

	return s.lobbyStore.Get(ctx, lobbyID)
}

func (s *service) Clean(ctx context.Context) {
	s.lobbyStore.Clean(ctx)
}

func (s *service) BelongsToLobby(ctx context.Context, userID int, lobbyID int) (bool, error) {
	return s.lobbyStore.BelongsToLobby(ctx, userID, lobbyID)
}

func (s *service) AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) (*core.CartState, error) {
	err := s.lobbyStore.Available(ctx, lobbyID)
	if err != nil {
		return nil, fmt.Errorf("error adding item to cart: %w", err)
	}

	err = s.lobbyStore.AddToCart(ctx, userID, lobbyID, mealID)
	if err != nil{
		return nil, err
	}

	return s.CollectCartInfo(ctx, userID, lobbyID)
}

func (s *service) DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) (*core.CartState, error) {
	err := s.lobbyStore.Available(ctx, lobbyID)
	if err != nil {
		return nil, fmt.Errorf("error deleting item from cart: %w", err)
	}

	err = s.lobbyStore.DelFromCart(ctx, userID, lobbyID, mealID)
	if err != nil{
		return nil, err
	}

	return s.CollectCartInfo(ctx, userID, lobbyID)
}

func (s *service) CollectCartInfo(ctx context.Context, userID int, lobbyID int) (*core.CartState, error) {
	return s.lobbyStore.CollectCartInfo(ctx, userID, lobbyID)
}
