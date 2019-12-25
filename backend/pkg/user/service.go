package user

import (
	"context"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	userStore core.UserStore
}


func NewService(userStore core.UserStore) core.UserService {
	return &service{userStore: userStore}
}

func (s *service) ListUsers(ctx context.Context) ([]*core.User, error) {
	return s.userStore.ListUsers(ctx)
}
