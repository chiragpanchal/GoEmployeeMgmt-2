package util

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "welcome"
	dbname   = "hrapps"
)

type RepoDb struct {
	DbClient *sqlx.DB
}

func DbConnection() *sqlx.DB {

	fmt.Println("entering DbConnection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		panic(err)
		return nil
	}

	fmt.Println("leaving DbConnection")

	return db
}
