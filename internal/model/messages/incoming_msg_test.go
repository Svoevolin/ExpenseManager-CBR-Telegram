package messages

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/messages"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func parseDate(date string) time.Time {
	v, _ := time.ParseInLocation(dateFormat, date, time.UTC)
	return v
}

func Test_OnAddCommand_ShouldAnswerWithAddedMessage(t *testing.T) {
	m := minimock.NewController(t)
	userID := int64(123)

	testCases := []struct {
		name    string
		command string
		amount  string
		kopecks int64
		title   string
		date    time.Time
		answer  string
	}{
		{
			name:    "normal",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("11.04.2024"),
			command: "/add 100.0; расход; 11.04.2024",
			answer:  "Расход добавлен: 100.00 RUB расход 11.04.2024",
		},
		{
			name:    "with split for ';' give a len 4",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("11.04.2024"),
			command: "/add 100.0; расход; 11.04.2024;",
			answer:  "Расход добавлен: 100.00 RUB расход 11.04.2024",
		},
		{
			name:    "without title",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    parseDate("11.04.2024"),
			command: "/add 100.0; ; 11.04.2024",
			answer:  "Расход добавлен: 100.00 RUB  11.04.2024",
		},
		//{
		//	name:    "without date",
		//	amount:  "100.0",
		//	kopecks: 10000,
		//	title:   "расход",
		//	date:    parseDate("01.01.2011"),
		//	command: "/add 100.0; расход;",
		//	answer:  "Расход добавлен: 100.00 RUB расход " + time.Now().UTC().Format(dateFormat),
		//},
		//{
		//	name:    "only amount with semicolon",
		//	amount:  "100.0",
		//	kopecks: 10000,
		//	title:   "",
		//	date:    parseDate("01.01.2011"),
		//	command: "/add 100.0;",
		//	answer:  "Расход добавлен: 100.00 RUB " + time.Now().UTC().Format(dateFormat),
		//},
		//{
		//	name:    "only amount without semicolon",
		//	amount:  "100.0",
		//	kopecks: 10000,
		//	title:   "",
		//	date:    parseDate("01.01.2011"),
		//	command: "/add 100.0",
		//	answer:  "Расход добавлен: 100.00 RUB  " + time.Now().UTC().Format(dateFormat),
		//},
		{
			name:    "without amount",
			amount:  "",
			kopecks: 0,
			title:   "расход",
			date:    parseDate("11.04.2024"),
			command: "/add ; расход; 11.04.2024",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid amount",
			amount:  "100.0.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("11.04.2024"),
			command: "/add 100.0.0; расход; 11.04.2024",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid date",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("11.04.2024"),
			command: "/add 100.0; расход; 11.54.2024",
			answer:  InvalidDateMessage,
		},
		{
			name:    "invalid date format",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("01.01.2011"),
			command: "/add 100.0; расход; 11.04.2024.0",
			answer:  InvalidDateMessage,
		},
		{
			name:    "empty command",
			date:    parseDate("01.01.2011"),
			command: "/add",
			answer:  InvalidCommandMessage,
		},
	}
	baseCurrency := "RUB"
	sender := mock.NewMessageSenderMock(m)
	settings := mock.NewConfigGetterMock(m)
	userDB := mock.NewUserDBMock(m)
	rateDB := mock.NewRateDBMock(m)
	expenseDB := mock.NewExpenseDBMock(m)
	updater := mock.NewExchangeRateUpdateMock(m)
	model := New(sender, settings, userDB, rateDB, expenseDB, updater)

	settings.GetBaseCurrencyMock.Expect().Return(baseCurrency)
	userDB.GetDefaultCurrencyMock.Expect(minimock.AnyContext, userID).Return(baseCurrency, nil)
	userDB.UserExistsMock.Expect(minimock.AnyContext, userID).Return(true)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expenseDB.AddExpenseMock.Expect(minimock.AnyContext, userID, tc.kopecks, tc.title, tc.date).Return(nil)
			sender.SendMessageMock.Expect(tc.answer, userID).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   tc.command,
				UserID: userID,
			})

			assert.NoError(t, err)
		})
	}
}

