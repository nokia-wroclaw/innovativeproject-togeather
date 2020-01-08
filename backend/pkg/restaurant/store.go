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

func (s *restaurantStore) ListRestaurants(ctx context.Context) ([]*core.Restaurant, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name, address, delivery FROM restaurants`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	rests := make([]*core.Restaurant, 0)
	for rows.Next(){
		rest := core.Restaurant{}
		err := rows.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.Delivery)
		if err != nil{
			return nil, err
		}

		rests = append(rests, &rest)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return rests, nil
}

func (s *restaurantStore) RestaurantMenu(ctx context.Context, restaurantID int) ([]*core.Meal, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name, price, description, owning_restaurant 
												FROM meals WHERE owning_restaurant = $1;`, restaurantID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	menu := make([]*core.Meal, 0)
	for rows.Next(){
		meal := core.Meal{}
		err := rows.Scan(&meal.ID, &meal.Name, &meal.Price, &meal.Description, &meal.RestaurantID)
		if err != nil{
			return nil, err
		}

		menu = append(menu, &meal)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return menu, nil
}

func (s *restaurantStore) GetRestaurant(ctx context.Context, restaurantID int) (*core.Restaurant, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name, price, description, owning_restaurant 
												FROM meals WHERE owning_restaurant = $1;`, restaurantID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	menu := make([]*core.Meal, 0)
	for rows.Next(){
		meal := core.Meal{}
		err := rows.Scan(&meal.ID, &meal.Name, &meal.Price, &meal.Description, &meal.RestaurantID)
		if err != nil{
			return nil, err
		}

		menu = append(menu, &meal)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	rest := core.Restaurant{}
	rest.Menu = menu

	row := s.db.QueryRowContext(ctx, `SELECT id, name, address, delivery FROM restaurants WHERE id = $1`, restaurantID)
	err = row.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.Delivery)
	if err != nil{
		return nil, err
	}

	return &rest, nil
}