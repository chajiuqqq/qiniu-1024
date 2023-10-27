package main

import (
	"go.uber.org/zap"
	"qiniu-1024-server/service"
)

type Handler struct {
	srv *service.Service
	log *zap.Logger
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
		log: srv.Logger,
	}
}
