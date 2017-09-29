package gtd_bot

import (
	"fmt"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (i *interaction) sendMessage(messageText string) {
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ReplyMarkup = telegram.ReplyKeyboardRemove{RemoveKeyboard: true}
	i.bot.Send(msg)
}

func (i *interaction) sendGoal(goal *Goal) {
	messageText := fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(goal.Text))
	if !goal.DueAt.IsZero() {
		dueLine := fmt.Sprintf("%s: <b>%s</b>", i.locale.Messages.Due, i.dueString(goal.DueAt))
		messageText = messageText + "\n" + dueLine
	}
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ParseMode = telegram.ModeHTML
	i.bot.Send(msg)
}

func (i *interaction) makeActionMessage(action *Action) *telegram.MessageConfig {
	messageText := fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(action.Text))
	if action.ContextID != 0 {
		contextLabel := fmt.Sprintf(" (<b>@%s</b>)", escapeForHTMLFormatting(action.Context.Text))
		messageText = messageText + contextLabel
	}
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ParseMode = telegram.ModeHTML
	return &msg
}

func (i *interaction) sendAction(action *Action) {
	i.bot.Send(i.makeActionMessage(action))
}

func (i *interaction) sendBoldMessage(messageText string) {
	i.bot.Send(i.makeBoldMessage(messageText))
}

func (i *interaction) makeBoldMessage(messageText string) *telegram.MessageConfig {
	boldText := fmt.Sprintf("<b>%s</b>", escapeForHTMLFormatting(messageText))
	msg := telegram.NewMessage(int64(i.state.userID()), boldText)
	msg.ParseMode = telegram.ModeHTML
	return &msg
}

func (i *interaction) sendHTMLMessage(messageText string) {
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	msg.ParseMode = telegram.ModeHTML
	i.bot.Send(&msg)
}

func escapeForHTMLFormatting(msg string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(msg, "&",
				"&amp;", -1),
			"<", "&lt;", -1),
		">", "&gt;", -1)
}

func (i *interaction) sendPrompt(messageText string, keyboard [][]string) {
	msg := telegram.NewMessage(int64(i.state.userID()), messageText)
	addKeyboardToMessage(&msg, keyboard)
	i.bot.Send(msg)
}

func (i *interaction) sendBoldPrompt(messageText string, keyboard [][]string) {
	msg := i.makeBoldMessage(messageText)
	addKeyboardToMessage(msg, keyboard)
	i.bot.Send(msg)
}

func (i *interaction) sendActionPrompt(action *Action, keyboard [][]string) {
	msg := i.makeActionMessage(action)
	addKeyboardToMessage(msg, keyboard)
	i.bot.Send(msg)
}

func addKeyboardToMessage(msg *telegram.MessageConfig, keyboard [][]string) {
	msg.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard:       makeTelegramKeyboard(keyboard),
		ResizeKeyboard: true,
	}
}

func makeTelegramKeyboard(buttons [][]string) [][]telegram.KeyboardButton {
	var telegramKeyboard [][]telegram.KeyboardButton
	for _, row := range buttons {
		var telegramRow []telegram.KeyboardButton
		for _, buttonText := range row {
			telegramRow = append(telegramRow, telegram.KeyboardButton{Text: buttonText})
		}
		telegramKeyboard = append(telegramKeyboard, telegramRow)
	}
	return telegramKeyboard
}

func (i *interaction) sendUnclear() {
	msg := telegram.NewMessage(int64(i.state.userID()), i.locale.Messages.PickOneOfTheOptions)
	i.bot.Send(msg)
}
