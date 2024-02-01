package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQL信息
type DbOpts struct {
	Host                  string        `yaml:"host"`
	Port                  int           `yaml:"port"`
	User                  string        `yaml:"user"`
	Password              string        `yaml:"password"`
	DataBase              string        `yaml:"database"`
	MaxIdleConnections    int           `yaml:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `yaml:"max-open-connections,omitempty"`
	MaxConnectionLifeTime time.Duration `yaml:"max-connection-life-time,omitempty"`
	LogLevel              int           `yaml:"log-level"`
}

func NewMySQL(opts *DbOpts) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%t&loc=%s&timeout=10s`,
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.DataBase,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.LogLevel(opts.LogLevel),
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      true,
				Colorful:                  true,
			}),
	})
	if err != nil {
		log.Fatalf("failed to connect database, err: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	return db, nil
}
