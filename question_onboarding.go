package gtd_bot

const questionOnboarding = ""

func init() {
	registerQuestion(questionOnboarding, nil, handleOnboarding)
}

func handleOnboarding(i *interaction) string {
	i.sendMessage("Hello! Onboarding message goes here!")
	return questionCollectingInbox
}
