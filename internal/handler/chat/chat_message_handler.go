package chat

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatMessage chat message list
// GET /api/v1/chat/message/list
func (h *handler) ChatMessageHandler(c *gin.Context) {
	var req chat.ChatMessageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	resp, err := h.logic.ChatMessage(c, uc, &req)
	response.HandleResponse(c, resp, err)
}
