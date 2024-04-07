package messages

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func (s *Model) setCurrency(ctx context.Context, msg Message) (string, error) {
	currencyCode := strings.Trim(strings.TrimPrefix(msg.Text, "/set_currency"), " ")
	userExists := s.userDB.UserExists(ctx, msg.UserID)

	supportedCurrencyCodesMap := make(map[string]any)

	for _, supportedCurrencyCode := range s.config.SupportedCurrencyCodes() {
		supportedCurrencyCodesMap[supportedCurrencyCode] = struct{}{}
	}

	if _, ok := supportedCurrencyCodesMap[currencyCode]; ok {
		if err := s.userDB.ChangeDefaultCurrency(ctx, msg.UserID, currencyCode); err != nil {
			log.Println(err)
			return "", ErrImpossibleToChangeUserCurrency
		}

		if userExists {
			return fmt.Sprintf("Установлена валюта по умолчанию %s", currencyCode), nil
		} else {
			return helpMessage, nil
		}
	}
	return fmt.Sprintf("Валюта %s не поддерживается, отправьте команду /set_currency с одним из значений %v",
		currencyCode, s.config.SupportedCurrencyCodes()), nil
}

func (s *Model) changeDefaultCurrency() (string, []map[string]string) {
	availableCodes := s.config.SupportedCurrencyCodes()

	rows := make(map[string]string, len(availableCodes))
	for _, availableCode := range availableCodes {
		rows[availableCode] = fmt.Sprintf("/set_currency %s", availableCode)
	}

	return "Выберите валюту в которой будете производить расходы", []map[string]string{rows}
}
