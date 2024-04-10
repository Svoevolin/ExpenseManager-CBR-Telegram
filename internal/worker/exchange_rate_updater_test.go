package worker

import (
	"context"
	"testing"
)

func TestCurrencyExchangeRateWorker_Run(t *testing.T) {
	type fields struct {
		updater CurrencyChangeUpdater
		config  ConfigGetter
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker := &CurrencyExchangeRateWorker{
				updater: tt.fields.updater,
				config:  tt.fields.config,
			}
			worker.Run(tt.args.ctx)
		})
	}
}