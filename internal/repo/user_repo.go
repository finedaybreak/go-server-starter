package repo

import (
	"context"
	"errors"
	"fmt"
	"go-server-starter/internal/model"
	"go-server-starter/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	BaseRepo[model.User]
	WithTx(tx *gorm.DB) UserRepo
	GenerateUniCode(ctx context.Context) (string, error)
	GetIDByUniCode(ctx context.Context, uniCode string) (uint64, error)
	GetByUniCode(ctx context.Context, uniCode string) (*model.User, error)
	GetRolesByUniCode(ctx context.Context, uniCode string) ([]*model.UserRole, error)
	GetRolesByID(ctx context.Context, id uint64) ([]*model.UserRole, error)
}

type UserRepoImpl struct {
	BaseRepo[model.User]
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepo(db *gorm.DB, logger *zap.Logger) UserRepo {
	return &UserRepoImpl{
		BaseRepo: NewBaseRepo[model.User](db, logger),
		db:       db,
		logger:   logger,
	}
}

func (r *UserRepoImpl) WithTx(tx *gorm.DB) UserRepo {
	return &UserRepoImpl{
		BaseRepo: NewBaseRepo[model.User](tx, r.logger),
		db:       tx,
		logger:   r.logger,
	}
}

func (r *UserRepoImpl) GenerateUniCode(ctx context.Context) (string, error) {
	for {
		uniCode := utils.RandomUserCode()
		exists, err := r.GetOne(ctx, Where("uni_code = ?", uniCode))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("query user by uni code failed: %w", err)
		}
		if exists == nil {
			return uniCode, nil
		}
	}
}

func (r *UserRepoImpl) GetIDByUniCode(ctx context.Context, uniCode string) (uint64, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("uni_code = ?", uniCode).Preload("Roles").First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepoImpl) GetByUniCode(ctx context.Context, uniCode string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("uni_code = ?", uniCode).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) GetRolesByUniCode(ctx context.Context, uniCode string) ([]*model.UserRole, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("uni_code = ?", uniCode).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	roles := make([]*model.UserRole, 0)
	for i := range user.Roles {
		var role = user.Roles[i]
		if role.Enabled {
			roles = append(roles, &role)
		}
	}
	return roles, nil
}

func (r *UserRepoImpl) GetRolesByID(ctx context.Context, id uint64) ([]*model.UserRole, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	roles := make([]*model.UserRole, 0)
	for i := range user.Roles {
		var role = user.Roles[i]
		if role.Enabled {
			roles = append(roles, &role)
		}
	}
	return roles, nil
}
