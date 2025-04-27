package tg_done_bot

const questionOnboarding = ""

func init() {
	registerQuestion(questionOnboarding, nil, handleOnboarding)
}

func handleOnboarding(i *interaction) string {
	i.sendText(i.locale.Onboarding.Text)
	return questionCollectingInbox
}
