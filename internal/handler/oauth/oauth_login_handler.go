package oauth

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/oauth"

	"github.com/gin-gonic/gin"
)

// OauthLogin user login
// POST /api/v1/login
func (h *handler) OauthLoginHandler(c *gin.Context) {
	var req oauth.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	resp, err := h.logic.OauthLogin(c, &req)
	response.HandleResponse(c, resp, err)
}
