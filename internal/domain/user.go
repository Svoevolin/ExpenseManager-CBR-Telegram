package domain

import (
	"database/sql"
	"time"
)

type User struct {
	UserID          int64 `gorm:"primarykey"`
	DefaultCurrency string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
}
