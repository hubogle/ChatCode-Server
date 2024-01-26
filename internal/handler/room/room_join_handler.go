package room

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/room"

	"github.com/gin-gonic/gin"
)

// RoomJoin join room
// POST /api/v1/room/join
func (h *handler) RoomJoinHandler(c *gin.Context) {
	var req room.RoomJoinReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	err := h.logic.RoomJoin(c, uc, &req)
	response.HandleResponse(c, nil, err)
}
