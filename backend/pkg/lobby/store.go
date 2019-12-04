package lobby

import (
	"context"

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
	rows, err := s.db.QueryxContext(ctx, `SELECT 
       id, restaurant, owner, expires FROM lobbies`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var lobbies []*core.Lobby
	for rows.Next(){
		var l core.Lobby
		err := rows.StructScan(&l)
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
