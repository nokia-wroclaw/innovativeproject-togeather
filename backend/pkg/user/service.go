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

func (s *service) List(ctx context.Context) ([]*core.User, error) {
	return s.userStore.List(ctx)
}

func (s *service) Create(ctx context.Context, userName string, lobbyID int, isOwner bool) (*core.User, error) {
	return s.userStore.Create(ctx, userName, lobbyID, isOwner)
}

func (s *service) Get(ctx context.Context, userID int) (*core.User, error) {
	return s.userStore.Get(ctx, userID)
}
