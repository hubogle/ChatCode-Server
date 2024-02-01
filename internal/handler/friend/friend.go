package friend

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/friend"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IFriendHandler interface {
	AddFriendHandler(ctx *gin.Context)
}

func NewFriendHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic friend.IFriendLogic) IFriendHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic friend.IFriendLogic
}
