package gtd_bot

import (
	"fmt"
	"time"
)

func (i *interaction) dueString(dueDate time.Time) string {
	until := time.Until(dueDate)
	if until < 0 {
		return fmt.Sprintf("%s (%s!)", dueDate.Format("2006-01-02"), i.locale.Date.Late)
	} else if until < time.Duration(24)*time.Hour {
		return i.locale.Date.Today
	} else if until < time.Duration(48)*time.Hour {
		return i.locale.Date.Tomorrow
	}

	return dueDate.Format("2006-01-02")
}
