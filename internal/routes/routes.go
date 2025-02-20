// Code generated by goctl. DO NOT EDIT.
package routes

import (
	"github.com/hubogle/chatcode-server/internal/routes/chat"
	"github.com/hubogle/chatcode-server/internal/routes/friend"
	"github.com/hubogle/chatcode-server/internal/routes/home"
	"github.com/hubogle/chatcode-server/internal/routes/ping"
	"github.com/hubogle/chatcode-server/internal/routes/room"
	"github.com/hubogle/chatcode-server/internal/routes/user"
	"github.com/hubogle/chatcode-server/internal/routes/ws"
	"github.com/hubogle/chatcode-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine, svcCtx *svc.ServiceContext) {
	chat.RegisterChatRoute(e, svcCtx)

	friend.RegisterFriendRoute(e, svcCtx)

	home.RegisterHomeRoute(e, svcCtx)

	ping.RegisterPingRoute(e, svcCtx)

	room.RegisterRoomRoute(e, svcCtx)

	user.RegisterUserRoute(e, svcCtx)

	ws.RegisterWsRoute(e, svcCtx)
}
