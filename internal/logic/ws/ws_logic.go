package ws

import (
	"context"

	"github.com/gorilla/websocket"
)

var connID uint64

// Ws websocket
// GET /api/v1/ws
func (l *logic) Ws(ctx context.Context, wsConn *websocket.Conn) {
	conn := NewConnection(l.manager, wsConn, connID)
	connID++

	go conn.Reader()
	go conn.Writer()

	return
}
