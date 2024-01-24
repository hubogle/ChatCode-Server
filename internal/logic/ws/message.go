package ws

import (
	"encoding/json"
	"fmt"

	"github.com/hubogle/chatcode-server/internal/code"
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
	SessionType_Group  SessionType = 2 // 群聊
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
	ReceiverID  uint64      `json:"receiver_uid"` // 接收者id, 群组id 或 用户 id
	SenderID    uint64      `json:"sender_uid"`   // 发送者id
	MessageType MessageType `json:"message_type"` // 消息类型
	Content     []byte      `json:"content"`      // 消息内容
	SendTime    uint64      `json:"send_time"`    // 消息发送时间
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
	Msg  string  `json:"msg"`
	Data string  `json:"data"`
}

// SendToUser 对指定 user id 的用户发送消息
func SendToUser(msg *Message, userID uint64) (uint64, error) {
	// TODO 保存数据到数据库

	conn := ConnManager.GetConn(userID)
	if conn == nil {
		fmt.Println("【消息】接收者不在线")
	}
	conn.SendMsg(userID, msg.Content)

	return 0, nil
}

// SendToGroup 对指定 group id 的群组发送消息
func SendToGroup() {
}

func MockServerMessage(msgType MsgType, statusCode int, data string) []byte {
	msg := &ServerMessage{
		Type: msgType,
		Code: statusCode,
		Msg:  code.GetCoder(statusCode).String(),
		Data: data,
	}
	result, _ := json.Marshal(msg)
	return result
}
