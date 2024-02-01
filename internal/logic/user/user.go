package user

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/internal/types/user"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IUserLogic interface {
	UserLogin(ctx context.Context, req *user.LoginReq) (resp user.LoginResp, err error)
	UserLogout(ctx context.Context) (err error)
	UserRegister(ctx context.Context, req *user.LoginReq) (err error)
}

func NewUserLogic(logicSvc *svc.ServiceContext, repo repository.UserRepo) IUserLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.UserRepo
}
