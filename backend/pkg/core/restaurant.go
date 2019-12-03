package core

import (
	"context"
)

type (
	Restaurant struct {
		ID    int  `json:"id"`
		Name  string  `json:"name"`
		Menu  []Meal    `json:"menu"`
		Address  string `json:"address"`
	}

	Meal struct {
		ID      	 int  `json:"id"`
		RestaurantID int  `json:"restaurant_id"`
		Name         string  `json:"meal"`
		Price        float32  `json:"price"`
	}

	RestaurantService interface {
		Exists(ctx context.Context, restaurantID int) (bool, error)
	}

	RestaurantStore interface {
		Exists(ctx context.Context, restaurantID int) (bool, error)
	}
)