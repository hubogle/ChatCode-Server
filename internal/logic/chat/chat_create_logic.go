package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/chat"
	"go.uber.org/zap"
)

// ChatCreate create chat
// POST /api/v1/chat/create
func (l *logic) ChatCreate(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatCreateReq) (err error) {
	if req.Type == 1 {
		if err = l.repo.InsertOneUserFriend(ctx, uc.UID, req.Uid); err != nil {
			l.Log.Error("InsertOneUserFriend", zap.Error(err))
			return code.WithCodeMsg(code.ChatCreateErr)
		}
	}

	if req.Type == 2 {
		if err = l.repo.InsertOneUserRoom(ctx, uc.UID, req.Uid); err != nil {
			l.Log.Error("InsertOneUserRoom", zap.Error(err))
			return code.WithCodeMsg(code.ChatCreateErr)
		}
	}

	return
}
