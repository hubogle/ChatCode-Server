package chat

import (
	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatRoom room chat
// POST /api/v1/chat/room
func (h *handler) ChatRoomHandler(c *gin.Context) {
	var req chat.ChatRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.HandleResponse(c, nil, code.WithCodeMsg(code.ChatWsNewErr))
		return
	}
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	err = h.logic.ChatRoom(c, ws, uc, &req)
	response.HandleResponse(c, nil, err)
}
