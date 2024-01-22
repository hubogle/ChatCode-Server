package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type PingRepo interface{}

func NewPingRepo(svcRepo *svc.ServiceContext) PingRepo {
	return &pingRepo{
		ServiceContext: svcRepo,
	}
}

type pingRepo struct {
	*svc.ServiceContext
}
