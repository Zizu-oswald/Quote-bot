package main

import (
	// "encoding/json"
	// "net/http"
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

	for update := range updates {
		telegram.HandleUpdate(bot, update)
	}

}

// type Quote []struct {
// 	Q string `json:"q"`
// 	A string `json:"a"`
// }

// url := "https://zenquotes.io/api/random"
// resp, err := http.Get(url)
// if err != nil || resp.StatusCode != 200 {
// 	log.Fatal("trash")
// }

// var quote Quote
// err = json.NewDecoder(resp.Body).Decode(&quote)
// if err != nil {
// 	log.Fatal(err)
// }
