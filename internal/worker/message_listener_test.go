package worker

import (
	"context"
	"testing"
)

func TestMessageListenerWorker_Run(t *testing.T) {
	type fields struct {
		fetcher   MessageFetcher
		processor MessageProcessor
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
			worker := &MessageListenerWorker{
				fetcher:   tt.fields.fetcher,
				processor: tt.fields.processor,
			}
			worker.Run(tt.args.ctx)
		})
	}
}