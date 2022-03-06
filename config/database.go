package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	conn := "root:dev123@tcp(localhost:3306)/northwind"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		panic(err)
	}

	return db
}
