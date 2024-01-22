package dal

import (
	"fmt"
	"sync"

	"github.com/hubogle/chatcode-server/config"
	"github.com/hubogle/chatcode-server/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repo interface {
	GetDb() *gorm.DB
	CloseDb() error
}

type dbRepo struct {
	Db *gorm.DB
}

func (d *dbRepo) GetDb() *gorm.DB {
	return d.Db
}

func (d *dbRepo) CloseDb() error {
	db, err := d.Db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory Repo
	mysqlOnce    sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr(opts *config.MySql) (Repo, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	mysqlOnce.Do(func() {
		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)
		opt := db.DbOpts(*opts)
		dbIns, err = db.NewMySQL(&opt)

		mysqlFactory = &dbRepo{dbIns}
	})

	if err != nil {
		return mysqlFactory, fmt.Errorf("failed to get mysql db")
	}

	return mysqlFactory, nil
}
