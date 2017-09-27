package gtd_bot

const questionReviewGoal = "review_goal"

func init() {
	registerQuestion(questionReviewGoal, askReviewGoal, handleReviewGoal)
}

func askReviewGoal(i *interaction) {
	if goal := i.state.goalToReview(); goal != nil {
		i.state.setCurrentGoal(goal)
		i.sendMessage(i.locale.ReviewGoal.LetsReviewThisGoal)
		i.sendGoal(goal)
		i.sendPrompt(i.locale.ReviewGoal.Prompt, [][]string{
			{i.locale.Commands.Yes},
			{i.locale.Commands.TrashGoal},
			{i.locale.Commands.BackToInbox},
		})
	} else {
		panic("bad precondition for review_goal question")
	}
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
