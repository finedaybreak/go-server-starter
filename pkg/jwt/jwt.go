package jwt

import (
	"errors"
	"go-server-starter/internal/config"
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/enum"
	"go-server-starter/internal/exception"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type CustomClaims struct {
	UniCode string `json:"uniCode"` // 用户唯一码
	jwt.RegisteredClaims
}

type JWT struct {
	config *config.JWTConfig
	logger *zap.Logger
}

func NewJWT(config *config.JWTConfig, logger *zap.Logger) *JWT {
	j := &JWT{
		config: config,
		logger: logger,
	}
	return j
}

func (j *JWT) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = ctx.FromGinCtx(c)
		token, err := j.GetTokenFromGinContext(c)
		if err != nil {
			ctx.ToError(exception.TokenNotFound.Append(err.Error()))
			return
		}
		claims, err := j.ParseAndVerifyToken(token)
		if err != nil {
			ctx.ToError(exception.TokenInvalid.Append(err.Error()))
			return
		}
		// 设置用户唯一码
		ctx.SetUserUniCode(claims.UniCode)
		// 判断是否需要刷新令牌
		if j.isTokenNeedRefresh(ctx, claims.ExpiresAt.Time) {
			// 刷新令牌
			newToken, err := j.GenerateToken(claims.UniCode, ctx.GetDeviceType())
			if err != nil {
				ctx.ToError(exception.TokenGenerateFailed.Append(err.Error()))
				return
			}
			c.Header("new-token", newToken)
		}
		c.Next()
	}
}

func (j *JWT) isTokenNeedRefresh(ctx *ctx.Context, expiresAt time.Time) bool {
	// 判断是否需要刷新令牌 还有多少时间过期
	var remaining = time.Until(expiresAt)
	expire := j.config.TokenExpires.Get(ctx.GetDeviceType())
	// 如果剩余时间小于过期时间的1/3，则需要刷新令牌
	return (remaining < expire/3) && (remaining > 0)
}

func (j *JWT) GetTokenFromGinContext(c *gin.Context) (string, error) {
	var token string
	keys := []string{"Authorization", "token"}
	for _, key := range keys {
		if value, exist := c.GetQuery(key); exist {
			token = strings.TrimPrefix(value, "Bearer ")
			break
		}
		if value := c.GetHeader(key); value != "" {
			token = strings.TrimPrefix(value, "Bearer ")
			break
		}
	}
	if token == "" {
		return "", errors.New("token is empty")
	}
	return token, nil
}

func (j *JWT) GenerateToken(uniCode string, deviceType enum.DeviceType) (string, error) {
	expire := j.config.TokenExpires.Get(deviceType)
	now := time.Now()
	claims := CustomClaims{
		UniCode: uniCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expire)),
			Issuer:    j.config.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (j *JWT) ParseAndVerifyToken(tokenStr string) (*CustomClaims, error) {
	var token, err = jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.config.TokenSecret), nil
	})
	if err != nil || token == nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid token claims type")
	}
	return claims, nil
}
