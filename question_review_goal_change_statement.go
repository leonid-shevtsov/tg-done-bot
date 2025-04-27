package tg_done_bot

const questionReviewGoalChangeStatement = "review_goal/change_statement"

func init() {
	registerQuestion(questionReviewGoalChangeStatement, askReviewGoalChangeStatement, handleReviewGoalChangeStatement)
}

func askReviewGoalChangeStatement(i *interaction) {
	i.reply().text(i.locale.ReviewGoalChangeStatement.Prompt).keyboard([][]string{
		{
			i.locale.Commands.Keep,
			i.locale.Commands.TrashGoal,
			i.locale.Commands.BackToInbox,
		},
	}).send()
}

func handleReviewGoalChangeStatement(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Keep:
		return questionReviewGoalDueDate
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendText(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	default:
		i.state.setGoalStatement(i.message.Text)
		return questionReviewGoalDueDate
	}
}
