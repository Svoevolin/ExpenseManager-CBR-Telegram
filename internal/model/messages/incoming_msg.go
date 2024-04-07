package messages

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

func (s *Model) IncomingMessage(ctx context.Context, msg Message) error {

	switch {
	case !s.userDB.UserExists(ctx, msg.UserID) && !strings.HasPrefix(msg.Text, "/set_currency"):

		_, buttons := s.changeDefaultCurrency()
		return s.tgClient.SendMessage(newUserMessage, msg.UserID, buttons...)

	case msg.Text == "/start":

		return s.tgClient.SendMessage(startMessage, msg.UserID)

	case strings.HasPrefix(msg.Text, "/set_currency"):

		answer, err := s.setCurrency(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrImpossibleToChangeUserCurrency) {
			return s.tgClient.SendMessage(FailedChangeCurrencyMessage, msg.UserID)
		}

		//fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)

	case strings.HasPrefix(msg.Text, "/change_currency"):

		answer, buttons := s.changeDefaultCurrency()
		return s.tgClient.SendMessage(answer, msg.UserID, buttons...)

	}

	return s.tgClient.SendMessage(unknownMessage, msg.UserID)
}
