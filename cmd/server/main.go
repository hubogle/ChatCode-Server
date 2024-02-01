package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/config"
	"github.com/hubogle/chatcode-server/internal/routes"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/pkg/middleware"
	"github.com/spf13/viper"
)

var (
	cfg        config.ServerConfig
	configPath string
)

func InitConfig() {
	flag.StringVar(&configPath, "c", "", "配置文件的路径")
	flag.Parse()

	if configPath == "" {
		workDir, _ := os.Getwd()
		configPath = workDir + "/config"
	}
	fmt.Println(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
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
	r.Use(middleware.Ginzap(time.RFC3339, false))
	routes.Setup(r, svc.NewServiceContext(cfg))

	r.Run(cfg.App.Addr)
	log.Println("监听端口:", cfg.App.Addr)
}
