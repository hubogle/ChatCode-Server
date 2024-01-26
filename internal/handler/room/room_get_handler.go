package room

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/room"

	"github.com/gin-gonic/gin"
)

// RoomGet room get
// GET /api/v1/room/:id
func (h *handler) RoomGetHandler(c *gin.Context) {
	var req room.RoomGetReq
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	err := h.logic.RoomGet(c, &req)
	response.HandleResponse(c, nil, err)
}
