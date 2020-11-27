package database

import (
	"database/sql"
	"fiber-go-crud/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB instance
var db *sql.DB

// DB setting
const (
	host = config.DBHOST
	port = config.DBPORT
	user = config.DBUSER
	pass = config.DBPASS
	dbnm = config.DBNAME
)

// DB instance for export
var PgCon = ConnectionDB()

// DB Connection
func ConnectionDB() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbnm))

	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

// DB Connect Postgres Instance
func ConnectPgInit() *sql.DB {
	client, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbnm))
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	return client
}
