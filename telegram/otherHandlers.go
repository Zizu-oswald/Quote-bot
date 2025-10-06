package telegram

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Zizu-oswald/Quote-bot/mymemory"
	zenquotesapi "github.com/Zizu-oswald/Quote-bot/zenquotesAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleGetQuote(Chat ChatStruct, b *tgbotapi.BotAPI, u tgbotapi.Update) error {
	quote, err := zenquotesapi.GetRandomQuote()
	if err != nil {
		return fmt.Errorf("%s could not get a quote: %w", u.Message.From.FirstName, err)
	}
	log.Println(quote)

	if quote == (zenquotesapi.Quote{}) {
		msg := tgbotapi.NewMessage(Chat.ID, GetLocale(Chat.Lang).FindingQuotes) // сообщение заглушка когда кончились цитаты
		_, err := b.Send(msg)
		if err != nil {
			return err
		}
		return nil
	}

	if Chat.Lang == "ru" { // перевод цитаты и имени
		quote.Quote, err = mymemory.TranslEngToRus(quote.Quote)
		if err != nil {
			log.Println(err)
		}
		quote.Author, err = mymemory.TranslEngToRus(quote.Author)
		if err != nil {
			log.Println(err)
		}
	}

	quoteStr := quote.IntoStringForMessage(Chat.Lang)
	log.Println(u.Message.From.FirstName, "get a quote: ", quote)

	msg := tgbotapi.NewMessage(Chat.ID, quoteStr)
	_, err = b.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func handleCallback(Chat ChatStruct, b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) error {
	Chat.changeLanguage(cq) // исполнение смены языка

	msg := makeButton(Chat, GetLocale(Chat.Lang).ButtonGetQuote)
	_, err := b.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func handleGettingUser(db *Database, id int64) (ChatStruct, error) {
	var err error
	Chat, err := db.GetUser(id)
	if err == sql.ErrNoRows {
		err := db.AddUser(ChatStruct{ID: id, Lang: "en", LastMessageID: 0})
		if err != nil {
			return ChatStruct{}, fmt.Errorf("problem to adding user to db: %w", err)
		}
		Chat, err = db.GetUser(id)
		if err != nil {
			return ChatStruct{}, fmt.Errorf("problem to get user after adding: %w", err)
		}
	}
	if err != nil {
		return ChatStruct{}, fmt.Errorf("problem to get user after adding: %w", err)
	}
	return Chat, nil
}

