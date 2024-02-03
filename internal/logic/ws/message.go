package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/pkg/errors"
)

type MsgType int32

const (
	Msg_Type_None    MsgType = iota // message type none
	Msg_Type_login                  // 注册链接，客户端向服务端发送请求，建立连接
	Msg_Type_Heart                  // 心跳保证连接存活
	Msg_Type_Message                // 消息发送
	Msg_Type_Ack                    // message type ack
	Msg_Type_Sync                   // 消息离线同步
)

// 会话类型
type SessionType uint8

const (
	SessionType_UnKnow SessionType = 0 // 未知
	SessionType_Single SessionType = 1 // 单聊
	SessionType_Room   SessionType = 2 // 群聊
)

// 用户所发送内容的消息类型
type MessageType uint8

const (
	MessageType_UnKnow  MessageType = 0 // 未知
	MessageType_Text    MessageType = 1 // 文本类型消息
	MessageType_Picture MessageType = 2 // 图片类型消息
)

// Message 客户端发送给服务端的消息
type Message struct {
	SessionType SessionType `json:"session_type"` // 会话类型
	ReceiverID  uint64      `json:"receiver_id"`  // 接收者id, 群组id 或 用户 id
	SenderID    uint64      `json:"sender_id"`    // 发送者id
	MessageType MessageType `json:"message_type"` // 消息类型
	Content     string      `json:"content"`      // 消息内容
	SendAt      int64       `json:"send_at"`      // 消息发送时间
}

type LoginMessage struct {
	Token string `json:"token"`
}

// ClientMessage 客户端发送给服务端的消息，data 内容才是 Message
type ClientMessage struct {
	Type MsgType         `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ServerMessage struct {
	Type MsgType `json:"type"`
	Code int     `json:"code"`
	Msg  string  `json:"msg,omitempty"`
	Data Message `json:"data,omitempty"`
}

// SendToUser 对指定 user id 的用户发送消息
func SendToUser(msg *Message, userID uint64) error {
	_, err := query.Q.UserBasic.Where(query.UserBasic.UID.Eq(userID)).Take()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("user id %d not found", userID))
	}

	if userID != msg.SenderID {
		now := time.Now().Unix()
		messageBasic := &model.MessageBasic{
			SenderID:    msg.SenderID,
			ReceiverID:  msg.ReceiverID,
			SessionType: int32(msg.SessionType),
			Content:     string(msg.Content),
			ContentType: int32(msg.MessageType),
			SendAt:      msg.SendAt,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		err = query.Q.MessageBasic.Create(messageBasic)
		if err != nil {
			return errors.Wrap(err, "message basic create error")
		}
	}

	conn := ConnManager.GetConn(userID)
	if conn == nil {
		return errors.New(fmt.Sprintf("user id %d not online", userID))
	}
	bytes := MockServerMessage(Msg_Type_Message, code.Success, msg)
	conn.SendMsg(userID, string(bytes))

	return nil
}

// SendToRoom 对指定 room id 的群组发送消息
func SendToRoom(msg *Message, roomID uint64) error {
	ctx := context.Background()
	repo := repository.NewRoomRepo(svc.GetServiceContext())
	_, err := repo.GetRoomUserIDByUserIDRoomID(ctx, msg.SenderID, roomID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("user id %d not in room id %d", msg.SenderID, roomID))
	}

	userIDList, err := repo.GetRoomUserIDByRoomID(ctx, roomID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("get room id %d user id list error", roomID))
	}

	now := time.Now().Unix()
	messageBasic := &model.MessageBasic{
		SenderID:    msg.SenderID,
		ReceiverID:  msg.ReceiverID,
		SessionType: int32(msg.SessionType),
		Content:     string(msg.Content),
		ContentType: int32(msg.MessageType),
		SendAt:      msg.SendAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err = query.Q.MessageBasic.Create(messageBasic)
	if err != nil {
		return errors.Wrap(err, "message basic create error")
	}

	bytes := MockServerMessage(Msg_Type_Message, code.Success, msg)
	for _, userID := range userIDList {
		conn := ConnManager.GetConn(userID)
		if conn == nil {
			continue
		}
		conn.SendMsg(userID, string(bytes))
	}

	return nil
}

func MockServerMessage(msgType MsgType, statusCode int, data *Message) []byte {
	msg := &ServerMessage{
		Type: msgType,
		Code: statusCode,
		Data: *data,
	}
	if statusCode != code.Success {
		msg.Msg = code.GetCoder(statusCode).String()
	}
	result, _ := json.Marshal(msg)
	return result
}
