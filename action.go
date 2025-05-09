package tg_done_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Action struct {
	ID          int
	UserID      int `sql:",notnull"`
	GoalID      int `sql:",notnull"`
	ContextID   int
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	ReviewedAt  time.Time `sql:",notnull"`
	CompletedAt time.Time
	DroppedAt   time.Time

	User    *User
	Goal    *Goal
	Context *Context
}

func (a *Action) BeforeInsert(db orm.DB) error {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
		// Subtract 10 minutes so action will be available for suggestion immediately
		a.ReviewedAt = time.Now().Add(-10 * time.Minute)
	}
	return nil
}
