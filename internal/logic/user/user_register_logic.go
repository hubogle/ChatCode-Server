package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hubogle/chatcode-server/internal/dal/model"
	"github.com/hubogle/chatcode-server/internal/types/user"
)

// UserRegister user register
// POST /api/v1/register
func (l *logic) UserRegister(ctx context.Context, req *user.LoginReq) (err error) {
	now := time.Now().Unix()
	userBasic := &model.UserBasic{
		UID:       uint64(uuid.New().ID()),
		Account:   req.Account,
		Password:  req.Password,
		Nickname:  req.Account,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = l.repo.InsertOneUserBasic(ctx, userBasic)
	if err != nil {
		return err
	}
	return
}
