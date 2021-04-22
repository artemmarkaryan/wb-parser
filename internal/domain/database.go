package domain

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DBconfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	ssl      bool
}

func NewDB() (db *sql.DB, err error) {
	dbConfig, err := newDBconfig()
	if err != nil {
		return
	}

	db = dbConfig.Connect()
	return
}

func newDBconfig() (config DBconfig, err error) {
	var (
		host     = "ec2-54-220-35-19.eu-west-1.compute.amazonaws.com"
		user     = "fjyyqaeamszurx"
		port     = "5432"
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = "d8j7e0r0rb8ggc"
	)
	if password == "" {
		err = errors.New("database config err: no password set")
	}
	config = DBconfig{
		host,
		port,
		user,
		password,
		dbname,
		true,
	}
	return
}

func (db *DBconfig) Connect() *sql.DB {
	database, err := sql.Open("postgres", db.buildConfigString())
	if err != nil {
		log.Fatal("cant connect to database:", err)
	}
	return database
}

func (db *DBconfig) buildConfigString() string {
	var sslMode string
	switch db.ssl {
	case true:
		sslMode = "require"
	default:
		sslMode = "disabled"
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.host, db.port, db.user, db.password, db.dbname, sslMode,
	)
}
