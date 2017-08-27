package gtd_bot

func (i *interaction) handleOnboarding() {
	i.sendMessage("Hello! Onboarding message goes here!")
	i.gotoInitialState()
}
