package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/config"
	"github.com/hubogle/chatcode-server/internal/routes"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/spf13/viper"
)

var cfg config.ServerConfig

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n", err))
	}
}

func main() {
	InitConfig()
	gin.SetMode(cfg.App.Env)
	r := gin.New()
	routes.Setup(r, svc.NewServiceContext(cfg))

	r.Run(cfg.App.Addr)
	log.Println("监听端口:", cfg.App.Addr)
}
