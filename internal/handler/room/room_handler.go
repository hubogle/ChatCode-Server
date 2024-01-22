package room

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/room"

	"github.com/gin-gonic/gin"
)

// Room room page
// GET /api/v1/room/:id
func (h *handler) RoomHandler(c *gin.Context) {
	var req room.RoomReq
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	err := h.logic.Room(c, &req)
	response.HandleResponse(c, nil, err)
}
