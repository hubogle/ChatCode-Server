package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IChatLogic interface {
	ChatList(ctx context.Context, uc *jwt.UserClaims) (resp chat.ChatListResp, err error)
	ChatCreate(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatCreateReq) (err error)
	ChatMessage(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatMessageReq) (resp chat.ChatMessageResp, err error)
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
