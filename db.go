package gtd_bot

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

func DBConnect() *pg.DB {
	dbURL := os.Getenv("DATABASE_URL")
	dbOptions, err := pg.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(dbOptions)
	if os.Getenv("DEBUG") != "" {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}
			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}
	// err := createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }
	return db
}

// func createSchema(db *pg.DB) error {
// 	for _, model := range []interface{}{&InboxItem{}, &User{}, &Goal{}, &Action{}} {
// 		err := db.CreateTable(model, nil)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
