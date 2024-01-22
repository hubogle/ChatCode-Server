package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IChatLogic interface {
	ChatPrivate(ctx context.Context, req *chat.ChatPrivateReq) (err error)
	ChatRoom(ctx context.Context, req *chat.ChatRoomReq) (err error)
}

func NewChatLogic(logicSvc *svc.ServiceContext, repo repository.ChatRepo) IChatLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.ChatRepo
}
