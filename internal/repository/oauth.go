package repository

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type OauthRepo interface {
	GetUserBasicByAccountPassword(ctx context.Context, account, password string) (userBasic *model.UserBasic, err error)
}

func NewOauthRepo(svcRepo *svc.ServiceContext) OauthRepo {
	return &oauthRepo{
		ServiceContext: svcRepo,
		Query:          query.Use(svcRepo.Db),
	}
}

type oauthRepo struct {
	*svc.ServiceContext
	*query.Query
}

// GetUserBasicByAccountPassword get user basic by account and password
func (r *oauthRepo) GetUserBasicByAccountPassword(ctx context.Context, account, password string) (userBasic *model.UserBasic, err error) {
	return r.UserBasic.WithContext(ctx).Where(
		r.UserBasic.Account.Eq(account),
		r.UserBasic.Password.Eq(password),
	).First()
}
