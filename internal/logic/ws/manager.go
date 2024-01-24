package ws

import (
	"fmt"
	"sync"
)

var (
	ConnManager *Manager
	once        sync.Once
)

// Manager 连接管理
// 1. 连接管理 2.工作队列
type Manager struct {
	connMap  sync.Map    // user id -> *Connection
	jobQueue []chan *Job // 任务处理队列
}

// GetManager 获取连接管理器
func GetManager() *Manager {
	once.Do(func() {
		ConnManager = &Manager{
			connMap:  sync.Map{},
			jobQueue: make([]chan *Job, 3),
		}
	})
	return ConnManager
}

func (m *Manager) Stop() {
}

// StartWorkerPool 启动工作池
func (m *Manager) StartWorkerPool() {
	for i := 0; i < len(m.jobQueue); i++ {
		m.jobQueue[i] = make(chan *Job, 10) // 每个队列处理的任务数量
		go func(i int) {
			fmt.Println("启动第", i, "个消息队列")
			for {
				select {
				case process := <-m.jobQueue[i]:
					fmt.Println("第", i, "个消息队列接收到任务")
					process.f()
				}
			}
		}(i)
	}
}

// SendTaskToJobQueue 发送任务到工作队列
func (m *Manager) SendTaskToJobQueue(job *Job) {
	if len(m.jobQueue) > 0 {
		// 根据链接id分配工作队列，保证同一个链接的任务在同一个工作队列中处理
		workID := job.conn.ConnID % uint64(len(m.jobQueue))
		m.jobQueue[workID] <- job
	} else {
		go job.f() // 会导致消息乱序
	}
}

func (m *Manager) AddConn(userID uint64, conn *Connection) {
	m.connMap.Store(conn.UserID, conn)
}

func (m *Manager) RemoveConn(userID uint64) {
	m.connMap.Delete(userID)
}

func (m *Manager) GetConn(userID uint64) *Connection {
	conn, ok := m.connMap.Load(userID)
	if !ok {
		return nil
	}
	return conn.(*Connection)
}

// GetConnAll 获取所有连接
func (m *Manager) GetConnAll() []*Connection {
	var connList []*Connection
	m.connMap.Range(func(key, value interface{}) bool {
		connList = append(connList, value.(*Connection))
		return true
	})
	return connList
}
