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
	Host                  string        `yaml:"host" mapstructure:"host"`
	Port                  int           `yaml:"port" mapstructure:"port"`
	User                  string        `yaml:"user" mapstructure:"user"`
	Password              string        `yaml:"password" mapstructure:"password"`
	DataBase              string        `yaml:"database" mapstructure:"database"`
	MaxIdleConnections    int           `yaml:"max-idle-connections,omitempty" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `yaml:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `yaml:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `yaml:"log-level" mapstructure:"log-level"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App   *app   `yaml:"app"`
	Mysql *MySql `yaml:"mysql"`
}
