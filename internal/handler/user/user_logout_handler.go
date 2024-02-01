package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/response"
)

// UserLogout user logout
// POST /api/v1/logout
func (h *handler) UserLogoutHandler(c *gin.Context) {

	err := h.logic.UserLogout(c)
	response.HandleResponse(c, nil, err)
}
