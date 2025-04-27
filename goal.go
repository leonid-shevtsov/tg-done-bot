package tg_done_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Goal struct {
	ID          int
	UserID      int       `sql:",notnull"`
	Text        string    `sql:",notnull"`
	CreatedAt   time.Time `sql:",notnull"`
	CompletedAt time.Time
	DroppedAt   time.Time
	DueAt       time.Time
	ReviewedAt  time.Time

	User        *User
	Actions     []*Action
	WaitingFors []*WaitingFor
}

func (g *Goal) BeforeInsert(db orm.DB) error {
	if g.CreatedAt.IsZero() {
		g.CreatedAt = time.Now()
	}
	if g.ReviewedAt.IsZero() {
		g.ReviewedAt = time.Now()
	}
	return nil
}
