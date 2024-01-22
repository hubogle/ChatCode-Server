package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type HomeRepo interface{}

func NewHomeRepo(svcRepo *svc.ServiceContext) HomeRepo {
	return &homeRepo{
		ServiceContext: svcRepo,
	}
}

type homeRepo struct {
	*svc.ServiceContext
}
