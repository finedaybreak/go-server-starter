package auth

import (
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/exception"
	"go-server-starter/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleCheckType string

const (
	RoleCheckTypeAll RoleCheckType = "all" // 所有角色都符合
	RoleCheckTypeAny RoleCheckType = "any" // 任意一个角色符合
)

type Auth interface {
	RoleCheck(roleCheckType RoleCheckType, roles ...enum.RoleCode) gin.HandlerFunc
	RoleCheckAny(roles ...enum.RoleCode) gin.HandlerFunc
	RoleCheckAll(roles ...enum.RoleCode) gin.HandlerFunc
}

type AuthImpl struct {
	service service.Service
	logger  *zap.Logger
}

func NewAuth(service service.Service, logger *zap.Logger) Auth {
	return &AuthImpl{service: service, logger: logger}
}

func (a *AuthImpl) RoleCheck(roleCheckType RoleCheckType, requiredRoles ...enum.RoleCode) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = ctx.FromGinCtx(c)
		uniCode, exc := ctx.GetUserUniCode()
		if exc != nil {
			ctx.ToError(exc)
			return
		}
		userRoles, exc := a.service.UserRole().GetCachedRolesCodeByUniCode(ctx, uniCode)
		if exc != nil {
			ctx.ToError(exc)
			return
		}
		if !a.checkRoles(roleCheckType, userRoles, requiredRoles) {
			ctx.ToError(exception.Forbidden)
			return
		}
		c.Next()
	}
}

// checkRoles 检查用户角色是否满足要求
func (a *AuthImpl) checkRoles(checkType RoleCheckType, userRoles, requiredRoles []enum.RoleCode) bool {
	if len(requiredRoles) == 0 {
		return true
	}

	userRoleSet := make(map[enum.RoleCode]struct{}, len(userRoles))
	for _, role := range userRoles {
		userRoleSet[role] = struct{}{}
	}

	switch checkType {
	case RoleCheckTypeAny:
		// 任意一个角色符合即可
		for _, required := range requiredRoles {
			if _, ok := userRoleSet[required]; ok {
				return true
			}
		}
		return false
	case RoleCheckTypeAll:
		// 所有角色都必须符合
		for _, required := range requiredRoles {
			if _, ok := userRoleSet[required]; !ok {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (a *AuthImpl) RoleCheckAny(roles ...enum.RoleCode) gin.HandlerFunc {
	return a.RoleCheck(RoleCheckTypeAny, roles...)
}

func (a *AuthImpl) RoleCheckAll(roles ...enum.RoleCode) gin.HandlerFunc {
	return a.RoleCheck(RoleCheckTypeAll, roles...)
}
