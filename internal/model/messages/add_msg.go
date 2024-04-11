package messages

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/helpers/money"
)

const FormatAddedExpenseMessage = "Расход добавлен: %v %s %s %s"

func (s *Model) addExpense(ctx context.Context, msg Message) (string, error) {
	date := time.Now().UTC()
	title := ""

	parts := strings.Split(strings.TrimPrefix(msg.Text, "/add"), ";")

	if len(parts) == 1 && strings.TrimSpace(parts[0]) == "" {
		return "", ErrInvalidCommand
	}

	if len(parts) >= 2 {
		title = strings.TrimSpace(parts[1])
	}

	kopecks, err := money.ConvertStringAmountToKopecks(parts[0])
	kopecksTemp := kopecks
	if err != nil {
		log.Printf("[%d]: %s", msg.UserID, err.Error())
		return "", ErrInvalidAmount
	}

	if len(parts) == 3 && strings.TrimSpace(parts[2]) != "" {
		date, err = time.ParseInLocation(dateFormat, strings.ReplaceAll(parts[2], " ", ""), time.UTC)
		if err != nil {
			log.Printf("[%d]: %s", msg.UserID, err.Error())
			return "", ErrInvalidDate
		}
	}

	userSelectedCurrency, err := s.userDB.GetDefaultCurrency(ctx, msg.UserID)
	if err != nil {
		return "", err
	}

	if userSelectedCurrency != s.config.GetBaseCurrency() {
		rate := s.rateDB.GetRate(ctx, userSelectedCurrency, date)

		if rate == nil {
			if err := s.rateUpdater.UpdateCurrency(ctx, date); err != nil {
				return "", err
			}

			rate = s.rateDB.GetRate(ctx, userSelectedCurrency, date)
		}

		kopecks = int64(float64(kopecks) * float64(rate.Kopecks) / float64(100) / float64(rate.Nominal))
	}

	if err := s.expenseDB.AddExpense(ctx, msg.UserID, kopecks, title, date); err != nil {
		log.Printf("[%d]: %s", msg.UserID, err.Error())
		return "", ErrWriteToDatabase
	}

	return fmt.Sprintf(FormatAddedExpenseMessage, money.ConvertKopecksToAmount(kopecksTemp), userSelectedCurrency, title, date.Format(dateFormat)), nil
}
