package database

import (
	"context"
	"sync"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	utils "github.com/Svoevolin/workshop_1_bot/internal/helpers/date"
)

type RateDB struct {
	store map[string]map[time.Time]domain.Rate
	sync.RWMutex
}

func NewRateDB() (*RateDB, error) {
	return &RateDB{
		store: make(map[string]map[time.Time]domain.Rate),
	}, nil
}

func (db *RateDB) AddRate(ctx context.Context, date time.Time, rate domain.Rate) error {
	db.Lock()
	defer db.Unlock()

	if _, ok := db.store[rate.Code][utils.GetDate(date)]; ok {
		return nil
	}

	db.store[rate.Code] = map[time.Time]domain.Rate{utils.GetDate(date): rate}

	return nil
}

func (db *RateDB) GetRate(ctx context.Context, code string, date time.Time) *domain.Rate {
	db.RLock()
	defer db.RUnlock()

	if rate, ok := db.store[code][utils.GetDate(date)]; ok {
		return &rate
	}

	return nil
}
