package database

import (
	"database/sql"
	"log"
	"github.com/davidobando99/APIRestWithGo/database"
	"time"
)

connection,err := Connection()
func CreateTable() {

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS domain (host STRING PRIMARY KEY, sslGrade STRING, previousGrade STRING, lastSearch TIMESTAMPTZ)"); err != nil {
		log.Fatal(err)
	}

}
func InsertDomain(host string, sslgrade string, previous string) {

	if _, err := db.Exec(
		"INSERT INTO domain (host, sslGrade,previousGrade,lastSearch) VALUES (host, sslgrade,previous,NOW())"); err != nil {
		log.Fatal(err)
	}

}

