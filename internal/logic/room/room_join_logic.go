package room

import (
	"context"
	"time"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/room"
)

// RoomJoin join room
// POST /api/v1/room/join
func (l *logic) RoomJoin(ctx context.Context, uc *jwt.UserClaims, req *room.RoomJoinReq) (err error) {
	now := time.Now().Unix()
	userRoom := &model.UserRoom{
		UserID:    uc.UID,
		RoomID:    req.RoomId,
		RoomType:  1,
		JoinedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = l.repo.InsertOneUserRoom(ctx, userRoom)
	if err != nil {
		return err
	}

	return
}
