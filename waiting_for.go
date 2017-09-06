package gtd_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type WaitingFor struct {
	ID          int
	UserID      int       `sql:",notnull"`
	GoalID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	ReviewedAt  time.Time `sql:",notnull"`
	CompletedAt time.Time
	DroppedAt   time.Time

	User *User
	Goal *Goal
}

func (w *WaitingFor) BeforeInsert(db orm.DB) error {
	if w.CreatedAt.IsZero() {
		w.CreatedAt = time.Now()
		w.ReviewedAt = time.Now()
	}
	return nil
}
