package config

import "time"

// 应用信息
type app struct {
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

// MySQL信息
type MySql struct {
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

// ServerConfig 配置信息
type ServerConfig struct {
	App   *app   `yaml:"app"`
	Mysql *MySql `yaml:"mysql"`
}
