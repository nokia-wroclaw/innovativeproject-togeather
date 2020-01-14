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
	s.Clean(ctx)
	return s.lobbyStore.List(ctx)
}

func (s *service) Create(
	ctx context.Context,
	restaurantID int,
	ownerName string,
	expires *time.Time,
	address string,
	order []*core.Item,
) (*core.Lobby, error) {
	return s.lobbyStore.Create(ctx, restaurantID, ownerName, expires, address, order)
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

func (s *service) Join(ctx context.Context, lobbyID int, clientName string)(*core.User, error){
	return s.lobbyStore.Join(ctx, lobbyID, clientName)
}

func (s *service) Get(ctx context.Context, lobbyID int) (*core.Lobby, error) {
	return s.lobbyStore.Get(ctx, lobbyID)

func (s *service) Clean(ctx context.Context) {
	s.lobbyStore.Clean(ctx)
}
