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
