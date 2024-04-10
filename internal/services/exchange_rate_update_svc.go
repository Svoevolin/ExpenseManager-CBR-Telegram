package services

import (
	"context"
	"log"
	"time"

	"github.com/Svoevolin/workshop_1_bot/internal/domain"
	"github.com/Svoevolin/workshop_1_bot/internal/helpers/money"
)

type ExchangeRateFetcher interface {
	FetchRates(ctx context.Context, date time.Time) ([]domain.Rate, error)
}

type RateStorage interface {
	AddRate(ctx context.Context, date time.Time, rate domain.Rate) error
}

type ConfigGetter interface {
	SupportedCurrencyCodes() []string
}

type ExchangeRateUpdateSvc struct {
	gateway ExchangeRateFetcher
	storage RateStorage
	config  ConfigGetter
}

func NewExchangeRateUpdateSvc(gateway ExchangeRateFetcher, storage RateStorage, config ConfigGetter) *ExchangeRateUpdateSvc {
	return &ExchangeRateUpdateSvc{
		gateway: gateway,
		storage: storage,
		config:  config,
	}
}

func (svc *ExchangeRateUpdateSvc) UpdateCurrency(ctx context.Context, time time.Time) error {
	rates, err := svc.gateway.FetchRates(ctx, time)
	if err != nil {
		return err
	}

	supportedCurrencyCodes := svc.config.SupportedCurrencyCodes()
	supportedCurrencyCodesAsMap := make(map[string]struct{}, len(supportedCurrencyCodes))
	for _, code := range supportedCurrencyCodes {
		supportedCurrencyCodesAsMap[code] = struct{}{}
	}

	for _, rate := range rates {
		if _, ok := supportedCurrencyCodesAsMap[rate.Code]; !ok {
			continue
		}

		rate.Kopecks, err = money.ConvertStringAmountToKopecks(rate.Original)
		if err != nil {
			log.Println(err)
			continue
		}
		rate.Ts = time

		err = svc.storage.AddRate(ctx, time, rate)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
