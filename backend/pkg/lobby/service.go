package lobby

import (
	"context"

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
