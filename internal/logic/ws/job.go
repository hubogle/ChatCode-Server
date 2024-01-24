package ws

import (
	"encoding/json"
	"fmt"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
)

type (
	Handler func()
	Job     struct {
		conn    *Connection     // 链接
		message json.RawMessage // 原始数据
		f       Handler         // 处理函数
	}
)

// 对应每种消息类型的处理函数
func (p *Job) Login() {
	if p.conn.GetUserID() != 0 {
		fmt.Println("【登录】用户已登录")
		return
	}

	msg := &LoginMessage{}
	err := json.Unmarshal(p.message, msg)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	uc, err := jwt.ParseToken(msg.Token)
	if err != nil {
		fmt.Println("【登录】token 解析失败", err)
		return
	}

	p.conn.SetUserID(uc.UID)
	p.conn.Manager.AddConn(uc.UID, p.conn)
	bytes := MockServerMessage(Msg_Type_login, code.Success, "ok")
	p.conn.SendMsg(uc.UID, bytes)
}

func (p *Job) HeartBeat() {
}

// Message 处理客户端发送给服务端的消息，包括：单聊，群聊
func (p *Job) Message() {
	msg := &Message{}
	userID := p.conn.GetUserID() // send message user id
	err := json.Unmarshal(p.message, msg)
	if err != nil {
		return
	}

	if msg.SenderID != userID {
		fmt.Println("【消息】发送有误")
		return
	}

	if msg.SessionType == SessionType_Single && msg.ReceiverID == userID {
		fmt.Println("【消息】接收着有误")
		return
	}

	switch msg.SessionType {
	case SessionType_Single:
	case SessionType_Group:
	default:
		fmt.Println("【消息】会话类型有误")
		return
	}
}
