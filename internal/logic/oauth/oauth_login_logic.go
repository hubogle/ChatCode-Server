package oauth

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/code"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
	"github.com/hubogle/chatcode-server/internal/types/oauth"
	"go.uber.org/zap"
)

// OauthLogin user login
// POST /api/v1/login
func (l *logic) OauthLogin(ctx context.Context, req *oauth.LoginReq) (resp oauth.LoginResp, err error) {
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

	return
}
