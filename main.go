package main

import (
	"log"
	"os"

	"github.com/Zizu-oswald/Quote-bot/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("keys.env")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal("Failed to create a new bot")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u) // поток обновлений

	// var LastMessageID int
	for update := range updates {
		telegram.HandleUpdate(bot, update)
	}

}
