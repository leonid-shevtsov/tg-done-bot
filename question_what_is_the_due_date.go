package gtd_bot

import (
	"fmt"
	"time"
)

const questionWhatIsTheDueDate = "process_inbox/what_is_the_due_date"

func init() {
	registerQuestion(questionWhatIsTheDueDate, askWhatIsTheDueDate, handleWhatIsTheDueDate)
}

func askWhatIsTheDueDate(i *interaction) {
	i.sendPrompt(i.locale.WhatIsTheDueDate.Prompt, [][]string{
		{
			i.locale.WhatIsTheDueDate.None,
			i.locale.WhatIsTheDueDate.Today,
		},
		{
			i.locale.WhatIsTheDueDate.Tomorrow,
			i.locale.WhatIsTheDueDate.EndOfWeek,
		},
		{i.locale.Processing.TrashIt},
	})
}

func handleWhatIsTheDueDate(i *interaction) string {
	date, err := recognizeDueDateFromMessage(i)
	if err != nil {
		switch i.message.Text {
		case i.locale.Processing.TrashIt:
			return endProcessingByTrashing(i)
		default:
			i.sendMessage(i.locale.WhatIsTheDueDate.FormatHelp)
			return questionWhatIsTheDueDate
		}
	} else if !date.IsZero() {
		i.state.setGoalDue(date)
		i.sendMessage(fmt.Sprintf(i.locale.WhatIsTheDueDate.Success, i.dueString(date)))
	}
	return questionWhatIsTheNextAction
}

func recognizeDueDateFromMessage(i *interaction) (time.Time, error) {
	switch i.message.Text {
	case i.locale.WhatIsTheDueDate.None:
		return time.Time{}, nil
	case i.locale.WhatIsTheDueDate.Today:
		return endOfToday(), nil
	case i.locale.WhatIsTheDueDate.Tomorrow:
		return endOfTomorrow(), nil
	case i.locale.WhatIsTheDueDate.EndOfWeek:
		return endOfWeek(), nil
	default:
		return parseDateInput(i.message.Text)
	}
}
