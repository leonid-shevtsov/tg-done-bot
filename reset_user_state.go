package gtd_bot

import (
	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"leonid.shevtsov.me/gtd_bot/i18n"
)

func resetUserState(bot *telegram.BotAPI, db *pg.DB) {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	for _, user := range repo.usersInDirtyState() {
		state := newState(repo, user.ID)
		interaction := interaction{state: state, bot: bot, message: nil, locale: &i18n.En}
		interaction.sendMessage(interaction.locale.Messages.ServerRestart)
		user.ActiveQuestion = questionCollectingInbox
		repo.update(user)
		askCollectingInbox(&interaction)
	}
}
