package messages

import (
	"testing"

	mocks "github.com/Svoevolin/workshop_1_bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	orderedmap "github.com/wk8/go-ordered-map"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("Hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}

func Test_OnAddCommand_ShouldAnswerWithInlineMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	buttons := orderedmap.New()
	buttons.Set("Развлечение", "play")
	buttons.Set("Еда", "food")
	buttons.Set("Медицина", "medical")
	buttons.Set("Вещи", "things")
	buttons.Set("Другие", "another")
	buttons.Set("Закрыть", "close")
	sender.EXPECT().SendInlineMenu("Добавляем новую трату\nВыбери категорию", int64(123), buttons)

	err := model.IncomingMessage(Message{
		Text:   "/add",
		UserID: 123,
	})

	assert.NoError(t, err)
}

func Test_OnUnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("Не знаю эту команду", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "some text",
		UserID: 123,
	})

	assert.NoError(t, err)
}
