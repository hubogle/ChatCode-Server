package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/chat"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IChatHandler interface {
	ChatListHandler(ctx *gin.Context)
	ChatPrivateHandler(ctx *gin.Context)
	ChatRoomHandler(ctx *gin.Context)
}

func NewChatHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic chat.IChatLogic) IChatHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  512,
			WriteBufferSize: 512,
			// CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic    chat.IChatLogic
	upgrader websocket.Upgrader
}
