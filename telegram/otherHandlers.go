package telegram

import (
	"fmt"
	"log"

	"github.com/Zizu-oswald/Quote-bot/mymemory"
	zenquotesapi "github.com/Zizu-oswald/Quote-bot/zenquotesAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleGetQuote(b *tgbotapi.BotAPI, u tgbotapi.Update) error {
	quote, err := zenquotesapi.GetRandomQuote()
	if err != nil {
		return fmt.Errorf("%s could not get a quote: %e", u.Message.From.FirstName, err)
	}
	log.Println(quote)

	if quote == (zenquotesapi.Quote{}) {
		err = messageFindingQuotes(b) // сообщение заглушка когда кончились цитаты
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

func handleCallback(b *tgbotapi.BotAPI, cq *tgbotapi.CallbackQuery) error {
	Chat.changeLanguage(cq) // исполнение смены языка

	msg := makeButton(GetLocale(Chat.Lang).ButtonGetQuote)
	_, err := b.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
