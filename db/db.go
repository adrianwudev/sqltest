package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/adrianwudev/sqltest/config"
)

func New() *sql.DB {
	dbConfig := config.GetDbConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}

	fmt.Println("DB connection established.")

	return db
}
