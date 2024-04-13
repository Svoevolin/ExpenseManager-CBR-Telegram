package database

import (
	"context"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"gorm.io/gorm"
)

type ExpenseDB struct {
	db *gorm.DB
}

func NewExpenseDB(db *gorm.DB) *ExpenseDB {
	return &ExpenseDB{db: db}
}

func (db *ExpenseDB) AddExpense(ctx context.Context, userID int64, kopecks int64, title string, date time.Time) error {
	result := db.db.WithContext(ctx).Create(&domain.Expense{
		UserID: userID,
		Title:  title,
		Date:   date,
		Amount: kopecks,
	})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (db *ExpenseDB) GetExpenses(ctx context.Context, userID int64) ([]domain.Expense, error) {
	var expenses []domain.Expense
	result := db.db.WithContext(ctx).Where(domain.Expense{UserID: userID}).Find(&expenses)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return expenses, nil
}
