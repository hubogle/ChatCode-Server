// Code generated by goctl. DO NOT EDIT.
package room

import (
	baseHandler "github.com/hubogle/chatcode-server/internal/handler"
	handler "github.com/hubogle/chatcode-server/internal/handler/room"
	logic "github.com/hubogle/chatcode-server/internal/logic/room"
	"github.com/hubogle/chatcode-server/internal/middleware"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoute(e *gin.Engine, svcCtx *svc.ServiceContext) {
	roomRepo := repository.NewRoomRepo(svcCtx)
	roomLogic := logic.NewRoomLogic(svcCtx, roomRepo)
	baseHandler := baseHandler.NewHandler(svcCtx)
	roomHandler := handler.NewRoomHandler(svcCtx, baseHandler, roomLogic)

	e.POST("/api/v1/room/create", middleware.Auth, roomHandler.RoomCreateHandler)

	e.POST("/api/v1/room/join", middleware.Auth, roomHandler.RoomJoinHandler)

	e.GET("/api/v1/room/:id", middleware.Auth, roomHandler.RoomGetHandler)

	e.GET("/api/v1/room/:id/person", middleware.Auth, roomHandler.RoomPersonHandler)
}
