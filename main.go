package main

import (
	"log"
	"os"

	"github.com/Zizu-oswald/Quote-bot/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

func main() {
	// .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// telegeram
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln("Failed to create a new bot: ", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u) // поток обновлений

	// postgres
	var db telegram.Database
	if err = db.ConnectToSql(); err != nil {
		log.Println(err)
	}
	if err = telegram.MakeTable(&db); err != nil {
		log.Println("cant make users table: ", err)
	}
	defer db.Close()

	// loop 
	for update := range updates {
		go telegram.HandleUpdate(bot, update, &db)
	}

}
