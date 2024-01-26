package room

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/room"
	"go.uber.org/zap"
)

// RoomJoin join room
// POST /api/v1/room/join
func (l *logic) RoomJoin(ctx context.Context, uc *jwt.UserClaims, req *room.RoomJoinReq) (err error) {
	err = l.repo.InsertOneUserRoom(ctx, uc.UID, req.RoomId)

	if err != nil {
		l.Log.Error("InsertOneUserRoom", zap.Error(err))
		return code.WithCodeMsg(code.RoomUserJoinErr)
	}

	return
}
