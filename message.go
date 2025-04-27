package tg_done_bot

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/leonid-shevtsov/tg-done-bot/i18n"
)

type message struct {
	interaction     *interaction
	locale          *i18n.Locale
	telegramMessage telegram.MessageConfig
}

func (i *interaction) reply() *message {
	telegramMessage := telegram.NewMessage(int64(i.state.userID()), "")
	telegramMessage.ReplyMarkup = telegram.ReplyKeyboardRemove{RemoveKeyboard: true}
	return &message{
		interaction:     i,
		locale:          i.locale,
		telegramMessage: telegramMessage,
	}
}

func (i *interaction) sendText(messageText string) {
	i.reply().text(messageText).send()
}

func (i *interaction) sendUnclear() {
	i.reply().text(i.locale.Messages.PickOneOfTheOptions).send()
}

func (m *message) text(messageText string) *message {
	m.telegramMessage.Text = messageText
	return m
}

func (m *message) html(messageHTML string) *message {
	m.telegramMessage.Text = messageHTML
	m.telegramMessage.ParseMode = telegram.ModeHTML
	return m
}

func (m *message) goal(goal *Goal) *message {
	return m.html(formatGoal(m.locale, goal))
}

func (m *message) action(action *Action) *message {
	return m.html(formatAction(action))
}

func (m *message) waitingFor(waitingFor *WaitingFor) *message {
	return m.html(formatWaitingFor(waitingFor))
}

func (m *message) inboxItem(inboxItem *InboxItem) *message {
	return m.html(formatInboxItem(inboxItem))
}

func (m *message) boldText(messageText string) *message {
	return m.html(formatBold(messageText))
}

func (m *message) acceptFreeInput() *message {
	m.telegramMessage.Text += " ✏️"
	return m
}

func (m *message) keyboard(keyboard [][]string) *message {
	m.telegramMessage.ReplyMarkup = telegram.ReplyKeyboardMarkup{
		Keyboard:        makeTelegramKeyboard(keyboard),
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	return m
}

func (m *message) send() {
	m.interaction.bot.Send(m.telegramMessage)
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
