package core

import (
	"context"
)

type (
	User struct {
		ID		int		`json:"id"`
		Name	string	`json:"name"`
	}

	UserService interface {
		ListUsers(ctx context.Context) ([]*User, error)
	}

	UserStore interface {
		ListUsers(ctx context.Context) ([]*User, error)
	}
)