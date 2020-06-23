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

func Connection() *sql.DB {

	db, err := sql.Open("postgres", "postgresql://davidobando99@localhost:26257/serversdb?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db

}
