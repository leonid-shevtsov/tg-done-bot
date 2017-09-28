package gtd_bot

import (
	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func resetUserState(bot *telegram.BotAPI, db *pg.DB) {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	for _, user := range repo.usersInDirtyState() {
		interaction := initiateInteraction(bot, repo, user.ID)
		interaction.sendMessage(interaction.locale.Messages.ServerRestart)
		user.ActiveQuestion = questionCollectingInbox
		repo.update(user)
		askCollectingInbox(interaction)
	}
}
