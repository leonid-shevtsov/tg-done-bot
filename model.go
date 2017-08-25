package gtd_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	ID                 int `sql:",pk"`
	State              int `sql:",notnull"`
	CurrentInboxItemID int
	CurrentGoalID      int
	CurrentActionID    int
	CreatedAt          time.Time `sql:",notnull"`
	LastMessageAt      time.Time `sql:",notnull"`

	CurrentInboxItem *InboxItem
	CurrentGoal      *Goal
	CurrentAction    *Action
	InboxItems       []*InboxItem
	Goals            []*Goal
	Actions          []*Action
}

func (u *User) BeforeInsert(db orm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}

type InboxItem struct {
	ID          int
	UserID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	ProcessedAt time.Time

	User *User
}

func (i *InboxItem) BeforeInsert(db orm.DB) error {
	if i.CreatedAt.IsZero() {
		i.CreatedAt = time.Now()
	}
	return nil
}

type Goal struct {
	ID          int
	UserID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	CompletedAt time.Time
	DroppedAt   time.Time

	User    *User
	Actions []*Action
}

func (g *Goal) BeforeInsert(db orm.DB) error {
	if g.CreatedAt.IsZero() {
		g.CreatedAt = time.Now()
	}
	return nil
}

type Action struct {
	ID          int
	UserID      int       `sql:",notnull"`
	GoalID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	CompletedAt time.Time

	User *User
	Goal *Goal
}

func (a *Action) BeforeInsert(db orm.DB) error {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
	}
	return nil
}
