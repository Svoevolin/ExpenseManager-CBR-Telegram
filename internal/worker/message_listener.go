package worker

import (
	"context"
	"log"

	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Stop()
}

type MessageProcessor interface {
	IncomingMessage(ctx context.Context, msg messages.Message) error
}

type MessageListenerWorker struct {
	fetcher   MessageFetcher
	processor MessageProcessor
}

func NewMessageListenerWorker(fetcher MessageFetcher, processor MessageProcessor) *MessageListenerWorker {
	return &MessageListenerWorker{
		fetcher:   fetcher,
		processor: processor,
	}
}

func (worker *MessageListenerWorker) processing(ctx context.Context, update tgbotapi.Update) error {

	if update.Message != nil {
		log.Printf("<message>[%s][%d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

		err := worker.processor.IncomingMessage(ctx, messages.Message{
			Text:   update.Message.Text,
			UserID: update.Message.From.ID,
		})
		if err != nil {
			log.Println("error processing message:", err)
		}
	} else if update.CallbackQuery != nil {
		log.Printf("<callback>[%s][%d] %s", update.CallbackQuery.From.UserName,
			update.CallbackQuery.From.ID, update.CallbackQuery.Data)

		err := worker.processor.IncomingMessage(ctx, messages.Message{
			Text:   update.CallbackQuery.Data,
			UserID: update.CallbackQuery.From.ID,
		})
		if err != nil {
			log.Println("error processing callback:", err)
		}
	}
	return nil
}

func (worker *MessageListenerWorker) Run(ctx context.Context) {
	for update := range worker.fetcher.Start() {
		select {
		case <-ctx.Done():
			worker.fetcher.Stop()
			log.Println("stopped receiving new message from tg bot")
			return
		default:
			if err := worker.processing(ctx, update); err != nil {
				log.Println(err)
			}
		}
	}
}
