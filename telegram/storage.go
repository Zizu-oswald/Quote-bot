package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatStruct struct {
	ID            int64
	Lang          string
	LastMessageID int
}

func (ch *ChatStruct) changeLanguage(cq *tgbotapi.CallbackQuery) {
	switch cq.Data {
	case "ru":
		ch.Lang = "ru"
	case "en":
		ch.Lang = "en"
	default:
	}
}

type Localizations struct {
	ButtonGetQuote string
	LanguageMsg    string
	FindingQuotes  string
}

var locales = map[string]Localizations{
	"ru": {
		ButtonGetQuote: "Получить цитату",
		LanguageMsg:    "Выбран язык: Русский",
		FindingQuotes:  "Ищем цитаты, попробуйте позже.",
	},
	"en": {
		ButtonGetQuote: "Get quote",
		LanguageMsg:    "Language selected: English",
		FindingQuotes:  "We are looking for quotes, try later.",
	},
}

func GetLocale(lang string) Localizations {
	if loc, ok := locales[lang]; ok {
		return loc
	}
	return locales["en"]
}

// var Chat ChatStruct // FIXME:
