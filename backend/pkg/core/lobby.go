package core

import (
	"context"
	"time"
)

type (
	Lobby struct {
		ID			int			`json:"id" db:"id"`
		Restaurant 	*Restaurant `json:"restaurant,omitempty"`
		Expires 	*time.Time	`json:"expires,omitempty" db:"expires"`
		Location 	*Location	`json:"location,omitempty"`
	}

	Location struct {
		GeoLat float64 `json:"lat" db:"geolat"`
		GeoLon float64 `json:"lon" db:"geolon"`
		Address string `json:"lobby_address" db:"address"`
	}

	CartState struct {
		Meals			[]*Meal	`json:"meals"`
		CartValue		float32	`json:"cart_value"`
		DeliveryCost	float32	`json:"delivery_cost"`
		LobbyCount		int		`json:"lobby_count"`
	}


LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerName string,
			expires *time.Time,
			address string,
		) (*Lobby, int, error)

		Edit(
			ctx context.Context,
			lobbyID int,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)

		Join(ctx context.Context, lobbyID int, userName string) error

		Get(ctx context.Context, lobbyID int) (*Lobby, error)

		Clean(ctx context.Context)

		BelongsToLobby(ctx context.Context, userID int, lobbyID int) (bool, error)

		AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) (*CartState, error)

		DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) (*CartState, error)

		CollectCartInfo(ctx context.Context, userID int, lobbyID int) (*CartState, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)

		Edit(
			ctx context.Context,
			lobbyID int,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)

		Join(ctx context.Context) error

		Get(ctx context.Context, lobbyID int)(*Lobby, error)

		Clean(ctx context.Context)

		BelongsToLobby(ctx context.Context, userID int, lobbyID int)(bool, error)

		AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) error

		DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) error

		CollectCartInfo(ctx context.Context, userID int, lobbyID int) (*CartState, error)
	}
)
