package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewHandler(botToken string) (*Handler, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	chatID, err := listen(bot)
	if err != nil {
		return nil, err
	}

	return &Handler{bot: bot, chatID: chatID}, nil
}

func listen(bot *tgbotapi.BotAPI) (int64, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			return chatID, nil
		}
	}
	return 0, errors.New("cant listen chatID")
}

func (h *Handler) SendMessage(message string) error {
	msg := tgbotapi.NewMessage(h.chatID, message)

	_, err := h.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Text == "/start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
		bot.Send(msg)
	}
}
