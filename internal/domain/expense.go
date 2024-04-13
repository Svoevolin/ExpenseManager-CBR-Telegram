package domain

import (
	"database/sql"
	"time"
)

type Expense struct {
	ID        uint `gorm:"primarykey"`
	UserID    int64
	Title     string
	Amount    int64
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
