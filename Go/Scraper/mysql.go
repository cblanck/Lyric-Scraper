package main

import (
	"database/sql"
	"fmt"
	"os"
    _ "code.google.com/p/go-mysql-driver/mysql"
)

var db *sql.DB

func query(q string) error {
	_, err := db.Exec(q)
	return err
}

func search() *sql.Rows {
	query := "SELECT NUMBER,ARTIST,TITLE,ALBUM FROM CART WHERE SCHED_CODES IS NULL LIMIT 100"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf(os.Stderr, "SQL ERROR: %s", err.Error())
		networkErrors.Printf("SQL ERROR: %s", err.Error())
		return nil
	}
	return rows
}

func connect(user, pass, dbname string) error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s?charset=utf8", user, pass, dbname))
	return err
}
