package chat

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

// ChatMessage chat message list
// GET /api/v1/chat/message/list
func (l *logic) ChatMessage(ctx context.Context, uc *jwt.UserClaims, req *chat.ChatMessageReq) (resp chat.ChatMessageResp, err error) {
	var messageList []*model.MessageBasic
	if req.Type == 1 {
		messageList, err = l.repo.GetMessageListByUserID(ctx, req.Uid, uc.UID, 1, 200)
		if err != nil {
			return chat.ChatMessageResp{}, err
		}
	}
	if req.Type == 2 {
		messageList, err = l.repo.GetMessageListByRoomID(ctx, req.Uid, 1, 200)
	}
	resp.List = make([]types.MessageItemInfo, 0, len(messageList))

	for _, message := range messageList {
		resp.List = append(resp.List, types.MessageItemInfo{
			Type:    req.Type,
			Content: message.Content,
			Uid:     message.SenderID,
			SendAt:  message.SendAt,
		})
	}
	return
}
