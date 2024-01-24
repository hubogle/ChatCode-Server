package oauth

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/oauth"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IOauthLogic interface {
	OauthLogin(ctx context.Context, req *oauth.LoginReq) (resp oauth.LoginResp, err error)
	OauthLogout(ctx context.Context) (err error)
}

func NewOauthLogic(logicSvc *svc.ServiceContext, repo repository.OauthRepo) IOauthLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.OauthRepo
}
