package repository

import (
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type RoomRepo interface{}

func NewRoomRepo(svcRepo *svc.ServiceContext) RoomRepo {
	return &roomRepo{
		ServiceContext: svcRepo,
	}
}

type roomRepo struct {
	*svc.ServiceContext
}
