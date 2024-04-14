package domain

import (
	"time"
)

type Rate struct {
	Code      string `gorm:"primaryKey"`
	Nominal   int64
	Kopecks   int64
	Original  string
	Date      time.Time `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
