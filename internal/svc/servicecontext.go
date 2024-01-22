package svc

import (
	"github.com/hubogle/chatcode-server/config"
	"github.com/hubogle/chatcode-server/internal/dal"
	"github.com/hubogle/chatcode-server/internal/dal/query"
	"github.com/hubogle/chatcode-server/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func NewServiceContext(cfg config.ServerConfig) *ServiceContext {
	dbIns, err := dal.GetMySQLFactoryOr(cfg.Mysql)
	if err != nil {
		zap.S().Errorw("failed to get mysql db", "err", err)
	}
	query.SetDefault(dbIns.GetDb())

	logger, err := log.NewLogger()
	if err != nil {
		zap.S().Errorw("failed to get logger", "err", err)
	}
	defer logger.Sync()

	return &ServiceContext{
		Db:  dbIns.GetDb(),
		Log: logger,
	}
}
