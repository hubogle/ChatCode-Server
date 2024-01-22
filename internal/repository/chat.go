package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type ChatRepo interface{}

func NewChatRepo(svcRepo *svc.ServiceContext) ChatRepo {
	return &chatRepo{
		ServiceContext: svcRepo,
	}
}

type chatRepo struct {
	*svc.ServiceContext
}
