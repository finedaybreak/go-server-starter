package repo

import (
	"go-server-starter/internal/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRoleRepo interface {
	BaseRepo[model.UserRole]
	WithTx(tx *gorm.DB) UserRoleRepo
}

type UserRoleRepoImpl struct {
	BaseRepo[model.UserRole]
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRoleRepo(db *gorm.DB, logger *zap.Logger) UserRoleRepo {
	return &UserRoleRepoImpl{
		BaseRepo: NewBaseRepo[model.UserRole](db, logger),
		db:       db,
		logger:   logger,
	}
}

func (r *UserRoleRepoImpl) WithTx(tx *gorm.DB) UserRoleRepo {
	return &UserRoleRepoImpl{
		BaseRepo: NewBaseRepo[model.UserRole](tx, r.logger),
		db:       tx,
		logger:   r.logger,
	}
}
