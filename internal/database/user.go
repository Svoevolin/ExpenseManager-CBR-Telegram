package database

import (
	"context"
	"fmt"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (db *UserDB) UserExists(ctx context.Context, userID int64) bool {
	var user domain.User
	result := db.db.WithContext(ctx).Where(domain.User{UserID: userID}).Find(&user)

	return result.RowsAffected != 0
}

func (db *UserDB) GetDefaultCurrency(ctx context.Context, userID int64) (string, error) {
	var user domain.User
	result := db.db.WithContext(ctx).Where(domain.User{UserID: userID}).Find(&user)

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("user №%d not found", userID)
	}
	return user.DefaultCurrency, nil
}

func (db *UserDB) ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error {
	db.db.WithContext(ctx).Where(domain.User{UserID: userID}).Assign(domain.User{
		DefaultCurrency: currency,
	}).FirstOrCreate(&domain.User{})

	return nil
}
