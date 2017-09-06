package gtd_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	ID                  int `sql:",pk"`
	ActiveQuestion      string
	CurrentInboxItemID  int
	CurrentGoalID       int
	CurrentActionID     int
	CurrentWaitingForID int
	CreatedAt           time.Time `sql:",notnull"`
	LastMessageAt       time.Time `sql:",notnull"`

	CurrentInboxItem  *InboxItem
	CurrentGoal       *Goal
	CurrentAction     *Action
	CurrentWaitingFor *WaitingFor
	InboxItems        []*InboxItem
	Goals             []*Goal
	Actions           []*Action
}

func (u *User) BeforeInsert(db orm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}
