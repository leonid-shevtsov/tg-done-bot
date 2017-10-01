package gtd_bot

import (
	"time"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

const resetTimeout = time.Duration(10) * time.Minute

func resetWhenInactive(bot *telegram.BotAPI, db *pg.DB) {
	for {
		<-time.After(time.Until(resetAndReturnNextTime(bot, db)))
	}
}

func resetAndReturnNextTime(bot *telegram.BotAPI, db *pg.DB) time.Time {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	for _, user := range repo.usersDirtySince(time.Now().Add(-resetTimeout)) {
		i := initiateInteraction(bot, repo, user.ID)
		i.sendText(i.locale.Messages.BackToCollectingInbox)
		user.ActiveQuestion = questionCollectingInbox
		repo.update(user)
		askCollectingInbox(i)
	}

	earliestDirtyActivityTime := repo.earliestDirtyActivityTime()

	if earliestDirtyActivityTime.IsZero() {
		earliestDirtyActivityTime = time.Now()
	}

	return earliestDirtyActivityTime.Add(resetTimeout)
}
