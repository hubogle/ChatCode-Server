package handler

import "github.com/hubogle/chatcode-server/internal/svc"

type Handler interface{}

type BaseHandler struct {
	*svc.ServiceContext
}

func NewHandler(handlerSvc *svc.ServiceContext) Handler {
	return &BaseHandler{
		ServiceContext: handlerSvc,
	}
}
