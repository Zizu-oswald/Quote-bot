package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		switch update.Message.Text {
		case "/start", "/changelang":
			Chat.ID = update.Message.Chat.ID
			msg := newMessageWithButtons(Chat.ID, "Change language:", "English", "Русский")
			delMsg, err := bot.Send(msg)
			if err != nil {
				log.Println("Cant send message with buttons ", err)
			}
			DeleteMessageID = delMsg.MessageID
		case "Получить цитату", "Get quote":
			msg := tgbotapi.NewMessage(Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}

	if update.CallbackQuery != nil { // нажата кнопка в сообщении
		HandleCallback(bot, update.CallbackQuery)
	}
}

func HandleCallback(b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) {
	delMsg := tgbotapi.NewDeleteMessage(Chat.ID, DeleteMessageID) // запрос на удаление 
	_, err := b.Request(delMsg)                                   // исполнение запроса на удаление
	if err != nil {
		log.Println("Cant delete message ", err)
	}
	// FIXME: удаление отдельно

	Chat.changeLanguage(cq) // исполнение смены языка

	var msg tgbotapi.MessageConfig
	switch Chat.Lang {
	case "ru":
		msg = makeButton("Получить цитату")
	case "en":
		msg = makeButton("Get quote")
	}
	b.Send(msg)

}

func makeButton(str string) tgbotapi.MessageConfig {
	msg := makeLanguageMsg() //создание сообщения о смене языка
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(str),
		),
	)
	return msg
}

func makeLanguageMsg() tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	switch Chat.Lang {
	case "ru":
		Chat.Lang = "ru"
		msg = tgbotapi.NewMessage(Chat.ID, "Выбран язык: Русский")
	case "en":
		Chat.Lang = "en"
		msg = tgbotapi.NewMessage(Chat.ID, "Language selected: English")
	default:
	}
	return msg
}

func newMessageWithButtons(ID int64, messageText string, butt1text string, butt2text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(ID, messageText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup( // создание кнопок в сообщении
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(butt1text, "en"),
			tgbotapi.NewInlineKeyboardButtonData(butt2text, "ru"),
		),
	)

	return msg
}
