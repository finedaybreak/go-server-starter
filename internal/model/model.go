package model

import (
	"time"

	"gorm.io/gorm"
)

// 基础模型 有主键
type Model struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`     // 主键
	CreatedAt *time.Time     `json:"createdAt"`                // 创建时间
	UpdatedAt *time.Time     `json:"updatedAt"`                // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`           // 删除时间
	Version   uint64         `gorm:"default:0" json:"version"` // 版本号 用于乐观锁
}
