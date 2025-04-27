package tg_done_bot

const questionReviewGoal = "review_goal"

func init() {
	registerQuestion(questionReviewGoal, askReviewGoal, handleReviewGoal)
}

func askReviewGoal(i *interaction) {
	i.sendText(i.locale.ReviewGoal.LetsReviewThisGoal)
	i.reply().goal(i.state.user.CurrentGoal).send()
	i.reply().text(i.locale.ReviewGoal.Prompt).keyboard([][]string{
		{i.locale.Commands.Yes},
		{i.locale.Commands.TrashGoal},
		{i.locale.Commands.BackToInbox},
	}).send()
}

func handleReviewGoal(i *interaction) string {
	switch i.message.Text {
	case i.locale.Commands.Yes:
		return questionReviewGoalStatement
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
