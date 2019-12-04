package core

import "context"

type (
	Lobby struct {

	}

	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)
	}
)
