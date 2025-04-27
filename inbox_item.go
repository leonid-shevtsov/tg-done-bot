package tg_done_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

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
