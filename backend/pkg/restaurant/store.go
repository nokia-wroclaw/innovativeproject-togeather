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

func (s *restaurantStore) ListRestaurants(ctx context.Context) ([]core.Restaurant, error) {
	query := "SELECT id, name, address FROM rests"

	rows, err := s.db.Query(query)

	if err != nil{
		//some error
		//TODO
	}
	defer rows.Close()

	rests := make([]core.Restaurant, 0)
	for rows.Next(){
		rest := core.Restaurant{}
		err := rows.Scan(&rest.ID, &rest.Name, &rest.Address)
		if err != nil{
			//some error
			//TODO
		}
		rests = append(rests, rest)
	}
	if err = rows.Err(); err != nil{
		//some error
		//TODO
	}

	return rests, nil

}