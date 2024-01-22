package ping

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/ping"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IPingHandler interface {
	PingHandler(ctx *gin.Context)
}

func NewPingHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic ping.IPingLogic) IPingHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic ping.IPingLogic
}
