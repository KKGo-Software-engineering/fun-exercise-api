package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

type Config struct {
	DatabaseURL string
}

func New(cfg Config) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	err = db.Ping()
	
	if err != nil {
		log.Fatal(err)
	}
	
	return &Postgres{Db: db}, nil
}
