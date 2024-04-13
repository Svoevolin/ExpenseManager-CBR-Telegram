package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	utils "github.com/Svoevolin/workshop_1_bot/internal/helpers/date"
	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/services"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCurrency(t *testing.T) {
	t.Run("returns error if gateway CBR could not get exchange rate on the date", func(t *testing.T) {
		t.Parallel()
		m := minimock.NewController(t)
		gateway := mock.NewExchangeRateFetcherMock(m)
		svc := NewExchangeRateUpdateSvc(gateway, nil, nil)

		date := utils.GetDate(time.Now())

		gateway.FetchRatesMock.Expect(minimock.AnyContext, date).
			Return(nil, fmt.Errorf("failed to get a list of currencies on the date %s", date.Format("02/01/2006")))

		err := svc.UpdateCurrency(context.TODO(), date)

		assert.EqualError(t, err, fmt.Sprintf("failed to get a list of currencies on the date %s", date.Format("02/01/2006")))
	})

	t.Run("we process only those currencies that are configured", func(t *testing.T) {
		t.Parallel()
		m := minimock.NewController(t)
		gateway := mock.NewExchangeRateFetcherMock(m)
		storage := mock.NewRateStorageMock(m)
		config := mock.NewConfigGetterMock(m)
		date := utils.GetDate(time.Now())
		svc := NewExchangeRateUpdateSvc(gateway, storage, config)

		config.SupportedCurrencyCodesMock.Expect().Return([]string{"RUB", "EUR"})
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "RUB", Original: "0", Date: date}).Then(nil)
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "EUR", Original: "0", Date: date}).Then(nil)

		gateway.FetchRatesMock.Expect(minimock.AnyContext, date).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "0"}, {Code: "EUR", Original: "0"}}, nil)

		err := svc.UpdateCurrency(context.TODO(), date)

		assert.NoError(t, err)
	})

	t.Run("penny converter should not block the saving of valid values", func(t *testing.T) {
		t.Parallel()
		m := minimock.NewController(t)
		gateway := mock.NewExchangeRateFetcherMock(m)
		storage := mock.NewRateStorageMock(m)
		config := mock.NewConfigGetterMock(m)
		date := utils.GetDate(time.Now())
		svc := NewExchangeRateUpdateSvc(gateway, storage, config)

		config.SupportedCurrencyCodesMock.Expect().Return([]string{"RUB", "USD", "EUR"})
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "RUB", Original: "0", Date: date}).Then(nil)
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "EUR", Original: "0", Date: date}).Then(nil)

		gateway.FetchRatesMock.Expect(minimock.AnyContext, date).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "Bad Value"}, {Code: "EUR", Original: "0"}}, nil)

		err := svc.UpdateCurrency(context.TODO(), date)

		assert.NoError(t, err)
	})

	t.Run("errors in saving currency rates should not block saving as a whole", func(t *testing.T) {
		t.Parallel()
		m := minimock.NewController(t)
		gateway := mock.NewExchangeRateFetcherMock(m)
		storage := mock.NewRateStorageMock(m)
		config := mock.NewConfigGetterMock(m)
		date := utils.GetDate(time.Now())
		svc := NewExchangeRateUpdateSvc(gateway, storage, config)

		config.SupportedCurrencyCodesMock.Expect().Return([]string{"RUB", "USD", "EUR"})
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "RUB", Original: "0", Date: date}).Then(errors.New("save failed"))
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "USD", Original: "0", Date: date}).Then(nil)
		storage.AddRateMock.When(minimock.AnyContext, domain.Rate{Code: "EUR", Original: "0", Date: date}).Then(nil)

		gateway.FetchRatesMock.Expect(minimock.AnyContext, date).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "0"}, {Code: "EUR", Original: "0"}}, nil)

		err := svc.UpdateCurrency(context.TODO(), date)

		assert.NoError(t, err)

	})
}
