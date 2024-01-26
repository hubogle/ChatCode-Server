package room

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/room"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IRoomLogic interface {
	RoomJoin(ctx context.Context, uc *jwt.UserClaims, req *room.RoomJoinReq) (err error)
	RoomGet(ctx context.Context, req *room.RoomGetReq) (err error)
	RoomCreate(ctx context.Context, uc *jwt.UserClaims, req *room.RoomCreateReq) (resp room.RoomCreateResp, err error)
}

func NewRoomLogic(logicSvc *svc.ServiceContext, repo repository.RoomRepo) IRoomLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.RoomRepo
}
