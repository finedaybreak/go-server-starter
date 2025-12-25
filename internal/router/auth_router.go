package router

func (r *Router) SetupAuthRoutes() {
	r.router.POST("/auth/login/by-mobile-and-code", r.handler.Auth().LoginByMobileAndCode)
	r.router.POST("/auth/login/by-email-and-code", r.handler.Auth().LoginByEmailAndCode)
}
