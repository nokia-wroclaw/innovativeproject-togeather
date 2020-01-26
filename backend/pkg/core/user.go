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
		Get(ctx context.Context, id int) (*User, error)
		List(ctx context.Context) ([]*User, error)
		Create(ctx context.Context, userName string, lobbyID int, isOwner bool) (*User, error)
	}

	UserStore interface {
		Get(ctx context.Context, id int) (*User, error)
		List(ctx context.Context) ([]*User, error)
		Create(ctx context.Context, userName string, lobbyID int, isOwner bool) (*User, error)
	}
)
