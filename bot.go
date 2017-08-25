package gtd_bot

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RunBot(db *pg.DB) {
	bot, err := telegram.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		interaction := newInteraction(db, bot, update.Message)
		interaction.handleMessage()
	}
}
