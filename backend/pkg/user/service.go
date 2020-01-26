package user

import (
	"context"
	"fmt"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type service struct {
	userStore core.UserStore
}

func NewService(userStore core.UserStore) core.UserService {
	return &service{userStore: userStore}
}

func (s *service) Login(ctx context.Context, id int) error {
	u, err := s.userStore.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("user %v does not exist: %w", id, err)
	}

	if u == nil {
		return fmt.Errorf("user %v does not exist", id)
	}

	return nil
}

func (s *service) Get(ctx context.Context, id int) (*core.User, error) {
	return s.userStore.Get(ctx, id)
}


func (s *service) List(ctx context.Context) ([]*core.User, error) {
	return s.userStore.List(ctx)
}

func (s *service) Create(ctx context.Context, userName string) (*core.User, error) {
	return s.userStore.Create(ctx, userName)
}

