package ping

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IPingLogic interface {
	Ping(ctx context.Context) (err error)
}

func NewPingLogic(logicSvc *svc.ServiceContext, repo repository.PingRepo) IPingLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.PingRepo
}
