package repository

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type FriendRepo interface {
	InsertOneUserFriend(ctx context.Context, userFriend *model.UserFriend) (err error)
}

func NewFriendRepo(svcRepo *svc.ServiceContext) FriendRepo {
	return &friendRepo{
		ServiceContext: svcRepo,
		Query:          query.Use(svcRepo.Db),
	}
}

type friendRepo struct {
	*svc.ServiceContext
	*query.Query
}

// InsertOneUserFriend insert one user friend
func (r *friendRepo) InsertOneUserFriend(ctx context.Context, userFriend *model.UserFriend) (err error) {
	return r.UserFriend.WithContext(ctx).Create(userFriend)
}
