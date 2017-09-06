package gtd_bot

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RunBot(db *pg.DB) {
	bot, err := telegram.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") != "" {
		bot.Debug = true
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	resetUserState(bot, db)

	var updates telegram.UpdatesChannel

	if webhookURL := os.Getenv("WEBHOOK_URL"); webhookURL != "" {
		updates = runWebhook(bot, webhookURL)
	} else {
		updates = runLongPoll(bot)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		handleMessage(bot, update.Message, db)
	}
}

func runWebhook(bot *telegram.BotAPI, webhookURL string) telegram.UpdatesChannel {
	rand.Seed(time.Now().UnixNano())
	randomSeed := rand.Uint64()
	webhookPath := fmt.Sprintf("/%d", randomSeed)
	_, err := bot.SetWebhook(telegram.NewWebhook(webhookURL + webhookPath))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook(webhookPath)

	var httpAddress string
	if port := os.Getenv("WEBHOOK_PORT"); port != "" {
		httpAddress = ":" + port
	} else {
		httpAddress = ":80"
	}
	go http.ListenAndServe(httpAddress, nil)

	return updates
}

func runLongPoll(bot *telegram.BotAPI) telegram.UpdatesChannel {
	// delete webhook, otherwise you cannot long poll
	_, err := bot.SetWebhook(telegram.NewWebhook(""))
	if err != nil {
		panic(err)
	}
	// start long poll
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return updates
}
