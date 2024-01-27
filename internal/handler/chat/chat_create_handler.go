package chat

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/chat"

	"github.com/gin-gonic/gin"
)

// ChatCreate create chat
// POST /api/v1/chat/create
func (h *handler) ChatCreateHandler(c *gin.Context) {
	var req chat.ChatCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	err := h.logic.ChatCreate(c, uc, &req)
	response.HandleResponse(c, nil, err)
}
