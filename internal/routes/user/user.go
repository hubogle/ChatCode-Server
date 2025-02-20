// Code generated by goctl. DO NOT EDIT.
package user

import (
	baseHandler "github.com/hubogle/chatcode-server/internal/handler"
	handler "github.com/hubogle/chatcode-server/internal/handler/user"
	logic "github.com/hubogle/chatcode-server/internal/logic/user"
	"github.com/hubogle/chatcode-server/internal/repository"
	"github.com/hubogle/chatcode-server/internal/svc"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(e *gin.Engine, svcCtx *svc.ServiceContext) {
	userRepo := repository.NewUserRepo(svcCtx)
	userLogic := logic.NewUserLogic(svcCtx, userRepo)
	baseHandler := baseHandler.NewHandler(svcCtx)
	userHandler := handler.NewUserHandler(svcCtx, baseHandler, userLogic)

	e.POST("/api/v1/login", userHandler.UserLoginHandler)

	e.POST("/api/v1/logout", userHandler.UserLogoutHandler)

	e.POST("/api/v1/register", userHandler.UserRegisterHandler)
}
