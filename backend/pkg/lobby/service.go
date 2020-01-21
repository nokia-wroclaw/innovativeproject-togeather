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

	u, err := s.userService.Create(ctx, ownerName, l.ID, true)
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

func (s *service) Join(ctx context.Context, lobbyID int, userName string)(*core.User, error){
	return s.userService.Create(ctx, userName, lobbyID, false)
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
