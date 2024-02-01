package user

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/user"
	"go.uber.org/zap"
)

// UserLogin user login
// POST /api/v1/login
func (l *logic) UserLogin(ctx context.Context, req *user.LoginReq) (resp user.LoginResp, err error) {
	userBasic, err := l.repo.GetUserBasicByAccountPassword(ctx, req.Account, req.Password)
	if err != nil {
		l.Log.Error("failed to get user basic", zap.Error(err))
		return resp, code.WithCodeMsg(code.LoginUserNotExist, "")
	}
	token, err := jwt.GenerateToken(userBasic.UID, userBasic.Account)
	if err != nil {
		return resp, code.WithCodeMsg(code.LoginFailed, "")
	}

	resp.Token = token
	resp.Uid = userBasic.UID

	return
}
