package messages

import (
	"context"
	"errors"
)

const (
	newUserMessage = `Привет, я бот - менеджер твоей бухгалтерии. Мне нужно знать валюту которой ты платишь`
	startMessage   = "Привет, я бот - менеджер твоей бухгалтерии.\n\n" + helpMessage
	helpMessage    = `Прочитай команды, чтобы понять что я умею делать:
/change_currency - Изменить валюту расходов
/add сумма; описание; <дата> - Добавь новую трату, если не укажешь дату - будет сегодня

Посмотреть расходы:
/spent - за всё время
/spent_day - за день 
/spent_week - за неделю
/spent_month - за месяц
/spent_year - за год`
	unknownMessage = `Неизвестная команда. Чтобы посмотреть список команд отправь /help`

	FailedMessage               = "Я временно не работаю, повторите попытку позже"
	FailedChangeCurrencyMessage = "Не удалось изменить текущую валюту, повторите попытку позже"
)

var (
	ErrImpossibleToChangeUserCurrency = errors.New("failed to change user currency")
)

type MessageSender interface {
	SendMessage(text string, userID int64, keyboardRows ...map[string]string) error
}

type ConfigGetter interface {
	SupportedCurrencyCodes() []string
	GetBaseCurrency() string
}

type UserDB interface {
	UserExists(ctx context.Context, userID int64) bool
	ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error
	GetDefaultCurrency(ctx context.Context, userID int64) (string, error)
}

type Model struct {
	tgClient MessageSender
	config   ConfigGetter
	userDB   UserDB
}

func New(tgClient MessageSender, config ConfigGetter, userDB UserDB) *Model {
	return &Model{
		tgClient: tgClient,
		config:   config,
		userDB:   userDB,
	}
}

type Message struct {
	Text   string
	UserID int64
}
