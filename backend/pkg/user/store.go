package user

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type userStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) core.UserStore {
	return &userStore{db: db}
}

func (s *userStore) List(ctx context.Context) ([]*core.User, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name FROM clients`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	users := make([]*core.User, 0)
	for rows.Next(){
		usr := core.User{}
		err := rows.Scan(&usr.ID, &usr.Name)
		if err != nil{
			return nil, err
		}

		users = append(users, &usr)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return users, nil
}

func (s *userStore) Create(ctx context.Context, userName string, lobbyID int,
	isOwner bool) (*core.User, error) {
	var userID int
	err := s.db.QueryRowContext(ctx, `INSERT INTO clients(name, lobby, is_owner) 
		VALUES ($1, $2, $3) RETURNING id`, userName, lobbyID, isOwner).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &core.User{ID:userID, Name:userName}, nil
}

func (s *userStore) Get(ctx context.Context, userID int) (*core.User, error) {
	var userName string
	err := s.db.QueryRowContext(ctx, `SELECT name FROM clients 
		WHERE id = $1`, userID).Scan(&userName)
	if err != nil {
		return nil, err
	}

	return &core.User{ID:userID, Name:userName}, nil
}
