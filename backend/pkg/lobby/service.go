package lobby

import (
	"context"
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
	ownerName string,
	expires *time.Time,
	address string,
) (*core.Lobby, int, error) {
	l, err := s.lobbyStore.Create(ctx, restaurantID, expires, address)
	if err != nil {
		return nil, 0, err
	}

	u, err := s.userService.Create(ctx, ownerName)
	if err != nil {
		return nil, 0, err
	}

	return l, u.ID, nil
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

func (s *service) Join(ctx context.Context, lobbyID int, userName string) error {
	_, err := s.userService.Create(ctx, userName)
	return err
}

func (s *service) Get(ctx context.Context, lobbyID int) (*core.Lobby, error) {
	return s.lobbyStore.Get(ctx, lobbyID)
}

func (s *service) Clean(ctx context.Context) {
	s.lobbyStore.Clean(ctx)
}

func (s *service) BelongsToLobby(ctx context.Context, userID int, lobbyID int) (bool, error) {
	return s.lobbyStore.BelongsToLobby(ctx, userID, lobbyID)
}

func (s *service) AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) (*core.CartState, error) {
	err := s.lobbyStore.AddToCart(ctx, userID, lobbyID, mealID)
	if err != nil{
		return nil, err
	}

	return s.CollectCartInfo(ctx, userID, lobbyID)
}

func (s *service) DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) (*core.CartState, error) {
	err := s.lobbyStore.DelFromCart(ctx, userID, lobbyID, mealID)
	if err != nil{
		return nil, err
	}

	return s.CollectCartInfo(ctx, userID, lobbyID)
}

func (s *service) CollectCartInfo(ctx context.Context, userID int, lobbyID int) (*core.CartState, error) {
	return s.lobbyStore.CollectCartInfo(ctx, userID, lobbyID)
}
