package gtd_bot

import (
	"log"
	"time"

	"github.com/go-pg/pg"
)

var db *pg.DB

func DBConnect() {
	db = pg.Connect(&pg.Options{
		Database: "gtd_bot",
		User:     "postgres",
		Addr:     ":5433",
	})
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})
	// err := createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&InboxItem{}, &UserState{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
