package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type OauthRepo interface{}

func NewOauthRepo(svcRepo *svc.ServiceContext) OauthRepo {
	return &oauthRepo{
		ServiceContext: svcRepo,
	}
}

type oauthRepo struct {
	*svc.ServiceContext
}
