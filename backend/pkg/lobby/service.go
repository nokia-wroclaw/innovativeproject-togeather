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

func (s *service) Join(ctx context.Context, lobbyID int, userID int) (*core.Lobby, error) {
	err := s.lobbyStore.Join(ctx, lobbyID, userID)
	if err != nil {
		return nil, err
	}

	return s.lobbyStore.Get(ctx, lobbyID)
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

func (s *service) AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) error {
	return s.lobbyStore.AddToCart(ctx, userID, lobbyID, mealID)
}

func (s *service) DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) error {
	return s.lobbyStore.DelFromCart(ctx, userID, lobbyID, mealID)
}
