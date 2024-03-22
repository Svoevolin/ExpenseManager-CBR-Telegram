package tg

import (
	"log"

	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	orderedmap "github.com/wk8/go-ordered-map"
)

type TokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func CreateInlineMenu(buttons *orderedmap.OrderedMap) tgbotapi.InlineKeyboardMarkup {
	tgButtonsRows := make([][]tgbotapi.InlineKeyboardButton, 0, 100)
	for pair := buttons.Oldest(); pair != nil; pair = pair.Next() {
		tgButtonsRows = append(tgButtonsRows,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(pair.Key.(string), pair.Value.(string))))
	}
	return tgbotapi.NewInlineKeyboardMarkup(tgButtonsRows...)
}

func (c *Client) SendInlineMenu(text string, userID int64, buttons *orderedmap.OrderedMap) error {
	newMessage := tgbotapi.NewMessage(userID, text)
	newMessage.ReplyMarkup = CreateInlineMenu(buttons)
	_, err := c.client.Send(newMessage)
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) EditTextAndMarkup(text string, userID int64, messageID int, buttons *orderedmap.OrderedMap) error {
	_, err := c.client.Send(tgbotapi.NewEditMessageTextAndMarkup(userID, messageID, text, CreateInlineMenu(buttons)))
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) DeleteMessage(userID int64, messageID int) error {
	_, err := c.client.Request(tgbotapi.NewDeleteMessage(userID, messageID))
	if err != nil {
		return errors.Wrap(err, "client.Request")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil {
			log.Printf("<message>[%s][%d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

			err := msgModel.IncomingMessage(messages.Message{
				Text:   update.Message.Text,
				UserID: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message:", err)
			}
		} else if update.CallbackQuery != nil {
			log.Printf("<callback>[%s][%d] %s", update.CallbackQuery.From.UserName,
				update.CallbackQuery.From.ID, update.CallbackQuery.Data)

			err := msgModel.IncomingCallback(messages.CallBack{
				UserID:    update.CallbackQuery.From.ID,
				MessageID: update.CallbackQuery.Message.MessageID,
				Data:      update.CallbackQuery.Data,
			})
			if err != nil {
				log.Println("error processing callback:", err)
			}
		}
	}
}
