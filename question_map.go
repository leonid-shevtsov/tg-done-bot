package gtd_bot

type question struct {
	ask          func(i *interaction)
	handleAnswer func(i *interaction) string
}

var questionMap = make(map[string]*question)

func registerQuestion(key string, ask func(i *interaction), handleAnswer func(i *interaction) string) {
	questionMap[key] = &question{ask, handleAnswer}
}
