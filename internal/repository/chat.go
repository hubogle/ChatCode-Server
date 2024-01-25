package repository

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type ChatRepo interface {
	GetUserRoomByUserIDRoomID(ctx context.Context, userID, roomID uint64) (err error)
	GetMessageListByRoomID(ctx context.Context, roomUID uint64, offset, limit int) (messageList []*model.MessageBasic, err error)
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

// GetUserRoomByUserIDRoomID get user room by user id and room id
func (r *chatRepo) GetUserRoomByUserIDRoomID(ctx context.Context, userID, roomID uint64) (err error) {
	_, err = r.UserRoom.WithContext(ctx).Where(
		r.UserRoom.UserID.Eq(userID),
		r.UserRoom.RoomID.Eq(roomID),
	).Take()
	return
}

// GetMessageListByRoomID get message list by room id
func (r *chatRepo) GetMessageListByRoomID(ctx context.Context, roomID uint64, offset, limit int) (messageList []*model.MessageBasic, err error) {
	messageList, _, err = r.MessageBasic.WithContext(ctx).
		Select(r.MessageBasic.Content, r.MessageBasic.SenderID).
		Where(r.MessageBasic.ReceiverID.Eq(roomID)).
		Order(r.MessageBasic.CreatedAt.Desc()).
		FindByPage(limit, (offset-1)*limit)
	return
}
