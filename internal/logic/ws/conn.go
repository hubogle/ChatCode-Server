package ws

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ConnID      uint64          // 链接id
	UserID      uint64          // 所属链接用户
	UserIDMutex sync.RWMutex    // 所属链接用户读写锁
	Socket      *websocket.Conn // 链接
	Manager     *Manager        // 链接管理器

	sendCh chan []byte   // 发送消息通道
	exitCh chan struct{} // 退出通道

	isClose      bool         // 是否关闭
	isCloseMutex sync.RWMutex // 是否关闭读写锁

	lastHeartBeatTime int64        // 最后一次心跳时间
	HeartMutex        sync.RWMutex // 最后一次心跳时间读写锁
}

func NewConnection(manager *Manager, conn *websocket.Conn, connID uint64) *Connection {
	return &Connection{
		ConnID:  connID,
		UserID:  0,
		Socket:  conn,
		Manager: manager,

		sendCh: make(chan []byte, 10),
		exitCh: make(chan struct{}, 1),

		isClose:           false,
		lastHeartBeatTime: time.Now().Unix(),
	}
}

// Stop 关闭链接
func (c *Connection) Stop() {
	c.isCloseMutex.Lock()
	defer c.isCloseMutex.Unlock()

	if c.isClose {
		return
	}

	_ = c.Socket.Close()
	c.exitCh <- struct{}{}
	// TODO 用户下线判断

	c.isClose = true

	close(c.exitCh)
	close(c.sendCh)
}

// KeepLive heartbeat connection
func (c *Connection) KeepLive() {
	c.HeartMutex.Lock()
	defer c.HeartMutex.Unlock()

	c.lastHeartBeatTime = time.Now().Unix()
}

// IsAlive check connection is alive
func (c *Connection) IsAlive() bool {
	now := time.Now().Unix()
	c.isCloseMutex.Lock()
	defer c.isCloseMutex.Unlock()
	c.HeartMutex.RLock()
	defer c.HeartMutex.RUnlock()

	if c.isClose || now-c.lastHeartBeatTime > 10 { // 心跳 10s
		return false
	}
	return true
}

// Reader client message reader
func (c *Connection) Reader() {
	defer c.Stop()

	for {
		_, messageRaw, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("read message error: ", err)
			return
		}
		message := &ClientMessage{}
		err = json.Unmarshal(messageRaw, message)
		if err != nil {
			fmt.Println("unmarshal message error: ", err)
			return
		}

		job := &Job{
			conn:    c,
			message: message.Data,
			f:       nil,
		}

		switch message.Type {
		case Msg_Type_login:
			job.f = job.Login
		case Msg_Type_Heart:
		case Msg_Type_Message:
			job.f = job.Message
		default:
			fmt.Println("unknown message type")
			return
		}

		c.KeepLive()
		c.Manager.SendTaskToJobQueue(job)
	}
}

// Writer client message writer
func (c *Connection) Writer() {
	for {
		select {
		case message := <-c.sendCh:
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("write message error: ", err)
				return
			}
		case <-c.exitCh:
			return
		}
	}
}

// SendMsg 对指定 user id 的用户发送消息
func (c *Connection) SendMsg(userID uint64, message []byte) {
	c.isCloseMutex.RLock()
	defer c.isCloseMutex.RUnlock()

	if c.isClose {
		return
	}

	conn := c.Manager.GetConn(userID)
	if conn == nil {
		return
	}

	conn.sendCh <- message

	return
}

// GetUserID 获取链接所属用户id
func (c *Connection) GetUserID() uint64 {
	c.UserIDMutex.RLock()
	defer c.UserIDMutex.RUnlock()

	return c.UserID
}

func (c *Connection) SetUserID(userID uint64) {
	c.UserIDMutex.Lock()
	defer c.UserIDMutex.Unlock()

	c.UserID = userID
}
