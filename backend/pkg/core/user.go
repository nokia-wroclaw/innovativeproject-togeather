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
		ListUsers(ctx context.Context) ([]*User, error)
	}

	UserStore interface {
		ListUsers(ctx context.Context) ([]*User, error)
	}
)