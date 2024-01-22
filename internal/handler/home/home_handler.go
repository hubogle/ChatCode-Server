package home

import (
	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/response"
)

// Home home page
// GET /api/v1/home
func (h *handler) HomeHandler(c *gin.Context) {
	err := h.logic.Home(c)
	response.HandleResponse(c, nil, err)
}
