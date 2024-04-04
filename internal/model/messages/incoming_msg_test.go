package messages

import (
	"context"
	"testing"

	mock "github.com/Svoevolin/workshop_1_bot/internal/mocks/messages"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	m := minimock.NewController(t)
	sender := mock.NewMessageSenderMock(m)
	model := New(sender)

	sender.SendMessageMock.Expect("Hello", int64(123)).Return(nil)
	err := model.IncomingMessage(context.TODO(), Message{
		Text:   "/start",
		UserID: 123,
	})
	assert.NoError(t, err)
}