func Test_OnAddCommand_WithConverting(t *testing.T) {
	m := minimock.NewController(t)
	sender := mock.NewMessageSenderMock(m)
	settings := mock.NewConfigGetterMock(m)
	userDB := mock.NewUserDBMock(m)
	rateDB := mock.NewRateDBMock(m)
	expenseDB := mock.NewExpenseDBMock(m)
	updater := mock.NewExchangeRateUpdateMock(m)
	model := New(sender, settings, userDB, rateDB, expenseDB, updater)

	userID := int64(123)
	settings.GetBaseCurrencyMock.Expect().Return("RUB")
	userDB.UserExistsMock.Expect(minimock.AnyContext, userID).Return(true)
	userDB.GetDefaultCurrencyMock.Expect(minimock.AnyContext, userID).Return("USD", nil)

	tcs := []struct {
		usd float64
		rub float64
	}{
		{usd: 1, rub: 93.21},
		{usd: 10, rub: 932.1},
		{usd: 15, rub: 1398.15},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(fmt.Sprintf("%f -> %f", tc.usd, tc.rub), func(t *testing.T) {
			t.Parallel()

			rateDB.GetRateMock.Expect(minimock.AnyContext, "USD", parseDate("11.10.2024")).Return(
				&domain.Rate{
					Nominal: 1,
					Kopecks: 9321,
				})

			expenseDB.AddExpenseMock.Expect(minimock.AnyContext, userID, int64(tc.rub*100), "купил снюс", parseDate("11.10.2024")).Return(nil)
			sender.SendMessageMock.Expect(fmt.Sprintf("Расход добавлен: %.2f USD купил снюс 11.10.2024", tc.usd), userID).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   fmt.Sprintf("/add %.2f; купил снюс; 11.10.2024", tc.usd),
				UserID: userID,
			})

			assert.NoError(t, err)
		})
	}
}

func TestStartCommand(t *testing.T) {

	t.Run("if user starts using for first time then we have to show him message about choosing currency", func(t *testing.T) {

		m := minimock.NewController(t)
		userDB := mock.NewUserDBMock(m)
		settings := mock.NewConfigGetterMock(m)
		sender := mock.NewMessageSenderMock(m)
		model := New(sender, settings, userDB, nil, nil, nil)

		exceptedCurrencies := []map[string]string{
			{
				"RUB": "/set_currency RUB",
				"USD": "/set_currency USD",
				"EUR": "/set_currency EUR",
			},
		}

		userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(false)
		settings.SupportedCurrencyCodesMock.Expect().Return([]string{"RUB", "USD", "EUR"})
		sender.SendMessageMock.Expect(newUserMessage, int64(123), exceptedCurrencies...).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{UserID: 123, Text: "/start"})
		assert.NoError(t, err)
	})

	t.Run("if user already exist we show help message", func(t *testing.T) {

		m := minimock.NewController(t)
		userDB := mock.NewUserDBMock(m)
		settings := mock.NewConfigGetterMock(m)
		sender := mock.NewMessageSenderMock(m)
		model := New(sender, settings, userDB, nil, nil, nil)

		userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(true)
		sender.SendMessageMock.Expect(startMessage, int64(123)).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{UserID: 123, Text: "/start"})
		assert.NoError(t, err)
	})
}

func Test_OnUnknownCommand_ShouldAnswerWithUnknownMessage(t *testing.T) {
	m := minimock.NewController(t)
	userDB := mock.NewUserDBMock(m)
	sender := mock.NewMessageSenderMock(m)
	model := New(sender, nil, userDB, nil, nil, nil)

	sender.SendMessageMock.Expect(unknownMessage, int64(123)).Return(nil)
	userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(true)

	err := model.IncomingMessage(context.TODO(), Message{Text: "some text", UserID: 123})
	assert.NoError(t, err)
}
