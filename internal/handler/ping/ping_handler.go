package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/response"
)

// Ping ping
// GET /ping
func (h *handler) PingHandler(c *gin.Context) {
	err := h.logic.Ping(c)
	response.HandleResponse(c, nil, err)
}
