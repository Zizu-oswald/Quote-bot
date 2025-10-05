package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *Database) {
	if update.Message != nil { // отпралено сообщение
		switch update.Message.Text {
		case "/start", "/changelang":
			// обработка в бд
			handleGettingUser(db, update.Message.Chat.ID)

			if Chat.LastMessageID != 0 {
				err := deleteMessage(bot, Chat.LastMessageID) // удаляет сообщение при повторной попытке выбора языка
				if err != nil {
					log.Println(err)
				}
			}

			msg := newMessageWithButtons(Chat.ID, "Change language:", "English", "Русский")
			delMsg, err := bot.Send(msg)
			if err != nil {
				log.Println("Cant send message with buttons ", err)
			}
			Chat.LastMessageID = delMsg.MessageID
			err = db.UpdateUserData(Chat)
			if err != nil {
				log.Println("cant update db (lastmessageid): ", err)
			}

		case "Получить цитату", "Get quote":
			handleGettingUser(db, update.Message.Chat.ID)

			err := handleGetQuote(bot, update)
			if err != nil {
				log.Println(Chat.ID, err)
			}
		}
	}

	if update.CallbackQuery != nil { // нажата кнопка в сообщении
		err := deleteMessage(bot, Chat.LastMessageID)
		if err != nil {
			log.Println(err)
		}

		err = handleCallback(bot, update.CallbackQuery)
		if err != nil {
			log.Println(err)
		}
		err = db.UpdateUserData(Chat)
		if err != nil {
			log.Println("cant update db (language): ", err)
		}
	}
}
