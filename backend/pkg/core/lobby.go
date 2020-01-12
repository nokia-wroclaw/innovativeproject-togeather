package core

import (
	"context"
	"time"
)

type (
	Lobby struct {
		ID			int			`json:"id" db:"id"`
		Restaurant 	*Restaurant `json:"restaurant"`
		Owner 		*Client		`json:"owner"`
		Expires 	time.Time	`json:"expires" db:"expires"`
		Location 	*Location	`json:"location"`
		Order 		[]*Item		`json:"order,omitempty"`
	}

	Location struct {
		GeoLat float64 `json:"lat" db:"geolat"`
		GeoLon float64 `json:"lon" db:"geolon"`
		Address string `json:"lobby_address" db:"address"`
	}

	Item struct {
		MealID   int `json:"meal_id, required"`
		Quantity int `json:"quantity, required"`
	}

	Client struct {
		ID 		int `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}


	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerName string,
			expires *time.Time,
			address string,
			order []*Item,
		) (*Lobby, error)

		Edit(
			ctx context.Context,
			lobbyID int,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)

		Join(ctx context.Context, lobbyID int, clientName string)(*User, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerName string,
			expires *time.Time,
			address string,
			order []*Item,
		) (*Lobby, error)

		Edit(
			ctx context.Context,
			lobbyID int,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)

		Join(ctx context.Context, lobbyID int, clientName string)(*User, error)
	}
)
