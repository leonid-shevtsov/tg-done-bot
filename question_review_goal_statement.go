package gtd_bot

const questionReviewGoalStatement = "review_goal/statement"

func init() {
	registerQuestion(questionReviewGoalStatement, askReviewGoalStatement, handleReviewGoalStatement)
}

func askReviewGoalStatement(i *interaction) {
	i.reply().text(i.locale.ReviewGoalStatement.Prompt).keyboard([][]string{
		{
			i.locale.Commands.Yes,
			i.locale.Commands.No,
		},
		{
			i.locale.Commands.TrashGoal,
			i.locale.Commands.BackToInbox,
		},
	}).send()
}

func handleReviewGoalStatement(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return questionReviewGoalDueDate
	case i.locale.Commands.No:
		return questionReviewGoalChangeStatement
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendText(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	default:
		return answerUnclear
	}
}
