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

	case msg.Text == "/help":
		return s.tgClient.SendMessage(helpMessage, msg.UserID)

	case strings.HasPrefix(msg.Text, "/add"):
		answer, err := s.addExpense(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrInvalidCommand) {
			return s.tgClient.SendMessage(InvalidCommandMessage, msg.UserID)
		}

		if errors.Is(err, ErrInvalidAmount) {
			return s.tgClient.SendMessage(InvalidAmountMessage, msg.UserID)
		}

		if errors.Is(err, ErrInvalidDate) {
			return s.tgClient.SendMessage(InvalidDateMessage, msg.UserID)
		}

		if errors.Is(err, ErrWriteToDatabase) {
			return s.tgClient.SendMessage(FailedWriteMessage, msg.UserID)
		}

		//fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)

	case strings.HasPrefix(msg.Text, "/spent"):
		answer, err := s.listExpenses(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrGetRecordsInDatabase) {
			return s.tgClient.SendMessage(FailedGetListExpensesMessage, msg.UserID)
		}

		//fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)

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
