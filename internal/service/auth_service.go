package service

import (
	"errors"
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/dto"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/exception"
	"go-server-starter/internal/model"
	"go-server-starter/internal/repo"
	"go-server-starter/pkg/jwt"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginByMobileAndCode(ctx *ctx.Context, params dto.AuthLoginByMobileAndCodeReqDto) (*dto.AuthTokenResDto, *exception.Exception)
	LoginByEmailAndCode(ctx *ctx.Context, params dto.AuthLoginByEmailAndCodeReqDto) (*dto.AuthTokenResDto, *exception.Exception)
}

type AuthServiceImpl struct {
	repo   repo.Repo
	jwt    *jwt.JWT
	logger *zap.Logger
}

func NewAuthService(repo repo.Repo, jwt *jwt.JWT, logger *zap.Logger) AuthService {
	return &AuthServiceImpl{
		repo:   repo,
		jwt:    jwt,
		logger: logger,
	}
}

func (s *AuthServiceImpl) LoginByMobileAndCode(ctx *ctx.Context, params dto.AuthLoginByMobileAndCodeReqDto) (*dto.AuthTokenResDto, *exception.Exception) {
	if params.Code != "666666" {
		return nil, exception.UserMobileVerificationCodeIsIncorrect
	}
	var deviceType = ctx.GetDeviceType()
	params.Mobile = strings.ReplaceAll(params.Mobile, " ", "")
	user, err := s.repo.User().GetOne(ctx.Ctx, repo.Where("mobile = ? AND country_code = ?", params.Mobile, params.CountryCode))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	if user == nil {
		// 创建用户
		tx := s.repo.DB().Begin()
		userRepo := s.repo.User().WithTx(tx)
		userRoleRepo := s.repo.UserRole().WithTx(tx)
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		uniCode, err := userRepo.GenerateUniCode(ctx.Ctx)
		if err != nil {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		// 绑定角色
		role, err := userRoleRepo.GetOne(ctx.Ctx, repo.Where("code = ?", enum.RoleCodeUser))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		if role == nil {
			tx.Rollback()
			return nil, exception.UserRoleNotFound
		}
		user = &model.User{
			UniCode:     uniCode,
			Mobile:      params.Mobile,
			CountryCode: params.CountryCode,
			Nickname:    params.Mobile,
			Roles:       []model.UserRole{*role},
		}
		err = userRepo.Create(ctx.Ctx, user)
		if err != nil {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		if err := tx.Commit().Error; err != nil {
			return nil, exception.InternalServerError.Append(err.Error())
		}
	}
	token, err := s.jwt.GenerateToken(user.UniCode, deviceType)
	if err != nil {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	return &dto.AuthTokenResDto{
		Token: token,
	}, nil
}

func (s *AuthServiceImpl) LoginByEmailAndCode(ctx *ctx.Context, params dto.AuthLoginByEmailAndCodeReqDto) (*dto.AuthTokenResDto, *exception.Exception) {
	if params.Code != "666666" {
		return nil, exception.UserEmailVerificationCodeIsIncorrect
	}
	var deviceType = ctx.GetDeviceType()
	params.Email = strings.ToLower(strings.TrimSpace(params.Email))
	user, err := s.repo.User().GetOne(ctx.Ctx, repo.Where("email = ?", params.Email))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	if user == nil {
		// 创建用户
		tx := s.repo.DB().Begin()
		userRepo := s.repo.User().WithTx(tx)
		userRoleRepo := s.repo.UserRole().WithTx(tx)
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		uniCode, err := userRepo.GenerateUniCode(ctx.Ctx)
		if err != nil {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		// 绑定角色
		role, err := userRoleRepo.GetOne(ctx.Ctx, repo.Where("code = ?", enum.RoleCodeUser))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		if role == nil {
			tx.Rollback()
			return nil, exception.UserRoleNotFound
		}
		user = &model.User{
			UniCode:  uniCode,
			Email:    params.Email,
			Nickname: params.Email,
			Roles:    []model.UserRole{*role},
		}
		err = userRepo.Create(ctx.Ctx, user)
		if err != nil {
			tx.Rollback()
			return nil, exception.InternalServerError.Append(err.Error())
		}
		if err := tx.Commit().Error; err != nil {
			return nil, exception.InternalServerError.Append(err.Error())
		}
	}
	token, err := s.jwt.GenerateToken(user.UniCode, deviceType)
	if err != nil {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	return &dto.AuthTokenResDto{
		Token: token,
	}, nil
}
