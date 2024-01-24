package ws

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/ws"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IWsHandler interface {
	WsHandler(ctx *gin.Context)
}

func NewWsHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic ws.IWsLogic) IWsHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic ws.IWsLogic
}
