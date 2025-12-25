package seed

import (
	"context"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/model"
	"go-server-starter/internal/repo"

	"go.uber.org/zap"
)

type Seed interface {
	Run() error
}

type seed struct {
	repo   repo.Repo
	logger *zap.Logger
}

func NewSeed(repo repo.Repo, logger *zap.Logger) Seed {
	return &seed{repo: repo, logger: logger}
}

func (s *seed) Run() error {
	if err := s.SeedUserRole(); err != nil {
		return err
	}
	return nil
}

func (s *seed) SeedUserRole() error {
	var ctx = context.Background()
	var roles = []enum.RoleCode{
		enum.RoleCodeSuperAdmin,
		enum.RoleCodeAdmin,
		enum.RoleCodeGuest,
		enum.RoleCodeUser,
		enum.RoleCodeUserVip,
		enum.RoleCodeUserSvip,
	}

	// 查询数据库中已存在的角色
	existingRoles, err := s.repo.UserRole().GetMany(ctx)
	if err != nil {
		s.logger.Error("failed to query existing roles", zap.Error(err))
		return err
	}

	// 创建已存在角色代码的 map，用于快速查找
	existingRoleMap := make(map[enum.RoleCode]bool)
	for _, role := range existingRoles {
		existingRoleMap[role.Code] = true
	}

	// 过滤出需要插入的角色
	var newRoles []*model.UserRole

	for _, roleCode := range roles {
		if !existingRoleMap[roleCode] {
			newRoles = append(newRoles, &model.UserRole{
				Code:    roleCode,
				Enabled: true,
			})
		}
	}

	// 如果有需要插入的角色，批量插入
	if len(newRoles) > 0 {
		if err := s.repo.UserRole().CreateBatch(ctx, newRoles); err != nil {
			s.logger.Error("failed to insert roles", zap.Error(err))
			return err
		}
		s.logger.Info("successfully seeded roles", zap.Int("count", len(newRoles)))
	} else {
		s.logger.Info("all roles already exist, no need to seed")
	}

	return nil
}
