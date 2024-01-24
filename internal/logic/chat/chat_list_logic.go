package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

// ChatList chat list
// GET /api/v1/chat/list
func (l *logic) ChatList(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatListReq) (err error) {
	err = l.repo.GetUserRoomByUserUIDRoomUID(ctx, uc.UID, req.RoomUid)
	if err != nil {
		return err
	}
	data, err := l.repo.GetMessageListByRoomUID(ctx, req.RoomUid, req.PageSize, req.Page)
	if err != nil {
		return err
	}
	_ = data

	return
}
