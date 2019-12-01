package restaurant

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type restaurantStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) core.RestaurantStore {
	return &restaurantStore{db: db}
}

func (s *restaurantStore) Exists(ctx context.Context, restaurantID int) (bool, error) {
	return false, nil
}
