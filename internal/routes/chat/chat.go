// Code generated by goctl. DO NOT EDIT.
package chat

import (
	baseHandler "github.com/hubogle/chatcode-server/internal/handler"
	handler "github.com/hubogle/chatcode-server/internal/handler/chat"
	logic "github.com/hubogle/chatcode-server/internal/logic/chat"
	"github.com/hubogle/chatcode-server/internal/middleware"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoute(e *gin.Engine, svcCtx *svc.ServiceContext) {
	chatRepo := repository.NewChatRepo(svcCtx)
	chatLogic := logic.NewChatLogic(svcCtx, chatRepo)
	baseHandler := baseHandler.NewHandler(svcCtx)
	chatHandler := handler.NewChatHandler(svcCtx, baseHandler, chatLogic)

	e.GET("/api/v1/chat/list", middleware.Auth, chatHandler.ChatListHandler)

	e.POST("/api/v1/chat/private", middleware.Auth, chatHandler.ChatPrivateHandler)

	e.POST("/api/v1/chat/room", middleware.Auth, chatHandler.ChatRoomHandler)
}
