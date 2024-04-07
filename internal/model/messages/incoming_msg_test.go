package messages

import (
	"context"
	"testing"

	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/messages"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestStartCommand(t *testing.T) {

	t.Run("if user starts using for first time then we have to show him message about choosing currency", func(t *testing.T) {

		m := minimock.NewController(t)
		userDB := mock.NewUserDBMock(m)
		settings := mock.NewConfigGetterMock(m)
		sender := mock.NewMessageSenderMock(m)
		model := New(sender, settings, userDB)

		exceptedCurrencies := []map[string]string{
			{
				"RUB": "/set_currency RUB",
				"USD": "/set_currency USD",
				"EUR": "/set_currency EUR",
			},
		}

		userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(false)
		settings.SupportedCurrencyCodesMock.Expect().Return([]string{"RUB", "USD", "EUR"})
		sender.SendMessageMock.Expect(newUserMessage, int64(123), exceptedCurrencies...).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{UserID: 123, Text: "/start"})
		assert.NoError(t, err)
	})

	t.Run("if user already exist we show help message", func(t *testing.T) {

		m := minimock.NewController(t)
		userDB := mock.NewUserDBMock(m)
		settings := mock.NewConfigGetterMock(m)
		sender := mock.NewMessageSenderMock(m)
		model := New(sender, settings, userDB)

		userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(true)
		sender.SendMessageMock.Expect(startMessage, int64(123)).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{UserID: 123, Text: "/start"})
		assert.NoError(t, err)
	})
}

func Test_OnUnknownCommand_ShouldAnswerWithUnknownMessage(t *testing.T) {
	m := minimock.NewController(t)
	userDB := mock.NewUserDBMock(m)
	sender := mock.NewMessageSenderMock(m)
	model := New(sender, nil, userDB)

	sender.SendMessageMock.Expect(unknownMessage, int64(123)).Return(nil)
	userDB.UserExistsMock.Expect(minimock.AnyContext, int64(123)).Return(true)

	err := model.IncomingMessage(context.TODO(), Message{Text: "some text", UserID: 123})
	assert.NoError(t, err)
}
