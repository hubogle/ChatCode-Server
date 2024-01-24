package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type WsRepo interface {
}

func NewWsRepo(svcRepo *svc.ServiceContext) WsRepo {

	return &wsRepo{
		ServiceContext: svcRepo,
	}
}

type wsRepo struct {
	*svc.ServiceContext
}
