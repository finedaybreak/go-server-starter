package service

import (
	"errors"
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/dto"
	"go-server-starter/internal/exception"
	"go-server-starter/internal/model"
	"go-server-starter/internal/repo"
	"go-server-starter/pkg/redis"
	"go-server-starter/pkg/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	GetByID(ctx *ctx.Context, id uint64) (*model.User, *exception.Exception)
	GetByUniCode(ctx *ctx.Context, uniCode string) (*model.User, *exception.Exception)
	GetInfoByUniCode(ctx *ctx.Context, uniCode string) (*dto.UserInfoResDto, *exception.Exception)
	GetTable(ctx *ctx.Context, params dto.UserTableQueryReqDto) (*dto.PaginationResDto[[]*dto.UserListItemResDto], *exception.Exception)
	UpdateInfo(ctx *ctx.Context, params dto.UserUpdateInfoReqDto) (*dto.UserInfoResDto, *exception.Exception)
}

type UserServiceImpl struct {
	repo   repo.Repo
	logger *zap.Logger
}

func NewUserService(repo repo.Repo, redis *redis.Redis, logger *zap.Logger) UserService {
	return &UserServiceImpl{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserServiceImpl) GetByID(ctx *ctx.Context, id uint64) (*model.User, *exception.Exception) {
	user, err := s.repo.User().GetByID(ctx.Ctx, id, repo.Preload("Roles"))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	if user == nil {
		return nil, exception.UserNotFound
	}
	return user, nil
}

func (s *UserServiceImpl) GetByUniCode(ctx *ctx.Context, uniCode string) (*model.User, *exception.Exception) {
	user, err := s.repo.User().GetByUniCode(ctx.Ctx, uniCode)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	if user == nil {
		return nil, exception.UserNotFound
	}
	return user, nil
}

func (s *UserServiceImpl) GetInfoByUniCode(ctx *ctx.Context, uniCode string) (*dto.UserInfoResDto, *exception.Exception) {
	user, err := s.GetByUniCode(ctx, uniCode)
	if err != nil {
		return nil, err
	}

	// 转换角色为字符串数组
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Code.String()
	}

	return &dto.UserInfoResDto{
		UniCode:     user.UniCode,
		Email:       user.Email,
		Mobile:      user.Mobile,
		CountryCode: user.CountryCode,
		Desc:        user.Desc,
		Nickname:    user.Nickname,
		AvatarURL:   user.AvatarURL,
		Roles:       roles,
	}, nil
}

func (s *UserServiceImpl) GetTable(ctx *ctx.Context, params dto.UserTableQueryReqDto) (*dto.PaginationResDto[[]*dto.UserListItemResDto], *exception.Exception) {
	opts := []repo.QueryOption{
		repo.Order("created_at DESC"),
		repo.Preload("Roles"),
		repo.WhereAutoLike("nickname", params.Nickname),
		repo.WhereAutoLikePrefix("email", params.Email),
		repo.WhereAutoLikePrefix("mobile", params.Mobile),
		repo.WhereAutoLikePrefix("country_code", params.CountryCode),
	}
	users, total, err := s.repo.User().GetTable(ctx.Ctx, params.Page, params.PageSize, opts...)
	if err != nil {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	res := make([]*dto.UserListItemResDto, len(users))
	for i, user := range users {
		res[i] = &dto.UserListItemResDto{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt.Format(time.RFC3339),
			UniCode:     user.UniCode,
			Email:       user.Email,
			Mobile:      user.Mobile,
			CountryCode: user.CountryCode,
			Desc:        user.Desc,
			Nickname:    user.Nickname,
			AvatarURL:   user.AvatarURL,
		}
		roles := make([]string, len(user.Roles))
		for i, role := range user.Roles {
			roles[i] = role.Code.String()
		}
		res[i].Roles = roles
	}
	return utils.AssemblePaginationResDto(res, total, params.Page, params.PageSize), nil
}

func (s *UserServiceImpl) UpdateInfo(ctx *ctx.Context, params dto.UserUpdateInfoReqDto) (*dto.UserInfoResDto, *exception.Exception) {
	userID, exc := ctx.GetUserID(s.repo)
	if exc != nil {
		return nil, exc
	}
	user, err := s.repo.User().GetByID(ctx.Ctx, userID, repo.Preload("Roles"))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.InternalServerError.Append(err.Error())
	}
	if user == nil {
		return nil, exception.UserNotFound
	}
	if params.Nickname != nil {
		user.Nickname = *params.Nickname
	}
	if params.AvatarURL != nil {
		user.AvatarURL = *params.AvatarURL
	}
	if params.Desc != nil {
		user.Desc = *params.Desc
	}
	if err := s.repo.User().UpdateByZeroFields(ctx.Ctx, userID, user); err != nil {
		return nil, exception.UserUpdateInfoFailed.Append(err.Error())
	}
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Code.String()
	}
	return &dto.UserInfoResDto{
		UniCode:     user.UniCode,
		Email:       user.Email,
		Mobile:      user.Mobile,
		CountryCode: user.CountryCode,
		Desc:        user.Desc,
		Nickname:    user.Nickname,
		AvatarURL:   user.AvatarURL,
		Roles:       roles,
	}, nil
}
