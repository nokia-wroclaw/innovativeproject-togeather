package lobby

import (
	"context"
	"time"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	lobbyStore core.LobbyStore
}

func NewService(lobbyStore core.LobbyStore) core.LobbyService {
	return &service{lobbyStore: lobbyStore}
}

func (s *service) List(ctx context.Context) ([]*core.Lobby, error) {
	return s.lobbyStore.List(ctx)
}

func (s *service) Create(
	ctx context.Context,
	restaurantID int,
	ownerID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	return s.lobbyStore.Create(ctx, restaurantID, ownerID, expires, address)
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
