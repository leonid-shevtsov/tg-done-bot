package gtd_bot

import "time"

type InboxItem struct {
	ID          int
	UserID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	ProcessedAt time.Time
}

type UserState struct {
	UserID             int       `sql:",pk"`
	LastMessageAt      time.Time `sql:",notnull"`
	State              int       `sql:",notnull"`
	CurrentInboxItemID int
	CurrentGoalID      int

	CurrentInboxItem *InboxItem
	CurrentGoal      *Goal
}

type Goal struct {
	ID          int
	UserID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	CompletedAt time.Time
	DroppedAt   time.Time

	Actions []*Action
}

type Action struct {
	ID          int
	UserID      int       `sql:",notnull"`
	GoalID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	CompletedAt time.Time

	Goal *Goal
}

type interactionState int

const (
	initialState = interactionState(iota)
	isItActionableState
	whatIsTheGoalState
	whatIsTheNextActionState
)
