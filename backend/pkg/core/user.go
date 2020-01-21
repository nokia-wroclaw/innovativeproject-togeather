package core

import (
	"context"
)

type (
	User struct {
		ID		int		`json:"user_id"`
		Name	string	`json:"user_name"`
	}

	UserService interface {
		List(ctx context.Context) ([]*User, error)
		Create(ctx context.Context, userName string, lobbyID int, isOwner bool) (*User, error)
	}

	UserStore interface {
		List(ctx context.Context) ([]*User, error)
		Create(ctx context.Context, userName string, lobbyID int, isOwner bool) (*User, error)
	}
)