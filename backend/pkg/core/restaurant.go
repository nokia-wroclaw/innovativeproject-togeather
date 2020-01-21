package core

import (
	"context"
)

type (
	Restaurant struct {
		ID		int		`json:"id"`
		Name	string	`json:"name"`
		Menu	[]*Meal	`json:"menu,omitempty"`
		Address	string	`json:"address,omitempty"`
		Delivery float32 `json:"delivery_cost"`
	}

	Meal struct {
		ID      	 int	 `json:"id"`
		RestaurantID int	 `json:"restaurant_id,omitempty"`
		Name         string	 `json:"name"`
		Price        float32 `json:"price"`
		Description	 string	 `json:"description"`
	}

	RestaurantService interface {
		List(ctx context.Context) ([]*Restaurant, error)
		GetMenu(ctx context.Context, restaurantID int) ([]*Meal, error)
		Get(ctx context.Context, restaurantID int) (*Restaurant, error)
	}

	RestaurantStore interface {
		List(ctx context.Context) ([]*Restaurant, error)
		GetMenu(ctx context.Context, restaurantID int) ([]*Meal, error)
		Get(ctx context.Context, restaurantID int) (*Restaurant, error)
	}
)