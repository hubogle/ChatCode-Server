package chat

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IChatLogic interface {
	ChatPrivate(ctx context.Context, req *chat.ChatPrivateReq) (err error)
	ChatRoom(ctx context.Context, ws *websocket.Conn, uc *jwt.UserClaims, req *chat.ChatRoomReq) (err error)
	ChatList(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatListReq) (err error)
}

func NewChatLogic(logicSvc *svc.ServiceContext, repo repository.ChatRepo) IChatLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
		OnLine:         make(map[uint64]*websocket.Conn), // TODO 这里应该用并发 map
	}
}

type logic struct {
	*svc.ServiceContext
	repo   repository.ChatRepo
	OnLine map[uint64]*websocket.Conn
}
