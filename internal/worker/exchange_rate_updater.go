package worker

import (
	"context"
	"log"
	"time"
)

type CurrencyChangeUpdater interface {
	UpdateCurrency(ctx context.Context, time time.Time) error
}

type ConfigGetter interface {
	FrequencyExchangeRateUpdates() time.Duration
}

type CurrencyExchangeRateWorker struct {
	updater CurrencyChangeUpdater
	config  ConfigGetter
}

func NewCurrencyExchangeRateWorker(updater CurrencyChangeUpdater, config ConfigGetter) *CurrencyExchangeRateWorker {
	return &CurrencyExchangeRateWorker{
		updater: updater,
		config:  config,
	}
}

func (worker *CurrencyExchangeRateWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(worker.config.FrequencyExchangeRateUpdates())

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("stopped receiving exchange rates")
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					log.Println("stopped receiving exchange rates")
					return
				default:
					if err := worker.updater.UpdateCurrency(ctx, time.Now()); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
}
