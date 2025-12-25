package handler

import (
	"go-server-starter/internal/service"

	"go.uber.org/zap"
)

type Handler interface {
	Hello() HelloHandler
	User() UserHandler
	UserRole() UserRoleHandler
	Auth() AuthHandler
}

type HandlerImpl struct {
	logger          *zap.Logger
	helloHandler    HelloHandler
	userHandler     UserHandler
	userRoleHandler UserRoleHandler
	authHandler     AuthHandler
}

func NewHandler(service service.Service, logger *zap.Logger) Handler {
	return &HandlerImpl{
		logger:          logger,
		helloHandler:    NewHelloHandler(logger),
		userHandler:     NewUserHandler(logger, service),
		userRoleHandler: NewUserRoleHandler(logger, service),
		authHandler:     NewAuthHandler(logger, service),
	}
}

func (h *HandlerImpl) Hello() HelloHandler {
	return h.helloHandler
}

func (h *HandlerImpl) User() UserHandler {
	return h.userHandler
}

func (h *HandlerImpl) UserRole() UserRoleHandler {
	return h.userRoleHandler
}

func (h *HandlerImpl) Auth() AuthHandler {
	return h.authHandler
}
