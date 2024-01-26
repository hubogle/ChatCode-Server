package room

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/room"

	"github.com/gin-gonic/gin"
)

// RoomCreate create room
// POST /api/v1/room/create
func (h *handler) RoomCreateHandler(c *gin.Context) {
	var req room.RoomCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	resp, err := h.logic.RoomCreate(c, uc, &req)
	response.HandleResponse(c, resp, err)
}
