package telegram

import (
	"database/sql"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *Database) {
	if update.Message != nil { // отпралено сообщение
		switch update.Message.Text {
		case "/start", "/changelang":
			// обработка в бд
			var err error
			Chat, err = db.TakeUser(update.Message.Chat.ID)
			if err == sql.ErrNoRows {
				err := db.AddUser(ChatStruct{ID: update.Message.Chat.ID, Lang: "en", LastMessageID: 0})
				if err != nil {
					log.Println("Problem to adding user to db: ", err)
				}
				Chat, err = db.TakeUser(update.Message.Chat.ID)
				if err != nil {
					log.Println("Problem to taking user from db: ", err)
				}
			}
			if err != nil {
				log.Println("Problem to taking user from db: ", err)
			}
			
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

		case "Получить цитату", "Get quote":
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
	}
}
