package gtd_bot

const questionReviewGoalDueDate = "review_goal/due_date"

func init() {
	registerQuestion(questionReviewGoalDueDate, askReviewGoalDueDate, handleReviewGoalDueDate)
}

func askReviewGoalDueDate(i *interaction) {
	i.sendPrompt(i.locale.ReviewGoalDueDate.Prompt, [][]string{
		{
			i.locale.Commands.Yes,
			i.locale.Commands.No,
		},
		{
			i.locale.Commands.TrashGoal,
			i.locale.Commands.BackToInbox,
		},
	})
}

func handleReviewGoalDueDate(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return completeGoalReview(i)
	case i.locale.Commands.No:
		return questionReviewGoalChangeDueDate
	case i.locale.Commands.TrashGoal:
		i.state.dropCurrentGoal()
		i.sendMessage(i.locale.Messages.GoalTrashed)
		return nextWorkQuestion(i)
	case i.locale.Commands.BackToInbox:
		return questionCollectingInbox
	default:
		return answerUnclear
	}
}

func completeGoalReview(i *interaction) string {
	i.state.markGoalReviewed()
	i.sendMessage(i.locale.ReviewGoal.Success)
	return nextWorkQuestion(i)
}
