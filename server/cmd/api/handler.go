package main

import (
	"go.uber.org/zap"
	"qiniu-1024-server/service"
)

type Handler struct {
	srv   *service.Service
	log   *zap.Logger
	sugar *zap.SugaredLogger
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv:   srv,
		log:   srv.Logger,
		sugar: srv.Logger.Sugar(),
	}
}
