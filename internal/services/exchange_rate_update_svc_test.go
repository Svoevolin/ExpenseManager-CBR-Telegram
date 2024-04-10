package services

import (
	"context"
	"testing"
	"time"
)

func TestExchangeRateUpdateSvc_UpdateCurrency(t *testing.T) {
	type fields struct {
		gateway ExchangeRateFetcher
		storage RateStorage
		config  ConfigGetter
	}
	type args struct {
		ctx  context.Context
		time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &ExchangeRateUpdateSvc{
				gateway: tt.fields.gateway,
				storage: tt.fields.storage,
				config:  tt.fields.config,
			}
			if err := svc.UpdateCurrency(tt.args.ctx, tt.args.time); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}