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

func (c *Client) SendInlineMenu(text string, userID int64, buttons *orderedmap.OrderedMap) error {
	tgButtonsRows := make([][]tgbotapi.InlineKeyboardButton, 0, 100)
	for pair := buttons.Oldest(); pair != nil; pair = pair.Next() {
		tgButtonsRows = append(tgButtonsRows,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(pair.Key.(string), pair.Value.(string))))
	}
	newMessage := tgbotapi.NewMessage(userID, text)
	newMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgButtonsRows...)
	_, err := c.client.Send(newMessage)
	if err != nil {
		return errors.Wrap(err, "client.Send")
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
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(messages.Message{
				Text:   update.Message.Text,
				UserID: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message:", err)
			}
		}
	}
}
