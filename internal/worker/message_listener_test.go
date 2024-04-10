package worker

import (
	"context"
	"testing"
	"time"

	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/worker"
	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestMessageListenerWorkerRun(t *testing.T) {
	t.Skip()
	t.Run("after interrupting context processing new messages will stop", func(t *testing.T) {
		t.Parallel()
		m := minimock.NewController(t)
		messageFetcher := mock.NewMessageFetcherMock(m)
		messageProcessor := mock.NewMessageProcessorMock(m)

		chWithUpdates := make(chan tgbotapi.Update, 100)
		messageProcessor.IncomingMessageMock.
			Expect(minimock.AnyContext, messages.Message{Text: "/command", UserID: int64(1230)}).Return(nil)

		messageFetcher.StopMock.Expect()
		messageFetcher.StartMock.Expect().Return(chWithUpdates)

		worker := NewMessageListenerWorker(messageFetcher, messageProcessor)

		ctx, cancel := context.WithCancel(context.TODO())

		go func(ch chan<- tgbotapi.Update) {
			for i := 0; i < 10; i++ {
				time.Sleep(50 * time.Millisecond)
				if i == 5 {
					cancel()
				}
				ch <- tgbotapi.Update{Message: &tgbotapi.Message{From: &tgbotapi.User{ID: 1230}, Text: "/command"}}
			}
			close(chWithUpdates)
		}(chWithUpdates)
		worker.Run(ctx)

		assert.Error(t, ctx.Err())
	})
}
