package gtd_bot

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Context struct {
	ID        int
	UserID    int       `sql:",notnull"`
	Text      string    `sql:",notnull"`
	CreatedAt time.Time `sql:",notnull"`

	Actions []*Action
}

func (c *Context) BeforeInsert(db orm.DB) error {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	return nil
}
