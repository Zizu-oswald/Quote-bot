package telegram

import (
	// "errors"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Text == "/start" {
			msg := newMessageWithButtons(update.Message.Chat.ID, "Change language:", "English", "Русский")
			bot.Send(msg)
		}
	}
	if update.CallbackQuery != nil {
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
	callback := tgbotapi.NewCallback(cq.ID, "Lang="+cq.Data)
	if _, err := b.Request(callback); err != nil {
		log.Println("Error: ", err)
	}
	b.Send(tgbotapi.NewMessage(cq.Message.Chat.ID, "Lang: "+cq.Data))
}
