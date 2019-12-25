package lobby

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type lobbyStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) core.LobbyStore {
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

	var lobbies []*core.Lobby
	for rows.Next(){
		l := core.Lobby{}
		err := rows.Scan(&l.ID, &l.Restaurant.ID, &l.Restaurant.Name, &l.Owner,
						&l.Expires, &l.GeoLat, &l.GeoLon)
		if err != nil{
			return nil, err
		}

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
	geolat float64,
	geolon float64,
) (*core.Lobby, error) {
	var id int

	err := s.db.QueryRowContext(ctx, `INSERT INTO
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
	}, nil
}
