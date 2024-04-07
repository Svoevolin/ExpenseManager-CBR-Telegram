package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
)

type UserDB struct {
	sync.RWMutex
	store map[int64]domain.User
}

func NewUserDB() (*UserDB, error) {
	return &UserDB{
		store: make(map[int64]domain.User),
	}, nil
}

func (db *UserDB) UserExists(ctx context.Context, userID int64) bool {
	db.RLock()
	defer db.RUnlock()

	_, ok := db.store[userID]
	return ok
}

func (db *UserDB) GetDefaultCurrency(ctx context.Context, userID int64) (string, error) {
	db.RLock()
	defer db.RUnlock()

	if user, ok := db.store[userID]; ok {
		return user.DefaultCurrency, nil
	}

	return "", fmt.Errorf("user №%d not found", userID)
}

func (db *UserDB) ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error {
	db.Lock()
	defer db.Unlock()

	db.store[userID] = domain.User{UserID: userID, DefaultCurrency: currency}

	return nil
}
