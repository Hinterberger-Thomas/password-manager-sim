package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	client *sql.DB
}

func Init_db() {
	db, err := sql.Open("mysql", "mysql:mysql@/datadb")
	if err != nil {
		fmt.Println(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func (db *DB) InsertAccount(accNam string) (bool, error) {
	return false, nil
}

func (db *DB) RemoveAccount(accNam string) (bool, error) {
	return false, nil
}
