package chat

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatPrivate private chat
// POST /api/v1/chat/private
func (h *handler) ChatPrivateHandler(c *gin.Context) {
	var req chat.ChatPrivateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	err := h.logic.ChatPrivate(c, &req)
	response.HandleResponse(c, nil, err)
}
