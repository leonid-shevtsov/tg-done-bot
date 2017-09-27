package gtd_bot

import "fmt"

const questionReviewGoalChangeDueDate = "review_goal/change_due_date"

func init() {
	registerQuestion(questionReviewGoalChangeDueDate, askReviewGoalChangeDueDate, handleReviewGoalChangeDueDate)
}

func askReviewGoalChangeDueDate(i *interaction) {
	i.sendPrompt(i.locale.ReviewGoalChangeDueDate.Prompt, [][]string{
		{
			i.locale.WhatIsTheDueDate.None,
			i.locale.WhatIsTheDueDate.Today,
		},
		{
			i.locale.WhatIsTheDueDate.Tomorrow,
			i.locale.WhatIsTheDueDate.EndOfWeek,
		},
		{
			i.locale.Commands.Keep,
			i.locale.Commands.TrashGoal,
			i.locale.Commands.BackToInbox,
		},
	})
}

func handleReviewGoalChangeDueDate(i *interaction) string {
	date, err := recognizeDueDateFromMessage(i)
	if err != nil {
		switch i.message.Text {
		case i.locale.Commands.Keep:
			return questionReviewGoalDueDate
		case i.locale.Commands.TrashGoal:
			i.state.dropCurrentGoal()
			i.sendMessage(i.locale.Messages.GoalTrashed)
			return nextWorkQuestion(i)
		default:
			i.sendMessage(i.locale.WhatIsTheDueDate.FormatHelp)
			return questionReviewGoalDueDate
		}
	}

	i.state.setGoalDue(date)
	if !date.IsZero() {
		i.sendMessage(fmt.Sprintf(i.locale.WhatIsTheDueDate.Success, i.dueString(date)))
	} else {
		i.sendMessage(i.locale.ReviewGoalChangeDueDate.Cleared)
	}
	return completeGoalReview(i)
}
