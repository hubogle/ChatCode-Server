package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
)

// ChatList chat list
// GET /api/v1/chat/list
func (h *handler) ChatListHandler(c *gin.Context) {
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	resp, err := h.logic.ChatList(c, uc)

	response.HandleResponse(c, resp, err)
}
