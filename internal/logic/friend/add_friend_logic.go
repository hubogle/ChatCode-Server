package friend

import (
	"context"
	"time"

	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/friend"
)

// AddFriend add friend
// POST /api/v1/friend/add
func (l *logic) AddFriend(ctx context.Context, uc *jwt.UserClaims, req *friend.AddFriendReq) (err error) {
	now := time.Now().Unix()
	userFriend := &model.UserFriend{
		UserID:    uc.UID,
		FriendID:  req.Uid,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = l.repo.InsertOneUserFriend(ctx, userFriend)
	if err != nil {
		return err
	}
	return
}
