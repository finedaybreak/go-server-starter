package router

import (
	"go-server-starter/internal/enum"
)

func (r *Router) SetupUserRoutes() {
	router := r.router.Group("/user")
	router.Use(r.jwt.JWT())
	{
		router.GET("/info", r.handler.User().GetInfo)
		router.PUT("/info", r.handler.User().UpdateInfo)
		// Admin User
		router.GET("/admin/table", r.auth.RoleCheckAny(enum.RoleCodeAdmin, enum.RoleCodeSuperAdmin), r.handler.User().GetTable)
	}
}
