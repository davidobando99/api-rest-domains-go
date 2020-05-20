package database

import (
	"database/sql"
	"log"
	"time"
)

type DomainDB struct {
	Host        string
	SslGrade    string
	PreviousSSL string
	LastTime    time.Time
}

func Connection() (*sql.DB, error) {

	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db, err

}
