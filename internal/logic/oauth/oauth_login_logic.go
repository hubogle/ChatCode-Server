package oauth

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/oauth"
)

// OauthLogin user login
// POST /api/v1/login
func (l *logic) OauthLogin(ctx context.Context, req *oauth.LoginReq) (err error) {
	userBasic, err := l.repo.GetUserBasicByAccountPassword(ctx, req.Account, req.Password)
	if err != nil {
		l.Log.Sugar().Errorw("failed to get user basic", "err", err)
		return code.WithCodeMsg(code.LoginUserNotExist, "")
	}
	token, err := jwt.GenerateToken(userBasic.UID, userBasic.Account)
	if err != nil {
		return code.WithCodeMsg(code.LoginFailed, "")
	}

	_ = token

	return
}
