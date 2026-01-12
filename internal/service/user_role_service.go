package service

import (
	"encoding/json"
	"go-server-starter/internal/constant"
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/exception"
	"go-server-starter/internal/repo"
	"go-server-starter/pkg/redis"

	goredis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserRoleService interface {
	GetRolesCodeByUniCode(ctx *ctx.Context, uniCode string) ([]enum.RoleCode, *exception.Exception)
	GetCachedRolesCodeByUniCode(ctx *ctx.Context, uniCode string) ([]enum.RoleCode, *exception.Exception)
}

type UserRoleServiceImpl struct {
	repo   repo.Repo
	redis  *redis.Redis
	logger *zap.Logger
}

func NewUserRoleService(repo repo.Repo, redis *redis.Redis, logger *zap.Logger) UserRoleService {
	return &UserRoleServiceImpl{
		repo:   repo,
		redis:  redis,
		logger: logger,
	}
}

func (s *UserRoleServiceImpl) GetRolesCodeByUniCode(ctx *ctx.Context, uniCode string) ([]enum.RoleCode, *exception.Exception) {
	roles, err := s.repo.User().GetRolesByUniCode(ctx.Ctx, uniCode)
	if err != nil {
		s.logger.Error("get roles code by uni code failed", zap.String("uniCode", uniCode), zap.Error(err))
		return nil, exception.InternalServerError.Append(err.Error())
	}
	rolesCode := make([]enum.RoleCode, len(roles))
	for i, role := range roles {
		rolesCode[i] = role.Code
	}
	return rolesCode, nil
}

func (s *UserRoleServiceImpl) GetCachedRolesCodeByUniCode(ctx *ctx.Context, uniCode string) ([]enum.RoleCode, *exception.Exception) {
	dataStr, err := s.redis.Get(ctx.Ctx, constant.RedisKeyOfAuthRoles(uniCode)).Result()
	if err != nil {
		// 如果redis 非正常报错，则返回错误
		if err != goredis.Nil {
			s.logger.Error("get cached roles code by uni code failed", zap.String("uniCode", uniCode), zap.Error(err))
			return nil, exception.InternalServerError.Append(err.Error())
		} else {
			// 如果redis 正常报错（goredis.Nil），则获取数据库中的角色
			roles, exc := s.GetRolesCodeByUniCode(ctx, uniCode)
			if exc != nil {
				return nil, exc
			}
			// 将角色转换为JSON
			rolesJSON, err := json.Marshal(roles)
			if err != nil {
				s.logger.Error("marshal roles code by uni code failed", zap.String("uniCode", uniCode), zap.Error(err))
				return nil, exception.InternalServerError.Append(err.Error())
			}
			// 将角色缓存到redis
			if err := s.redis.Set(ctx.Ctx, constant.RedisKeyOfAuthRoles(uniCode), rolesJSON, constant.REDIS_EXPIRE_OF_AUTH_ROLES).Err(); err != nil {
				s.logger.Error("set cached roles code by uni code failed", zap.String("uniCode", uniCode), zap.Error(err))
				return nil, exception.InternalServerError.Append(err.Error())
			}
			return roles, nil
		}
	} else {
		// 如果redis 正常返回，则将dataStr 反序列化为角色
		var roles []enum.RoleCode
		if err := json.Unmarshal([]byte(dataStr), &roles); err != nil {
			s.logger.Error("unmarshal roles code by uni code failed", zap.String("uniCode", uniCode), zap.Error(err))
			return nil, exception.InternalServerError.Append(err.Error())
		}
		return roles, nil
	}
}
