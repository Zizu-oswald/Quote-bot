package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Zizu-oswald/Quote-bot/telegram"

	"github.com/joho/godotenv"
)

type Quote []struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func main() {
	err := godotenv.Load("keys.env")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := telegram.NewHandler(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	url := "https://zenquotes.io/api/random"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("trash")
	}

	var quote Quote
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		log.Fatal(err)
	}

	bot.SendMessage(fmt.Sprintf("Quote: %s \nAuthor: %s", quote[0].Q, quote[0].A))

	// select {}
}
