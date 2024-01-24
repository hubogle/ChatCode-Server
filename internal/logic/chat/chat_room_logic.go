package chat

import (
	"context"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/chat"
)

type RoomMessageStruct struct {
	Message string `json:"message"`
	RoomUID string `json:"room_uid"`
}

// ChatRoom room chat
// POST /api/v1/chat/room
func (l *logic) ChatRoom(ctx context.Context, ws *websocket.Conn, uc *jwt.UserClaims, req *chat.ChatRoomReq) (err error) {
	l.OnLine[uc.UID] = ws
	for {
		ms := new(RoomMessageStruct)
		err = ws.ReadJSON(ms)
		if err != nil {
			return
		}

		err = l.repo.GetUserRoomByUserUIDRoomUID(ctx, uc.UID, ms.RoomUID)
		if err != nil {
			return err
		}
		nowTime := time.Now().Unix()
		messageBasic := &model.MessageBasic{
			UserUID:   uc.UID,
			RoomUID:   ms.RoomUID,
			Content:   ms.Message,
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		}

		err = l.repo.InsertOneMessageBasic(ctx, messageBasic)
		if err != nil {
			return err
		}

		userUIDs, err := l.repo.GetUserRoomByRoomUID(ctx, ms.RoomUID)
		if err != nil {
			return err
		}
		for _, userUID := range userUIDs {
			if conn, ok := l.OnLine[userUID]; ok {
				err = conn.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					return err
				}
			}
		}
	}
}
