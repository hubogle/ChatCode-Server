package repository

import (
	"context"
	"time"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../mocks/repository/$GOFILE -package mock_repo

type ChatRepo interface {
	GetUserRoomByUserIDRoomID(ctx context.Context, userID, roomID uint64) (err error)
	GetMessageListByRoomID(ctx context.Context, roomUID uint64, offset, limit int) (messageList []*model.MessageBasic, err error)
	GetMessageListByUserID(ctx context.Context, userID, senderID uint64, offset, limit int) (messageList []*model.MessageBasic, err error)
	InsertOneMessageBasic(ctx context.Context, messageBasic *model.MessageBasic) (err error)
	GetUserFriendByUserID(ctx context.Context, userID uint64) (userBasicList []*model.UserBasic, err error)
	GetRoomBasicByUserID(ctx context.Context, userID uint64) (roomBasicList []*model.RoomBasic, err error)
	InsertOneUserFriend(ctx context.Context, userID, friend uint64) (err error)
	InsertOneUserRoom(ctx context.Context, userID, roomID uint64) (err error)
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
		FindByPage((offset-1)*limit, limit)
	return
}

// GetMessageListByUserID get message list by user id
func (r *chatRepo) GetMessageListByUserID(ctx context.Context, userID, senderID uint64, offset, limit int) (messageList []*model.MessageBasic, err error) {
	subQuery := r.MessageBasic.WithContext(ctx).Where(r.MessageBasic.ReceiverID.Eq(senderID), r.MessageBasic.SenderID.Eq(userID))
	messageList, _, err = r.MessageBasic.Debug().WithContext(ctx).
		Select(r.MessageBasic.Content, r.MessageBasic.SenderID, r.MessageBasic.SendAt).
		Where(r.MessageBasic.ReceiverID.Eq(userID), r.MessageBasic.SenderID.Eq(senderID)).
		Or(subQuery).
		Order(r.MessageBasic.CreatedAt).
		FindByPage((offset-1)*limit, limit)
	return
}

func (r *chatRepo) InsertOneMessageBasic(ctx context.Context, messageBasic *model.MessageBasic) (err error) {
	return r.MessageBasic.WithContext(ctx).Create(messageBasic)
}

// GetUserFriendByUserID get user friend by user id
func (r *chatRepo) GetUserFriendByUserID(ctx context.Context, userID uint64) (userBasicList []*model.UserBasic, err error) {
	userBasicList, err = r.UserBasic.WithContext(ctx).
		Select(r.UserBasic.UID, r.UserBasic.Account).
		LeftJoin(r.UserFriend, r.UserFriend.FriendID.EqCol(r.UserBasic.UID)).
		Where(r.UserFriend.UserID.Eq(userID)).
		Find()
	return
}

// GetRoomBasicByUserID get room id by user id
func (r *chatRepo) GetRoomBasicByUserID(ctx context.Context, userID uint64) (roomBasicList []*model.RoomBasic, err error) {
	roomBasicList, err = r.RoomBasic.WithContext(ctx).
		Select(r.RoomBasic.UID, r.RoomBasic.Name).
		LeftJoin(r.UserRoom, r.UserRoom.RoomID.EqCol(r.RoomBasic.UID)).
		Where(r.UserRoom.UserID.Eq(userID)).
		Find()

	return
}

// InsertOneUserFriend insert one user friend
func (r *chatRepo) InsertOneUserFriend(ctx context.Context, userID, friend uint64) (err error) {
	now := time.Now().Unix()
	userFriend := &model.UserFriend{
		UserID:    userID,
		FriendID:  friend,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return r.UserFriend.WithContext(ctx).Create(userFriend)
}

// InsertOneUserRoom insert one user room
func (r *chatRepo) InsertOneUserRoom(ctx context.Context, userID, roomID uint64) (err error) {
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
