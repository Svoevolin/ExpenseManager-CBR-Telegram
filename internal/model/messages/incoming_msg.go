package messages

import (
	"context"
)

type MessageSender interface {
	SendMessage(text string, userID int64, keyboardRows ...map[string]string) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

type CallBack struct {
	UserID    int64
	MessageID int
	Data      string
}

func (s *Model) IncomingMessage(ctx context.Context, msg Message) error {
	switch msg.Text {
	case "/start":
		return s.tgClient.SendMessage("Hello", msg.UserID)
	default:
		return s.tgClient.SendMessage("Не знаю эту команду", msg.UserID)
	}
}
