package room

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/room"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IRoomHandler interface {
	RoomHandler(ctx *gin.Context)
}

func NewRoomHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic room.IRoomLogic) IRoomHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic room.IRoomLogic
}
