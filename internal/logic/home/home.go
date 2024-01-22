package home

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IHomeLogic interface {
	Home(ctx context.Context) (err error)
}

func NewHomeLogic(logicSvc *svc.ServiceContext, repo repository.HomeRepo) IHomeLogic {
	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
	}
}

type logic struct {
	*svc.ServiceContext
	repo repository.HomeRepo
}
