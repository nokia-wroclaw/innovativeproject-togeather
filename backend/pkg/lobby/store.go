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
	rows, err := s.db.QueryxContext(ctx, `SELECT l.id, l.restaurant, r.name, 
		r.delivery, l.owner,	l.expires, l.geolat, l.geolon, l.address FROM lobbies l
		JOIN restaurants r ON r.id = l.restaurant`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	lobbies := make([]*core.Lobby, 0)
	for rows.Next(){
		l := core.Lobby{}
		r := core.Restaurant{}
		loc := core.Location{}

		err := rows.Scan(&l.ID, &r.ID, &r.Name, &r.Delivery, &l.Owner,
						&l.Expires, &loc.GeoLat, &loc.GeoLon, &loc.Address)
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
	ownerID int,
	expires *time.Time,
	address string,
) (*core.Lobby, error) {
	var id int

	geolat, geolon, err := geocoder.Geocode(address)
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, `INSERT INTO
    	lobbies(restaurant, owner, expires, geolat, geolon, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		restaurantID, ownerID, expires, geolat, geolon, address).Scan(&id)
	if err != nil {
		return nil, err
	}

	var restaurantName string
	var restaurantDelivery float32
	row := s.db.QueryRowContext(ctx, `SELECT name, delivery FROM restaurants WHERE id = $1`, restaurantID)
	err = row.Scan(&restaurantName, &restaurantDelivery)
	if err != nil{
		return nil, err
	}

	return &core.Lobby{
		ID:           id,
		Restaurant: &core.Restaurant{
			ID:   restaurantID,
			Name: restaurantName,
			Delivery: restaurantDelivery,
		},
		Owner:        ownerID,
		Expires:      *expires,
		Location: &core.Location{
			GeoLat:  geolat,
			GeoLon:  geolon,
			Address: address,
		},
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

	_, err = s.db.ExecContext(ctx, `UPDATE lobbies SET restaurant = $1, expires = $2,
		geolat = $3, geolon = $4, address = $5 WHERE id = $6`, restaurantID, expires, geolat, geolon, address, lobbyID)
	if err != nil {
		return nil, err
	}

	var restaurantName string
	var restaurantDelivery float32
	row := s.db.QueryRowContext(ctx, `SELECT name, delivery FROM restaurants WHERE id = $1`, restaurantID)
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
		Owner:        ownerID,
		Expires:      *expires,
		Location: &core.Location{
			GeoLat:  geolat,
			GeoLon:  geolon,
			Address: address,
		},
	}, nil
}
