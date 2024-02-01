package friend

import (
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/response"
	"github.com/hubogle/chatcode-server/internal/types/friend"

	"github.com/gin-gonic/gin"
)

// AddFriend add friend
// POST /api/v1/friend/add
func (h *handler) AddFriendHandler(c *gin.Context) {
	var req friend.AddFriendReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandlerParamsResponse(c, err)
		return
	}
	uc := c.MustGet("user_claims").(*jwt.UserClaims)

	err := h.logic.AddFriend(c, uc, &req)
	response.HandleResponse(c, nil, err)
}
