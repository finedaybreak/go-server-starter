package model

import "go-server-starter/internal/enum"

type UserRole struct {
	Model
	Code    enum.RoleCode `gorm:"unique;not null" json:"code"`
	Enabled bool          `gorm:"default:true" json:"enabled"`
}
