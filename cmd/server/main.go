package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/config"
	"github.com/hubogle/chatcode-server/internal/routes"
	"github.com/hubogle/chatcode-server/internal/svc"
	"github.com/hubogle/chatcode-server/pkg/log"
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
	log.NewSlog("json", -4)
	gin.SetMode(cfg.App.Env)
	r := gin.New()
	r.Use(middleware.GinSlog(), middleware.RecoverySlog(true))
	routes.Setup(r, svc.NewServiceContext(cfg))

	server := &http.Server{
		Addr:    cfg.App.Addr,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown:", err)
	}
	slog.Info("Server exiting")
}
