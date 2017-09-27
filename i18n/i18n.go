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

type LocaleWhatIsTheDueDate struct {
	Prompt     string
	None       string
	Today      string
	Tomorrow   string
	EndOfWeek  string
	FormatHelp string
	Success    string
}

type LocaleWhatIsTheNextAction struct {
	Prompt     string
	WaitingFor string
}

type LocaleWhatAreYouWaitingFor struct {
	Prompt  string
	Nothing string
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
	ChangeNextAction      string
	ItIsDone              string
	Skipping              string
}

type LocaleDoing struct {
	Prompt    string
	Completed string
}

type LocaleMoveGoalForward struct {
	Prompt                  string
	GoalIsAchieved          string
	ReviewLater             string
	CongratulationsComplete string
	WillReviewLater         string
	AddedAction             string
}

type LocaleCheckWaitingFor struct {
	YourGoal         string
	IsWaitingFor     string
	ItIsReady        string
	StillWaiting     string
	Success          string
	ContinuingToWait string
}

type LocaleWhatIsTheGoalWaitingFor struct {
	Prompt  string
	Nothing string
}

type LocaleProcessing struct {
	TrashIt string
	Abort   string
	Aborted func(count int) string
}

type LocaleReviewGoal struct {
	LetsReviewThisGoal string
	Prompt             string
	Success            string
}

type LocaleReviewGoalStatement struct {
	Prompt string
}

type LocaleReviewGoalChangeStatement struct {
	Prompt string
}

type LocaleReviewGoalDueDate struct {
	Prompt       string
	PromptNoDate string
}

type LocaleReviewGoalChangeDueDate struct {
	Prompt  string
	Cleared string
}

type LocaleMessages struct {
	PickOneOfTheOptions string
	GoalTrashed         string
	ServerRestart       string
	Due                 string
}

type LocaleDate struct {
	Today    string
	Tomorrow string
	Late     string
}

type LocaleCommands struct {
	Yes         string
	No          string
	Done        string
	DoItLater   string
	TrashGoal   string
	BackToInbox string
	WaitingFor  string
	Keep        string
}

type Locale struct {
	CollectingInbox           LocaleCollectingInbox
	IsItActionable            LocaleIsItActionable
	WhatIsTheGoal             LocaleWhatIsTheGoal
	WhatIsTheDueDate          LocaleWhatIsTheDueDate
	WhatIsTheNextAction       LocaleWhatIsTheNextAction
	WhatAreYouWaitingFor      LocaleWhatAreYouWaitingFor
	CanYouDoItNow             LocaleCanYouDoItNow
	DoItNow                   LocaleDoItNow
	ActionSuggestion          LocaleActionSuggestion
	Doing                     LocaleDoing
	MoveGoalForward           LocaleMoveGoalForward
	CheckWaitingFor           LocaleCheckWaitingFor
	WhatIsTheGoalWaitingFor   LocaleWhatIsTheGoalWaitingFor
	Processing                LocaleProcessing
	ReviewGoal                LocaleReviewGoal
	ReviewGoalStatement       LocaleReviewGoalStatement
	ReviewGoalChangeStatement LocaleReviewGoalChangeStatement
	ReviewGoalDueDate         LocaleReviewGoalDueDate
	ReviewGoalChangeDueDate   LocaleReviewGoalChangeDueDate
	Messages                  LocaleMessages
	Date                      LocaleDate
	Commands                  LocaleCommands
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
		"Does this require action in the next one or two weeks?",
		"Trashed! Moving on.",
		"No - trash it",
	},
	LocaleWhatIsTheGoal{
		"What is the goal here?",
	},
	LocaleWhatIsTheDueDate{
		"What is the due date?",
		"None",
		"Today",
		"Tomorrow",
		"End of week",
		"Please enter date in YYYY-MM-DD format, or pick one of the options",
		"Due date is now: %s",
	},
	LocaleWhatIsTheNextAction{
		"What is the next physical action?",
		"I am blocked (waiting)",
	},
	LocaleWhatAreYouWaitingFor{
		"What are you waiting for?",
		"Actually, I'm unblocked.",
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
		"Let's set another next action.",
		"It is already done.",
		"OK, skipping for now...",
	},
	LocaleDoing{
		"Great! Waiting for you to finish.",
		"Awesome!",
	},
	LocaleMoveGoalForward{
		"What is the next action towards this goal?",
		"This goal is achieved",
		"Let's review it later",
		"Congratulations on succeeding!",
		"OK, marked goal for review.",
		"OK, recorded next action for this goal.",
	},
	LocaleCheckWaitingFor{
		"Your goal:",
		"is waiting for:",
		"It is ready.",
		"Still waiting.",
		"Awesome! Moving on.",
		"OK, waiting for now.",
	},
	LocaleWhatIsTheGoalWaitingFor{
		"What is the goal waiting for?",
		"Actually, it's not blocked.",
	},
	LocaleProcessing{
		"Let's just trash it",
		"Let's do this later",
		func(count int) string {
			return pluralize(count, "OK. %d inbox item left to process.", "OK. %d inbox items left to process.")
		},
	},
	LocaleReviewGoal{
		"Let's review this goal:",
		"Is this something you are going to work on in the upcoming week?",
		"Goal is reviewed and up-to-date.",
	},
	LocaleReviewGoalStatement{
		"Is the goal statement still relevant?",
	},
	LocaleReviewGoalChangeStatement{
		"What is the current goal statement?",
	},
	LocaleReviewGoalDueDate{
		"Is the current due date still in effect?",
		"Is this goal still without a due date?",
	},
	LocaleReviewGoalChangeDueDate{
		"What is the current due date?",
		"Due date is now cleared",
	},
	LocaleMessages{
		"Please pick one of the options.",
		"OK, goal trashed.",
		"Goooood morning! I've got restarted.",
		"due",
	},
	LocaleDate{
		"today",
		"tomorrow",
		"late",
	},
	LocaleCommands{
		"Yes",
		"No",
		"Done!",
		"I'll do it later",
		"Trash this goal.",
		"Done working for now.",
		"I am waiting for something",
		"Keep the current one",
	},
}
