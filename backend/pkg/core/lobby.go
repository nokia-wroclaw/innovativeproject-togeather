package core

import "context"

type (
	Lobby struct {
		ID int `json:"id" db:"id"`
		RestaurantID int `json:"restaurant_id" db:"restaurant"`
		Owner int `json:"owner" db:"owner"`
	}

	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)
	}
)
