package lobby

import (
	"context"
	"time"

	"github.com/jasonwinn/geocoder"
	"github.com/jmoiron/sqlx"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type lobbyStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) core.LobbyStore {
	geocoder.SetAPIKey("PdBSQAE97uUFd6NKJYsBO35voZcXX0qD")
	return &lobbyStore{db: db}
}

func (s *lobbyStore) List(ctx context.Context) ([]*core.Lobby, error) {
	currentTime := time.Now().Format(time.RFC3339)

	rows, err := s.db.QueryxContext(ctx, `SELECT l.id, l.restaurant, r.name, 
		r.delivery, l.expires, l.geolat, l.geolon, l.address 
		FROM lobbies l 
		JOIN restaurants r ON r.id = l.restaurant 
		WHERE l.expires > $1`, currentTime)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	lobbies := make([]*core.Lobby, 0)
	for rows.Next(){
		l := core.Lobby{}
		r := core.Restaurant{}
		loc := core.Location{}

		err := rows.Scan(&l.ID, &r.ID, &r.Name, &r.Delivery, &l.Expires,
			&loc.GeoLat, &loc.GeoLon, &loc.Address)
		if err != nil{
			return nil, err
		}

		l.Location = &loc
		l.Restaurant = &r
		lobbies = append(lobbies, &l)
	}
	if err = rows.Err(); err != nil{
		return nil, err
	}

	return lobbies, nil
}

func (s *lobbyStore) Create(
	ctx context.Context,
	restaurantID int,
	userID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	geolat, geolon, err := geocoder.Geocode(address)
	if err != nil {
		return nil, err
	}

	var lobbyID int
	err = s.db.QueryRowContext(ctx, `INSERT INTO
    	lobbies(restaurant, expires, geolat, geolon, address) 
    	VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		restaurantID, expires, geolat, geolon, address).Scan(&lobbyID)
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, `INSERT INTO 
    	lobbys_users (lobby_id, client_id, is_owner) VALUES ($1, $2, $3)`,
		lobbyID, userID, true,
	)

	return &core.Lobby{
		ID:		lobbyID,
	}, nil
}

func (s *lobbyStore) Edit(
	ctx context.Context,
	lobbyID int,
	restaurantID int,
	ownerID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	geolat, geolon, err := geocoder.Geocode(address)
	if err != nil {
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, `UPDATE lobbies SET restaurant = $1, 
		expires = $2, geolat = $3, geolon = $4, address = $5 WHERE id = $6`,
		restaurantID, expires, geolat, geolon, address, lobbyID)
	if err != nil {
		return nil, err
	}

	var restaurantName string
	var restaurantDelivery float32
	row := s.db.QueryRowContext(ctx, `SELECT name, delivery 
		FROM restaurants WHERE id = $1`, restaurantID)
	err = row.Scan(&restaurantName, &restaurantDelivery)
	if err != nil{
		return nil, err
	}

	return &core.Lobby{
		ID:           lobbyID,
		Restaurant: &core.Restaurant{
			ID:   restaurantID,
			Name: restaurantName,
			Delivery: restaurantDelivery,
		},
		Expires:      expires,
		Location: &core.Location{
			GeoLat:  geolat,
			GeoLon:  geolon,
			Address: address,
		},
	}, nil
}

func (s *lobbyStore) Join(ctx context.Context, lobbyID int, userID int) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO 
    	lobbys_users (lobby_id, client_id, is_owner) VALUES ($1, $2, $3)`,
    	lobbyID, userID, false,
	)

	return err
}

func (s *lobbyStore) Get(ctx context.Context, lobbyID int) (*core.Lobby, error) {
	row := s.db.QueryRowContext(ctx, `SELECT l.restaurant, r.name, 
		r.delivery, l.expires, l.geolat, l.geolon, l.address 
		FROM lobbies l 
		JOIN restaurants r ON r.id = l.restaurant 
		WHERE l.id = $1`, lobbyID)

	lobby := core.Lobby{ID: lobbyID}
	r := core.Restaurant{}
	l := core.Location{}

	err := row.Scan(&r.ID, &r.Name, &r.Delivery, &lobby.Expires,
		&l.GeoLat, &l.GeoLon, &l.Address)
	if err != nil{
		return nil, err
	}

	lobby.Location = &l
	lobby.Restaurant = &r

	return &lobby, nil
}

func (s *lobbyStore) Clean(ctx context.Context) {
	limitTime := time.Now().Add(time.Minute*(-30)).Format(time.RFC3339)
	s.db.ExecContext(ctx, `DELETE FROM lobbies WHERE expires < $1`, limitTime)
}

func (s *lobbyStore) BelongsToLobby(ctx context.Context, userID int, lobbyID int) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM clients 
		WHERE id = $1 AND lobby = $2)`, userID, lobbyID).Scan(&exists)
	if err != nil{
		return false, err
	}

	return exists, nil
}

func (s *lobbyStore) AddToCart(ctx context.Context, userID int, lobbyID int, mealID int) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO orders(client_id, lobby_id, meal_id)
		VALUES ($1, $2, $3)`, userID, lobbyID, mealID)
	if err != nil{
		return err
	}

	return nil
}

func (s *lobbyStore) DelFromCart(ctx context.Context, userID int, lobbyID int, mealID int) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM orders WHERE id IN 
		(SELECT id FROM orders 
		WHERE client_id = $1 AND lobby_id = $2 AND meal_id = $3 LIMIT 1)`,
		userID, lobbyID, mealID)
	if err != nil{
		return err
	}

	return nil
}

func (s *lobbyStore) CollectCartInfo(ctx context.Context, userID int, lobbyID int) (*core.CartState, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT m.id, m.name, m.price 
		FROM meals m JOIN orders o on m.id = o.meal_id
		WHERE o.client_id = $1 AND o.lobby_id = $2`, userID, lobbyID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	cartValue := float32(0.0)
	meals := make([]*core.Meal, 0)
	for rows.Next(){
		m := core.Meal{}

		err := rows.Scan(&m.ID, &m.Name, &m.Price)
		if err != nil{
			return nil, err
		}

		cartValue += m.Price
		meals = append(meals, &m)
	}
	if err = rows.Err(); err != nil{
		return nil, err
	}

	var lobbyCount int
	err = s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM lobbys_users
		WHERE lobby_id = $1`, lobbyID).Scan(&lobbyCount)
	if err != nil{
		return nil, err
	}

	var restDelivery float32
	err = s.db.QueryRowContext(ctx, `SELECT r.delivery FROM restaurants r 
		JOIN lobbies l on r.id = l.restaurant WHERE l.id = $1`, lobbyID).Scan(&restDelivery)
	if err != nil{
		return nil, err
	}

	cartDelivery := restDelivery / float32(lobbyCount)
	cartDelivery = float32(int(cartDelivery * 100)) / 100
	cartValue  = float32(int(cartValue * 100)) / 100

	return &core.CartState{
			Meals: meals,
			CartValue: cartValue,
			DeliveryCost: cartDelivery,
			LobbyCount: lobbyCount}, nil
}
