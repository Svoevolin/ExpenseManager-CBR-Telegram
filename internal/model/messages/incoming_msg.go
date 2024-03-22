package messages

import orderedmap "github.com/wk8/go-ordered-map"

type MessageSender interface {
	SendMessage(text string, userID int64) error
	SendInlineMenu(text string, userID int64, buttons *orderedmap.OrderedMap) error
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

func (s *Model) IncomingMessage(msg Message) error {
	switch msg.Text {
	case "/start":
		return s.tgClient.SendMessage("Hello", msg.UserID)
	case "/add":
		buttons := orderedmap.New()
		buttons.Set("Развлечение", "play")
		buttons.Set("Еда", "food")
		buttons.Set("Медицина", "medical")
		buttons.Set("Вещи", "things")
		buttons.Set("Другие", "another")
		buttons.Set("Закрыть", "close")
		return s.tgClient.SendInlineMenu("Добавляем новую трату\nВыбери категорию", msg.UserID, buttons)
	default:
		return s.tgClient.SendMessage("Не знаю эту команду", msg.UserID)
	}
}
