package chat

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/chat"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IChatHandler interface {
	ChatPrivateHandler(ctx *gin.Context)
	ChatRoomHandler(ctx *gin.Context)
}

func NewChatHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic chat.IChatLogic) IChatHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic chat.IChatLogic
}
