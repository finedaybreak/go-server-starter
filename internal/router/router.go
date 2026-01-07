package router

import (
	"go-server-starter/internal/handler"
	"go-server-starter/internal/middleware"
	"go-server-starter/pkg/auth"
	"go-server-starter/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler   handler.Handler
	router    *gin.RouterGroup
	jwt       *jwt.JWT
	auth      auth.Auth
	ratelimit *middleware.RateLimit
}

func NewRouter(handler handler.Handler, router *gin.RouterGroup, jwt *jwt.JWT, auth auth.Auth, ratelimit *middleware.RateLimit) *Router {
	return &Router{handler: handler, router: router, jwt: jwt, auth: auth, ratelimit: ratelimit}
}

func (r *Router) SetupRoutes() {
	// Hello (公开接口)
	r.router.GET("/hello", r.handler.Hello().Hello)

	// Auth - 认证相关
	r.SetupAuthRoutes()

	// User - 用户相关
	r.SetupUserRoutes()
}
