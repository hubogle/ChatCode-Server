package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/response"
)

// OauthLogout user logout
// POST /api/v1/logout
func (h *handler) OauthLogoutHandler(c *gin.Context) {
	err := h.logic.OauthLogout(c)
	response.HandleResponse(c, nil, err)
}
