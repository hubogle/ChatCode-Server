package chat

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatList chat list
// GET /api/v1/chat/list
func (h *handler) ChatListHandler(c *gin.Context) {
	var req chat.ChatListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	err := h.logic.ChatList(c, uc, &req)
	response.HandleResponse(c, nil, err)
}
