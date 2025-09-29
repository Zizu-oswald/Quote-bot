package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func deleteMessage(b *tgbotapi.BotAPI, LastMessageID int) error {
	delMsg := tgbotapi.NewDeleteMessage(Chat.ID, LastMessageID) // запрос на удаление
	_, err := b.Request(delMsg)                                 // исполнение запроса на удаление
	if err != nil {
		return err
	}
	return nil
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

func messageFindingQuotes(b *tgbotapi.BotAPI) error {
	var msg tgbotapi.MessageConfig
	switch Chat.Lang {
	case "ru":
		msg = tgbotapi.NewMessage(Chat.ID, "Ищем цитаты, попробуйте позже.")
	case "en":
		msg = tgbotapi.NewMessage(Chat.ID, "We are looking for quotes, try later.")
	}
	_, err := b.Send(msg)
	return err
}
