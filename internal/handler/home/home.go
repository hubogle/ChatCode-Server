package home

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/home"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IHomeHandler interface {
	HomeHandler(ctx *gin.Context)
}

func NewHomeHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic home.IHomeLogic) IHomeHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic home.IHomeLogic
}
