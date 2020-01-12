package restaurant

import (
	"context"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	restaurantStore core.RestaurantStore
}

type command struct {
	Command string `json:"command"`
	Food Meal `json:"meal"`
}

type Meal struct {
	ID      	 int	 `json:"id"`
	RestaurantID int	 `json:"restaurant_id,omitempty"`
	Name         string	 `json:"name"`
	Price        float32 `json:"price"`
	Description	 string	 `json:"description"`
}


func NewService(restaurantStore core.RestaurantStore) core.RestaurantService {
	return &service{restaurantStore: restaurantStore}
}

func (s *service) Exists(ctx context.Context, restaurantID int) (bool, error) {
	return false, nil
}

func (s *service) ListRestaurants(ctx context.Context) ([]*core.Restaurant, error) {
	return s.restaurantStore.ListRestaurants(ctx)
}

func (s *service) RestaurantMenu(ctx context.Context, restaurantID int) ([]*core.Meal, error) {
	return s.restaurantStore.RestaurantMenu(ctx, restaurantID)
}

func (s *service) GetRestaurant(ctx context.Context, restaurantID int) (*core.Restaurant, error) {
	return s.restaurantStore.GetRestaurant(ctx, restaurantID)
}
