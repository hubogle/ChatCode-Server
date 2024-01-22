package chat

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatRoom room chat
// POST /api/v1/root/chat
func (h *handler) ChatRoomHandler(c *gin.Context) {
	var req chat.ChatRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	err := h.logic.ChatRoom(c, &req)
	response.HandleResponse(c, nil, err)
}
