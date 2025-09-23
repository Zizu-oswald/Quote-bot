package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


type ChatStruct struct {
	ID   int64
	Lang string
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

var Chat ChatStruct
var DeleteMessageID int // id сообщения которое будет удалено