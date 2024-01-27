package repository

import (
	"context"
	"time"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type RoomRepo interface {
	InsertOneUserRoom(ctx context.Context, userID, roomID uint64) (err error)
	InsertOneRoomBasic(ctx context.Context, roomBasic *model.RoomBasic) (err error)
	GetRoomUserIDByRoomID(ctx context.Context, roomID uint64) (userIDList []uint64, err error)
	GetRoomUserIDByUserIDRoomID(ctx context.Context, userID, roomID uint64) (userRoom *model.UserRoom, err error)
}

func NewRoomRepo(svcRepo *svc.ServiceContext) RoomRepo {
	return &roomRepo{
		ServiceContext: svcRepo,
		Query:          query.Use(svcRepo.Db),
	}
}

type roomRepo struct {
	*svc.ServiceContext
	*query.Query
}

// InsertOneRoomBasic insert one room basic
func (r *roomRepo) InsertOneRoomBasic(ctx context.Context, roomBasic *model.RoomBasic) (err error) {
	return r.RoomBasic.WithContext(ctx).Create(roomBasic)
}

// InsertOneUserRoom insert one user room
func (r *roomRepo) InsertOneUserRoom(ctx context.Context, userID, roomID uint64) (err error) {
	now := time.Now().Unix()
	userRoom := &model.UserRoom{
		UserID:    userID,
		RoomID:    roomID,
		RoomType:  1,
		JoinedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return r.UserRoom.WithContext(ctx).Create(userRoom)
}

func (r *roomRepo) GetRoomUserIDByRoomID(ctx context.Context, roomID uint64) (userIDList []uint64, err error) {
	err = r.UserRoom.WithContext(ctx).
		Where(r.UserRoom.RoomID.Eq(roomID)).
		Pluck(r.UserRoom.UserID, &userIDList)
	return
}

func (r *roomRepo) GetRoomUserIDByUserIDRoomID(ctx context.Context, userID, roomID uint64) (userRoom *model.UserRoom, err error) {
	return r.UserRoom.WithContext(ctx).
		Where(r.UserRoom.UserID.Eq(userID), r.UserRoom.RoomID.Eq(roomID)).
		First()
}
