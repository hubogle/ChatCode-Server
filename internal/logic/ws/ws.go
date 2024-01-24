package ws

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"
)

//go:generate mockgen -source=$GOFILE -destination ../../mocks/logic/$GOFILE -package mock_logic

type IWsLogic interface {
	Ws(ctx context.Context, conn *websocket.Conn)
}

func NewWsLogic(logicSvc *svc.ServiceContext, repo repository.WsRepo) IWsLogic {
	manager := GetManager()
	manager.StartWorkerPool()

	return &logic{
		ServiceContext: logicSvc,
		repo:           repo,
		manager:        manager,
	}
}

type logic struct {
	*svc.ServiceContext
	repo    repository.WsRepo
	manager *Manager
}
