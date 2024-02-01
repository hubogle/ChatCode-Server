package friend

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/friend"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IFriendLogic interface {
	AddFriend(ctx context.Context, uc *jwt.UserClaims, req *friend.AddFriendReq) (err error)
}

func NewFriendLogic(logicSvc *svc.ServiceContext, repo repository.FriendRepo) IFriendLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.FriendRepo
}
