package service

import (
	"go-server-starter/internal/config"
	"go-server-starter/internal/repo"
	"go-server-starter/pkg/jwt"
	"go-server-starter/pkg/redis"
	"go-server-starter/pkg/snowflake"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	User() UserService
	UserRole() UserRoleService
	Auth() AuthService
}

type ServiceImpl struct {
	db              *gorm.DB
	config          *config.Config
	jwt             *jwt.JWT
	redis           *redis.Redis
	snowflake       *snowflake.Snowflake
	logger          *zap.Logger
	userService     UserService
	userRoleService UserRoleService
	authService     AuthService
}

func NewService(db *gorm.DB, config *config.Config, jwt *jwt.JWT, redis *redis.Redis, snowflake *snowflake.Snowflake, repo repo.Repo, logger *zap.Logger) Service {
	return &ServiceImpl{
		db:              db,
		config:          config,
		jwt:             jwt,
		redis:           redis,
		snowflake:       snowflake,
		logger:          logger,
		userService:     NewUserService(repo, redis, logger),
		userRoleService: NewUserRoleService(repo, redis, logger),
		authService:     NewAuthService(repo, jwt, logger),
	}
}

func (s *ServiceImpl) User() UserService {
	return s.userService
}

func (s *ServiceImpl) UserRole() UserRoleService {
	return s.userRoleService
}

func (s *ServiceImpl) Auth() AuthService {
	return s.authService
}
