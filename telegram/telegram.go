package telegram

import (
	// "errors"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatStruct struct {
	ID          int64
	Lang        string
	langChanged bool // изменить
}

var Chat ChatStruct
var deleteMessageID int // id сообщения которое будет удалено

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		if (update.Message.Text == "/start") || (update.Message.Text == "/changelang") {
			Chat.ID = update.Message.Chat.ID
			msg := newMessageWithButtons(Chat.ID, "Change language:", "English", "Русский")
			delMsg, err := bot.Send(msg)
			if err != nil {
				log.Println("Cant send message with buttons ", err)
			}
			deleteMessageID = delMsg.MessageID
		} else if (update.Message.Text == "Получить цитату") || (update.Message.Text == "Get quote") {
			msg := tgbotapi.NewMessage(Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}

	if update.CallbackQuery != nil {
		delMsg := tgbotapi.NewDeleteMessage(Chat.ID, deleteMessageID) // запрос на удаление
		_, err := bot.Request(delMsg)                                 // исполнение запроса на удаление
		if err != nil {
			log.Println("Cant delete message ", err)
		}
		handleCallback(bot, update.CallbackQuery)
		Chat.langChanged = true // обозначаем что язык был изменен
	}
//че за бретто
	if Chat.langChanged { // если язык был изменен то изменяем кнопку
		Chat.langChanged = false
		log.Println(Chat.Lang)
		var msg tgbotapi.MessageConfig
		switch Chat.Lang {
		case "ru":
			msg = makeButton("Получить цитату")
		case "en":
			msg = makeButton("Get quote")
		}
		bot.Send(msg)
	}
}

func makeButton(str string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(Chat.ID, ".")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(str),
		),
	)
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

func handleCallback(b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(cq.ID, "") // создание колбэка чтобы кнопки перестали гореть
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
