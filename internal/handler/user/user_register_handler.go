package user

import (
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/user"

	"github.com/gin-gonic/gin"
)

// UserRegister user register
// POST /api/v1/register
func (h *handler) UserRegisterHandler(c *gin.Context) {
	var req user.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}

	err := h.logic.UserRegister(c, &req)
	response.HandleResponse(c, nil, err)
}
