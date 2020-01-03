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
	rows, err := s.db.QueryxContext(ctx, `SELECT l.id, l.restaurant, r.name, l.owner,
 								l.expires, l.geolat, l.geolon FROM lobbies l
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

		err := rows.Scan(&l.ID, &r.ID, &r.Name, &l.Owner,
						&l.Expires, &loc.GeoLat, &loc.GeoLon)
		if err != nil{
			return nil, err
		}

		adrs, err := geocoder.ReverseGeocode(loc.GeoLat, loc.GeoLon)
		if err != nil {
			return nil, err
		}

		loc.Address = adrs.Street + " " + adrs.City
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
    	lobbies(restaurant, owner, expires, geolat, geolon) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		restaurantID, ownerID, expires, geolat, geolon).Scan(&id)
	if err != nil {
		return nil, err
	}

	var restaurantName string
	row := s.db.QueryRowContext(ctx, `SELECT name FROM restaurants WHERE id = $1`, restaurantID)
	err = row.Scan(&restaurantName)
	if err != nil{
		return nil, err
	}

	return &core.Lobby{
		ID:           id,
		Restaurant: &core.Restaurant{
			ID:   restaurantID,
			Name: restaurantName,
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
