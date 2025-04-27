package tg_done_bot

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func DBConnect() *pg.DB {
	dbURL := os.Getenv("DATABASE_URL")

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		panic(err)
	}
	// m.Force(1)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

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

	return db
}
