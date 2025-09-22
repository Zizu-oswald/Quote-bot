package telegram

import (
	// "errors"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatStruct struct {
	ID   int64
	Lang string
}

var Chat ChatStruct
var deleteMessageID int

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Text == "/start" {
			Chat.ID = update.Message.Chat.ID
			msg := newMessageWithButtons(Chat.ID, "Change language:", "English", "Русский")
			delMsg, _ := bot.Send(msg)
			deleteMessageID = delMsg.MessageID
		}
	}
	if update.CallbackQuery != nil {
		delMsg := tgbotapi.NewDeleteMessage(Chat.ID, deleteMessageID)
		bot.Request(delMsg)
		handleCallback(bot, update.CallbackQuery)
	}

}

func newMessageWithButtons(ID int64, messageText string, butt1text string, butt2text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(ID, messageText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(butt1text, "en"),
			tgbotapi.NewInlineKeyboardButtonData(butt2text, "ru"),
		),
	)

	return msg
}

func handleCallback(b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) {
	// callback := tgbotapi.NewCallback(cq.ID, "Lang="+cq.Data)
	callback := tgbotapi.NewCallback(cq.ID, "")
	if _, err := b.Request(callback); err != nil {
		log.Println("Error: ", err)
	}

	switch cq.Data {
	case "ru":
		Chat.Lang = "ru"
		b.Send(tgbotapi.NewMessage(Chat.ID, "Выбран язык: Русский"))
	case "en":
		Chat.Lang = "en"
		b.Send(tgbotapi.NewMessage(Chat.ID, "Language selected: English"))
	default:
	}
}
