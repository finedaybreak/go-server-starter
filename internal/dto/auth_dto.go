package dto

// 手机和手机验证码登录(没有注册自动注册)
type AuthLoginByMobileAndCodeReqDto struct {
	Mobile      string `json:"mobile" binding:"required"`
	Code        string `json:"code" binding:"required"`
	CountryCode string `json:"countryCode" binding:"required"`
}

// 邮箱和邮箱验证码登录 (没有注册自动注册)
type AuthLoginByEmailAndCodeReqDto struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type AuthTokenResDto struct {
	Token string `json:"token"`
}
