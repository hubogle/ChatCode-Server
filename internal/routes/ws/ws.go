// Code generated by goctl. DO NOT EDIT.
package ws

import (
	baseHandler "github.com/hubogle/chatcode-server/internal/handler"
	handler "github.com/hubogle/chatcode-server/internal/handler/ws"
	logic "github.com/hubogle/chatcode-server/internal/logic/ws"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func RegisterWsRoute(e *gin.Engine, svcCtx *svc.ServiceContext) {
	wsRepo := repository.NewWsRepo(svcCtx)
	wsLogic := logic.NewWsLogic(svcCtx, wsRepo)
	baseHandler := baseHandler.NewHandler(svcCtx)
	wsHandler := handler.NewWsHandler(svcCtx, baseHandler, wsLogic)

	e.GET("/api/v1/ws", wsHandler.WsHandler)
}
