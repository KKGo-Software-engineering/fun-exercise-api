package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func New() (*Postgres, error) {
	// databaseSource := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable", "localhost", 5432, "root", "password", "wallet")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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
