package gtd_bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func dailyUpdates(bot *telegram.BotAPI, db *pg.DB) {
	for {
		sendDailyUpdates(bot, db)
		<-time.After(time.Duration(24) * time.Hour)
	}
}

func sendDailyUpdates(bot *telegram.BotAPI, db *pg.DB) {
	repo := newRepo(db)
	defer repo.finalizeTransaction()

	for _, user := range repo.usersForDailyUpdate() {
		i := initiateInteraction(bot, repo, user.ID)
		messageLines := []string{
			fmt.Sprintf("<b>%s</b>", i.locale.StatusUpdate.Title),
			"",
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.GoalsTotal, i.state.goalCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.ActionsTotal, i.state.actionCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.WaitingForTotal, i.state.waitingForCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.InboxItemsProcessedToday, i.state.inboxItemsProcessedTodayCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.GoalsCreatedToday, i.state.goalsCreatedTodayCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.ActionsCompletedToday, i.state.actionsCompletedTodayCount()),
			fmt.Sprintf("%s: <b>%d</b>", i.locale.StatusUpdate.InboxItemsCount, i.state.inboxCount()),
		}
		message := strings.Join(messageLines, "\n")
		i.sendHTMLMessage(message)
		askCollectingInbox(i)
	}
}
