package gtd_bot

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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
	switch i.message.Text {
	case i.locale.Processing.TrashIt:
		return endProcessingByTrashing(i)
	case i.locale.WhatIsTheDueDate.None:
		return questionWhatIsTheNextAction
	case i.locale.WhatIsTheDueDate.Today:
		return setDueDate(i, endOfToday())
	case i.locale.WhatIsTheDueDate.Tomorrow:
		return setDueDate(i, endOfTomorrow())
	case i.locale.WhatIsTheDueDate.EndOfWeek:
		return setDueDate(i, endOfWeek())
	default:
		date, err := parseDateInput(i.message.Text)
		if err != nil {
			i.sendMessage(i.locale.WhatIsTheDueDate.FormatHelp)
			return questionWhatIsTheDueDate
		}
		return setDueDate(i, date)
	}
}

func setDueDate(i *interaction, date time.Time) string {
	i.state.setGoalDue(date)
	i.sendMessage(fmt.Sprintf(i.locale.WhatIsTheDueDate.Success, i.dueString(date)))
	return questionWhatIsTheNextAction
}

func parseDateInput(input string) (time.Time, error) {
	dateRegexp, err := regexp.Compile("(\\d{4})-(\\d{2})-(\\d{2})")
	if err != nil {
		panic("bad regular expression")
	}
	matches := dateRegexp.FindAllStringSubmatch(input, 1)
	if len(matches) == 0 {
		return time.Now(), errors.New("bad string")
	}
	year, _ := strconv.Atoi(matches[0][1])
	month, _ := strconv.Atoi(matches[0][2])
	day, _ := strconv.Atoi(matches[0][3])
	return time.Date(year, time.Month(month), day, 23, 59, 59, 0, time.UTC), nil
}

func endOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

func nextDay(t time.Time) time.Time {
	return t.Add(time.Duration(24) * time.Hour)
}

func endOfTomorrow() time.Time {
	return nextDay(endOfToday())
}

func endOfWeek() time.Time {
	endOfDay := endOfToday()
	for endOfDay.Weekday() != time.Sunday {
		endOfDay = nextDay(endOfDay)
	}
	return endOfDay
}
