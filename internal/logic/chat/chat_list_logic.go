package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types"
	"github.com/hubogle/chatcode-server/internal/types/chat"
	"go.uber.org/zap"
)

// ChatList chat list
// GET /api/v1/chat/list
func (l *logic) ChatList(ctx context.Context, uc *jwt.UserClaims) (resp chat.ChatListResp, err error) {
	userFriendList, err := l.repo.GetUserFriendByUserID(ctx, uc.UID)
	if err != nil {
		l.Log.Error("failed to get user friend", zap.Error(err))
		return chat.ChatListResp{}, code.WithCodeMsg(code.UserGetListErr)
	}

	roomBasicList, err := l.repo.GetRoomBasicByUserID(ctx, uc.UID)
	if err != nil {
		l.Log.Error("failed to get room basic", zap.Error(err))
		return chat.ChatListResp{}, code.WithCodeMsg(code.RoomGetListErr)
	}

	resp.List = make([]types.ChatItemInfo, 0, len(userFriendList)+len(roomBasicList))
	for _, userFriend := range userFriendList {
		resp.List = append(resp.List, types.ChatItemInfo{
			Name: userFriend.Account,
			Uid:  userFriend.UID,
			Type: 1,
		})
	}

	for _, roomBasic := range roomBasicList {
		resp.List = append(resp.List, types.ChatItemInfo{
			Name: roomBasic.Name,
			Uid:  roomBasic.UID,
			Type: 2,
		})
	}
	return
}
