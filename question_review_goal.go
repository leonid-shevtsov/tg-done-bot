package gtd_bot

const questionReviewGoal = "review_goal"

func init() {
	registerQuestion(questionReviewGoal, askReviewGoal, handleReviewGoal)
}

func askReviewGoal(i *interaction) {
	i.sendMessage(i.locale.ReviewGoal.LetsReviewThisGoal)
	i.sendGoal(i.state.user.CurrentGoal)
	i.sendPrompt(i.locale.ReviewGoal.Prompt, [][]string{
		{i.locale.Commands.Yes},
		{i.locale.Commands.TrashGoal},
		{i.locale.Commands.BackToInbox},
	})
}

func handleReviewGoal(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return questionReviewGoalStatement
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
