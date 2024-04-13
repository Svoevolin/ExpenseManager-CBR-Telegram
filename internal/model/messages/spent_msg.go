package messages

import (
	"context"
	"fmt"
	"strings"
	"time"

	utils "github.com/Svoevolin/workshop_1_bot/internal/helpers/date"
	"github.com/Svoevolin/workshop_1_bot/internal/helpers/money"
)

const (
	formatTotalListMessage = "Расходов за %s: %v %s\n------------------------------"
	formatItemListMessage  = "\n- %v %s %s (%s)"
)

const FailedGetListExpensesMessage = "Не удалось получить расходы, повторите команду позже."

func (s *Model) listExpenses(ctx context.Context, msg Message) (string, error) {
	today := time.Now()

	var dur time.Duration

	interval := "всё время"
	switch strings.TrimPrefix(msg.Text, "/spent_") {
	case "day":
		dur = time.Hour * 24
		interval = "день"
	case "week":
		dur = time.Hour * 24 * 7
		interval = "неделя"
	case "month":
		dur = time.Hour * 24 * 31
		interval = "месяц"
	case "year":
		dur = time.Hour * 24 * 365
		interval = "год"
	}

	userSelectedCurrency, _ := s.userDB.GetDefaultCurrency(ctx, msg.UserID)
	baseCurrency := s.config.GetBaseCurrency()
	expenses, err := s.expenseDB.GetExpenses(ctx, msg.UserID)
	if err != nil {
		return "", ErrGetRecordsInDatabase
	}

	list := ""
	var total int64
	for _, v := range expenses {
		if today.Sub(v.Date) < dur || interval == "всё время" {
			if userSelectedCurrency != baseCurrency {
				rate := s.rateDB.GetRate(ctx, userSelectedCurrency, utils.GetDate(v.Date))

				if rate == nil {
					if err := s.rateUpdater.UpdateCurrency(ctx, v.Date); err != nil {
						return "", err
					}
					rate = s.rateDB.GetRate(ctx, userSelectedCurrency, utils.GetDate(v.Date))
				}

				v.Amount = int64(float64(v.Amount*rate.Nominal) / float64(rate.Kopecks) * 100)
			}

			total += v.Amount
			list += fmt.Sprintf(formatItemListMessage, money.ConvertKopecksToAmount(v.Amount), userSelectedCurrency, v.Title, v.Date.Format(dateFormat))
		}
	}
	return fmt.Sprintf(formatTotalListMessage+list, interval, money.ConvertKopecksToAmount(total), userSelectedCurrency), nil
}
