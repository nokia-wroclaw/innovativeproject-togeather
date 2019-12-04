package core

import (
	"context"
	"time"
)

type (
	Lobby struct {
		ID int `json:"id" db:"id"`
		RestaurantID int `json:"restaurant_id" db:"restaurant"`
		Owner int `json:"owner" db:"owner"`
		Expires time.Time `json:"expires" db:"expires"`
	}

	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)
	}
)
