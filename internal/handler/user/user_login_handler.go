package user

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/user"

	"github.com/gin-gonic/gin"
)

// UserLogin user login
// POST /api/v1/login
func (h *handler) UserLoginHandler(c *gin.Context) {
	var req user.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	resp, err := h.logic.UserLogin(c, &req)
	response.HandleResponse(c, resp, err)
}
