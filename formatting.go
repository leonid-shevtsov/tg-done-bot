package gtd_bot

import (
	"fmt"
	"strings"
	"time"

	"leonid.shevtsov.me/gtd_bot/i18n"
)

func dueString(locale *i18n.Locale, dueDate time.Time) string {
	until := time.Until(dueDate)
	if until < 0 {
		return fmt.Sprintf("%s (%s!)", dueDate.Format("2006-01-02"), locale.Date.Late)
	} else if until < time.Duration(24)*time.Hour {
		return locale.Date.Today
	} else if until < time.Duration(48)*time.Hour {
		return locale.Date.Tomorrow
	}

	return dueDate.Format("2006-01-02")
}

func formatGoal(locale *i18n.Locale, goal *Goal) string {
	messageText := fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(goal.Text))
	if !goal.DueAt.IsZero() {
		dueLine := fmt.Sprintf("%s: <b>%s</b>", locale.Messages.Due, dueString(locale, goal.DueAt))
		messageText = messageText + "\n" + dueLine
	}
	return messageText
}

func formatAction(action *Action) string {
	messageText := fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(action.Text))
	if action.ContextID != 0 {
		contextLabel := fmt.Sprintf(" (<b>@%s</b>)", escapeForHTMLFormatting(action.Context.Text))
		messageText = messageText + contextLabel
	}
	return messageText
}

func formatBold(messageText string) string {
	return fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(messageText))
}

func escapeForHTMLFormatting(msg string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(msg, "&",
				"&amp;", -1),
			"<", "&lt;", -1),
		">", "&gt;", -1)
}
