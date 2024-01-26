package ws

import (
	"encoding/json"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/pkg/errors"
)

type (
	Handler func() error
	Job     struct {
		conn    *Connection     // 链接
		message json.RawMessage // 原始数据
		f       Handler         // 处理函数
	}
)

// Login user login
func (p *Job) Login() error {
	if p.conn.GetUserID() != 0 {
		return errors.New("user already login")
	}

	msg := &LoginMessage{}
	err := json.Unmarshal(p.message, msg)
	if err != nil {
		return errors.WithMessage(err, "unmarshal message error")
	}

	uc, err := jwt.ParseToken(msg.Token)
	if err != nil {
		return errors.WithMessage(err, "token parse error")
	}

	p.conn.SetUserID(uc.UID)
	p.conn.Manager.AddConn(uc.UID, p.conn)
	bytes := MockServerMessage(Msg_Type_login, code.Success, "ok")
	p.conn.SendMsg(uc.UID, string(bytes))
	return nil
}

func (p *Job) HeartBeat() {
}

// Message 处理客户端发送给服务端的消息，包括：单聊，群聊
func (p *Job) Message() error {
	msg := &Message{}
	userID := p.conn.GetUserID() // send message user id
	err := json.Unmarshal(p.message, msg)
	if err != nil {
		return errors.WithMessage(err, "unmarshal message error")
	}

	if msg.SenderID != userID {
		return errors.New("sender id error")
	}

	if msg.SessionType == SessionType_Single && msg.ReceiverID == userID {
		return errors.New("receiver id error")
	}

	// if err = SendToUser(msg, msg.SenderID); err != nil {
	// 	return errors.WithMessage(err, "send to myself error")
	// }

	switch msg.SessionType {
	case SessionType_Single:
		if err = SendToUser(msg, msg.ReceiverID); err != nil {
			return errors.WithMessage(err, "send to user error")
		}
	case SessionType_Room:
		if err = SendToRoom(msg, msg.ReceiverID); err != nil {
			return errors.WithMessage(err, "send to room error")
		}
	default:
		return errors.New("session type error")
	}
	return nil
}
