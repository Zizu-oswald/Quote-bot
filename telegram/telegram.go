package telegram

import (
	"log"

	zenquotesapi "github.com/Zizu-oswald/Quote-bot/zenquotesAPI"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		switch update.Message.Text {
		case "/start", "/changelang":
			Chat.ID = update.Message.Chat.ID

			if LastMessageID != 0 { 
				deleteMessage(bot) // удаляет сообщение при повторной попытке выбора языка
			}

			msg := newMessageWithButtons(Chat.ID, "Change language:", "English", "Русский")
			delMsg, err := bot.Send(msg)
			if err != nil {
				log.Println("Cant send message with buttons ", err)
			}
			LastMessageID = delMsg.MessageID

		case "Получить цитату", "Get quote":
			err := handleGetQuote(bot, update)
			if err != nil {
				log.Println(err) // FIXME: исправить вывод ошибок
			}
		}
	}

	if update.CallbackQuery != nil { // нажата кнопка в сообщении
		err := deleteMessage(bot)
		if err != nil {
			log.Println(err) // FIXME: исправить вывод ошибок
		}

		err = handleCallback(bot, update.CallbackQuery)
		if err != nil {
			log.Println(err) // FIXME: исправить вывод ошибок
		}
	}
}

func handleGetQuote(b *tgbotapi.BotAPI, u tgbotapi.Update) error {
	quote, err := zenquotesapi.GetRandomQuote()
	if err != nil {
		log.Println(u.Message.From.FirstName, "Could not get a quote: ", err)
	}
	quoteStr := quote.QuoteIntoString()
	log.Println(u.Message.From.FirstName, "get a quote: ", quote)
	msg := tgbotapi.NewMessage(Chat.ID, quoteStr)
	_, err = b.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func deleteMessage(b *tgbotapi.BotAPI) error {
	delMsg := tgbotapi.NewDeleteMessage(Chat.ID, LastMessageID) // запрос на удаление
	_, err := b.Request(delMsg)                                 // исполнение запроса на удаление
	if err != nil {
		return err
	}
	return nil
}

func handleCallback(b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) error {

	Chat.changeLanguage(cq) // исполнение смены языка

	var msg tgbotapi.MessageConfig
	switch Chat.Lang {
	case "ru":
		msg = makeButton("Получить цитату")
	case "en":
		msg = makeButton("Get quote")
	}
	_, err := b.Send(msg)
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
