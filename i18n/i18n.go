package i18n

import "fmt"

func pluralize(count int, one string, other string) string {
	if (count % 10) == 1 {
		return fmt.Sprintf(one, count)
	} else {
		return fmt.Sprintf(other, count)
	}
}

type LocaleCollectingInbox struct {
	Prompt       string
	StartWorking string
	NoMoreWork   string
	Added        func(count int) string
}

type LocaleIsItActionable struct {
	ProcessingInboxItem string
	Prompt              string
	Trashed             string
	NoTrashIt           string
}

type LocaleWhatIsTheGoal struct {
	Prompt string
}

type LocaleWhatIsTheNextAction struct {
	Prompt string
}

type LocaleCanYouDoItNow struct {
	Prompt string
}

type LocaleDoItNow struct {
	Prompt string
}

type LocaleActionSuggestion struct {
	IThinkYouShouldWorkOn string
	ByDoing               string
	Doing                 string
	Skip                  string
	ItIsDone              string
	BackToInbox           string
	Skipping              string
}

type LocaleDoing struct {
	Prompt    string
	Completed string
}

type LocaleMoveGoalForward struct {
	NowYouAreCloserToAchieving string
	Prompt                     string
	GoalIsAchieved             string
	ReviewLater                string
	CongratulationsComplete    string
	WillReviewLater            string
	AddedAction                string
}

type LocaleProcessing struct {
	TrashIt string
	Abort   string
	Aborted func(count int) string
}

type LocaleMessages struct {
	PickOneOfTheOptions string
}

type LocaleCommands struct {
	Yes       string
	No        string
	Done      string
	DoItLater string
}

type Locale struct {
	CollectingInbox     LocaleCollectingInbox
	IsItActionable      LocaleIsItActionable
	WhatIsTheGoal       LocaleWhatIsTheGoal
	WhatIsTheNextAction LocaleWhatIsTheNextAction
	CanYouDoItNow       LocaleCanYouDoItNow
	DoItNow             LocaleDoItNow
	ActionSuggestion    LocaleActionSuggestion
	Doing               LocaleDoing
	MoveGoalForward     LocaleMoveGoalForward
	Processing          LocaleProcessing
	Messages            LocaleMessages
	Commands            LocaleCommands
}

var En = Locale{
	LocaleCollectingInbox{
		"Collecting inbox.",
		"I'm ready for some work",
		"No more work!",
		func(count int) string {
			return pluralize(count, "Added to inbox. Now in inbox: %d item.", "Added to inbox. Now in inbox: %d items.")
		},
	},
	LocaleIsItActionable{
		"Processing inbox item:",
		"Is it actionable?",
		"Trashed! Moving on.",
		"No - trash it",
	},
	LocaleWhatIsTheGoal{
		"What is the goal here?",
	},
	LocaleWhatIsTheNextAction{
		"What is the next physical action?",
	},
	LocaleCanYouDoItNow{
		"Can you do it now in 2 minutes?",
	},
	LocaleDoItNow{
		"Do it! 2 minutes and counting.",
	},
	LocaleActionSuggestion{
		"I think you should work on:",
		"by doing:",
		"Yes, I'll do this.",
		"Skip this one for now.",
		"It is already done.",
		"Done working for now.",
		"OK, skipping for now...",
	},
	LocaleDoing{
		"Great! Waiting for you to finish.",
		"Awesome!",
	},
	LocaleMoveGoalForward{
		"Now you are closer to achieving:",
		"What is the next action towards this goal?",
		"This goal is achieved",
		"Let's review it later",
		"Congratulations on succeeding!",
		"OK, marked goal for review.",
		"OK, recorded next action for this goal.",
	},
	LocaleProcessing{
		"Let's just trash it",
		"Let's do this later",
		func(count int) string {
			return pluralize(count, "OK. %d inbox item left to process.", "OK. %d inbox items left to process.")
		},
	},
	LocaleMessages{
		"Please pick one of the options.",
	},
	LocaleCommands{
		"Yes",
		"No",
		"Done!",
		"I'll do it later",
	},
}
