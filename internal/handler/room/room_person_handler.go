package room

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/room"

	"github.com/gin-gonic/gin"
)

// RoomPerson room person
// GET /api/v1/room/:id/person
func (h *handler) RoomPersonHandler(c *gin.Context) {
	var req room.RoomPersonReq
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	resp, err := h.logic.RoomPerson(c, &req)
	response.HandleResponse(c, resp, err)
}
