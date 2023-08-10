package model

import (
	"database/sql"
	"time"
)

// Basic 同 gorm.Model  为了自定义序列化字段
type Basic struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
