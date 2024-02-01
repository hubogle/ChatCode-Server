package user

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/user"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IUserHandler interface {
	UserLoginHandler(ctx *gin.Context)
	UserLogoutHandler(ctx *gin.Context)
	UserRegisterHandler(ctx *gin.Context)
}

func NewUserHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic user.IUserLogic) IUserHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic user.IUserLogic
}
