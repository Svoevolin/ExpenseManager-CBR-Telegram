package domain

import (
	"time"
)

type Expense struct {
	ID     uint `gorm:"primarykey"`
	UserID int64
	Title  string
	Amount int64
	Date   time.Time
}
