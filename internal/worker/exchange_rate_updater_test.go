package worker

import (
	"context"
	"sync"
	"testing"
	"time"

	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/worker"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyExchangeRateWorkerRun(t *testing.T) {
	t.Run("after interrupting context receiving for current exchange rates will stop", func(t *testing.T) {
		t.Parallel()

		m := minimock.NewController(t)

		rateUpdate := mock.NewCurrencyChangeUpdaterMock(m)
		config := mock.NewConfigGetterMock(m)

		config.FrequencyExchangeRateUpdatesMock.Expect().Return(30 * time.Millisecond)
		i := 0
		rateUpdate.UpdateCurrencyMock.Set(
			func(ctx context.Context, _ time.Time) error {
				assert.NoError(t, ctx.Err())
				rateUpdate.UpdateCurrencyBeforeCounter()
				i++
				return nil
			})

		worker := NewCurrencyExchangeRateWorker(rateUpdate, config)

		ctx, cancel := context.WithTimeout(context.TODO(), 140*time.Millisecond)
		defer cancel()

		var wg sync.WaitGroup

		wg.Add(1)
		worker.Run(ctx, &wg)
		wg.Wait()

		assert.Error(t, ctx.Err())
		assert.EqualValues(t, i, 4)
	})
}
