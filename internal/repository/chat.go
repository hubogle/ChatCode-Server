package repository

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type ChatRepo interface {
	GetUserRoomByUserUIDRoomUID(ctx context.Context, userUID uint64, roomUID string) (err error)
	GetUserRoomByRoomUID(ctx context.Context, roomUID string) (userUIDs []uint64, err error)
	InsertOneMessageBasic(ctx context.Context, messageBasic *model.MessageBasic) (err error)
	GetMessageListByRoomUID(ctx context.Context, roomUID string, offset, limit int) (messageList []*model.MessageBasic, err error)
}

func NewChatRepo(svcRepo *svc.ServiceContext) ChatRepo {
	return &chatRepo{
		ServiceContext: svcRepo,
		Query:          query.Use(svcRepo.Db),
	}
}

type chatRepo struct {
	*svc.ServiceContext
	*query.Query
}

// GetUserRoomByUserUIDRoomUID get user room by user uid and room uid
func (r *chatRepo) GetUserRoomByUserUIDRoomUID(ctx context.Context, userUID uint64, roomUID string) (err error) {
	_, err = r.UserRoom.WithContext(ctx).Where(
		r.UserRoom.UserUID.Eq(userUID),
		r.UserRoom.RoomUID.Eq(roomUID),
	).Take()
	return
}

// GetUserRoomByRoomUID get user room by room uid
func (r *chatRepo) GetUserRoomByRoomUID(ctx context.Context, roomUID string) (userUIDs []uint64, err error) {
	err = r.UserRoom.WithContext(ctx).
		Where(r.UserRoom.RoomUID.Eq(roomUID)).
		Pluck(r.UserRoom.UserUID, &userUIDs)
	return
}

func (r *chatRepo) InsertOneMessageBasic(ctx context.Context, messageBasic *model.MessageBasic) (err error) {
	return r.MessageBasic.Create(messageBasic)
}

// GetMessageListByRoomUID get message list by room uid
func (r *chatRepo) GetMessageListByRoomUID(ctx context.Context, roomUID string, offset, limit int) (messageList []*model.MessageBasic, err error) {
	messageList, _, err = r.MessageBasic.WithContext(ctx).
		Select(r.MessageBasic.Content, r.MessageBasic.UserUID).
		Where(r.MessageBasic.RoomUID.Eq(roomUID)).
		Order(r.MessageBasic.CreatedAt.Desc()).
		FindByPage(limit, (offset-1)*limit)
	return
}
