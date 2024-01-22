package oauth

import (
	"github.com/gin-gonic/gin"
	hdl "github.com/hubogle/chatcode-server/internal/handler"
	"github.com/hubogle/chatcode-server/internal/logic/oauth"
	"github.com/hubogle/chatcode-server/internal/svc"
)

type IOauthHandler interface {
	OauthLoginHandler(ctx *gin.Context)
	OauthLogoutHandler(ctx *gin.Context)
}

func NewOauthHandler(handlerSvc *svc.ServiceContext, hdl hdl.Handler, logic oauth.IOauthLogic) IOauthHandler {
	return &handler{
		Handler:        hdl,
		ServiceContext: handlerSvc,
		logic:          logic,
	}
}

type handler struct {
	hdl.Handler
	*svc.ServiceContext
	logic oauth.IOauthLogic
}
