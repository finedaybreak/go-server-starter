package repo

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repo interface {
	DB() *gorm.DB
	Logger() *zap.Logger
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	User() UserRepo
	UserRole() UserRoleRepo
}

type RepoImpl struct {
	db           *gorm.DB
	logger       *zap.Logger
	userRepo     UserRepo
	userRoleRepo UserRoleRepo
}

func NewRepo(db *gorm.DB, logger *zap.Logger) Repo {
	return &RepoImpl{
		db:           db,
		logger:       logger,
		userRepo:     NewUserRepo(db, logger),
		userRoleRepo: NewUserRoleRepo(db, logger),
	}
}

func (r *RepoImpl) DB() *gorm.DB {
	return r.db
}

func (r *RepoImpl) Logger() *zap.Logger {
	return r.logger
}

func (r *RepoImpl) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}

func (r *RepoImpl) User() UserRepo {
	return r.userRepo
}

func (r *RepoImpl) UserRole() UserRoleRepo {
	return r.userRoleRepo
}
