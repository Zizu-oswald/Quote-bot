package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *Database) {
	if update.Message != nil { // отпралено сообщение
		Chat, err := handleGettingUser(db, update.Message.Chat.ID)
		if err != nil {
			log.Println("Problem with getting user from message update: ", err)
		}
		switch update.Message.Text {
		case "/start", "/changelang":

			if Chat.LastMessageID != 0 {
				err := deleteMessage(Chat, bot) // удаляет сообщение при повторной попытке выбора языка
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

			err := handleGetQuote(Chat, bot, update)
			if err != nil {
				log.Println(Chat.ID, err)
			}
		}
	}

	if update.CallbackQuery != nil { // нажата кнопка в сообщении
		Chat, err := handleGettingUser(db, update.CallbackQuery.Message.Chat.ChatConfig().ChatID)
		if err != nil {
			log.Println("Problem with getting user from callback: ", err)
		}
		err = deleteMessage(Chat, bot)
		if err != nil {
			log.Println(err)
		}

		err = handleCallback(&Chat, bot, update.CallbackQuery)
		if err != nil {
			log.Println(err)
		}
		err = db.UpdateUserData(Chat)
		if err != nil {
			log.Println("cant update db (language): ", err)
		}
	}
}
