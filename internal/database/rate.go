package database

import (
	"context"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"gorm.io/gorm"
)

type RateDB struct {
	db *gorm.DB
}

func NewRateDB(db *gorm.DB) *RateDB {
	return &RateDB{db: db}
}

func (db *RateDB) AddRate(ctx context.Context, rate domain.Rate) error {
	db.db.WithContext(ctx).Where(domain.Rate{Code: rate.Code, Date: rate.Date}).Assign(rate).FirstOrCreate(&rate)
	return nil
}

func (db *RateDB) GetRate(ctx context.Context, code string, date time.Time) *domain.Rate {
	var resp domain.Rate
	result := db.db.WithContext(ctx).Where(domain.Rate{Code: code, Date: date}).Take(&resp)

	if result.RowsAffected == 0 {
		return nil
	}
	return &resp
}
