package room

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/room"
	"go.uber.org/zap"
)

// RoomCreate create room
// POST /api/v1/room/create
func (l *logic) RoomCreate(ctx context.Context, uc *jwt.UserClaims, req *room.RoomCreateReq) (resp room.RoomCreateResp, err error) {
	uid := uuid.New().ID()
	now := time.Now().Unix()

	if req.Salt != nil && *req.Salt == "" {
		req.Salt = nil
	}

	roomBasic := model.RoomBasic{
		UID:       uint64(uid),
		UserID:    uc.UID,
		Name:      req.Name,
		Info:      req.Info,
		Salt:      req.Salt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = l.repo.InsertOneRoomBasic(ctx, &roomBasic)

	if err != nil {
		l.Log.Error("InsertOneRoomBasic", zap.Error(err))
		err = code.WithCodeMsg(code.RoomCreateErr)
		return resp, err
	}
	resp.RoomId = uint64(uid)

	userRoom := &model.UserRoom{
		UserID:    uc.UID,
		RoomID:    uint64(uid),
		RoomType:  1,
		JoinedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = l.repo.InsertOneUserRoom(ctx, userRoom)

	if err != nil {
		return resp, err
	}

	return
}
