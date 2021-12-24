package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDatabase() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+`:`+os.Getenv("DB_PASSWORD")+`@/`+os.Getenv("DB_DATABASE"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected To Database")
	return db
}
