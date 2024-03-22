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
	buttons.Set("Развлечение", "category#play")
	buttons.Set("Еда", "category#food")
	buttons.Set("Медицина", "category#medical")
	buttons.Set("Вещи", "category#things")
	buttons.Set("Другие", "category#another")
	buttons.Set("Закрыть", "close")
	sender.EXPECT().SendInlineMenu("Добавляем новую трату\n\nКатегория:", int64(123), buttons)

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

func Test_OnCategoryCallback_ShouldEditTextAndMarkup(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	option := "food"
	buttons := orderedmap.New()
	buttons.Set("Сегодня", "date#today#"+option)
	buttons.Set("Указать", "date#select#"+option)
	buttons.Set("Закрыть", "close")
	sender.EXPECT().EditTextAndMarkup("Добавляем новую трату\n\nДата:", int64(123), 321, buttons)

	err := model.IncomingCallback(CallBack{
		UserID:    123,
		MessageID: 321,
		Data:      "category#" + option,
	})

	assert.NoError(t, err)
}

func Test_OnCloseCallback_ShouldDeleteMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().DeleteMessage(int64(123), 321)

	err := model.IncomingCallback(CallBack{
		UserID:    123,
		MessageID: 321,
		Data:      "close",
	})

	assert.NoError(t, err)

}
