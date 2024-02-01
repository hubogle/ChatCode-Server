package repository

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type UserRepo interface {
	GetUserBasicByAccountPassword(ctx context.Context, account, password string) (userBasic *model.UserBasic, err error)
	InsertOneUserBasic(ctx context.Context, userBasic *model.UserBasic) (err error)
}

func NewUserRepo(svcRepo *svc.ServiceContext) UserRepo {
	return &userRepo{
		ServiceContext: svcRepo,
		Query:          query.Use(svcRepo.Db),
	}
}

type userRepo struct {
	*svc.ServiceContext
	*query.Query
}

// GetUserBasicByAccountPassword get user basic by account and password
func (r *userRepo) GetUserBasicByAccountPassword(ctx context.Context, account, password string) (userBasic *model.UserBasic, err error) {
	return r.UserBasic.WithContext(ctx).Where(
		r.UserBasic.Account.Eq(account),
		r.UserBasic.Password.Eq(password),
	).First()
}

// InsertOneUserBasic insert one user basic
func (r *userRepo) InsertOneUserBasic(ctx context.Context, userBasic *model.UserBasic) (err error) {
	return r.UserBasic.WithContext(ctx).Create(userBasic)
}
