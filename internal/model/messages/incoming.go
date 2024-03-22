package messages

import (
	orderedmap "github.com/wk8/go-ordered-map"
	"strings"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
	SendInlineMenu(text string, userID int64, buttons *orderedmap.OrderedMap) error
	EditTextAndMarkup(text string, userID int64, messageID int, buttons *orderedmap.OrderedMap) error
	DeleteMessage(userID int64, messageID int) error
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

func (s *Model) IncomingMessage(msg Message) error {
	switch msg.Text {
	case "/start":
		return s.tgClient.SendMessage("Hello", msg.UserID)
	case "/add":
		buttons := orderedmap.New()
		buttons.Set("Развлечение", "category#play")
		buttons.Set("Еда", "category#food")
		buttons.Set("Медицина", "category#medical")
		buttons.Set("Вещи", "category#things")
		buttons.Set("Другие", "category#another")
		buttons.Set("Закрыть", "close")
		return s.tgClient.SendInlineMenu("Добавляем новую трату\n\nКатегория:", msg.UserID, buttons)
	default:
		return s.tgClient.SendMessage("Не знаю эту команду", msg.UserID)
	}
}

func (s *Model) IncomingCallback(call CallBack) error {
	switch parts := strings.Split(call.Data, "#"); parts[0] {
	case "category":
		buttons := orderedmap.New()
		buttons.Set("Сегодня", "date#today#"+parts[1])
		buttons.Set("Указать", "date#select#"+parts[1])
		buttons.Set("Закрыть", "close")
		return s.tgClient.EditTextAndMarkup("Добавляем новую трату\n\nДата:", call.UserID, call.MessageID, buttons)
	case "close":
		return s.tgClient.DeleteMessage(call.UserID, call.MessageID)
	}
	return nil
}
