package restaurant

import (
	"context"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	restaurantStore core.RestaurantStore
}


func NewService(restaurantStore core.RestaurantStore) core.RestaurantService {
	return &service{restaurantStore: restaurantStore}
}

func (s *service) Exists(ctx context.Context, restaurantID int) (bool, error) {
	return false, nil
}

func (s *service) List(ctx context.Context) ([]*core.Restaurant, error) {
	return s.restaurantStore.List(ctx)
}

func (s *service) GetMenu(ctx context.Context, restaurantID int) ([]*core.Meal, error) {
	return s.restaurantStore.GetMenu(ctx, restaurantID)
}

func (s *service) Get(ctx context.Context, restaurantID int) (*core.Restaurant, error) {
	m, err := s.GetMenu(ctx, restaurantID)
	if err != nil{
		return nil, err
	}

	r, err := s.restaurantStore.Get(ctx, restaurantID)
	if err != nil{
		return nil, err
	}

	r.Menu = m
	return r, nil
}
