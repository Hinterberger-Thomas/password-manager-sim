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

type account struct {
	id       int64
	account  string
	password string
}

func Init_db() *DB {
	db, err := sql.Open("mysql", "mysql:mysql@/datadb")
	if err != nil {
		fmt.Println(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{
		client: db,
	}
}

func (db *DB) InsertAccount(accNam string, password string) (int64, error) {
	stmt, e := db.client.Prepare("insert into data(id, account, password) values (?, ?, ?)")
	if e != nil {
		return -1, e
	}
	rows, e := db.client.Query("SELECT * FROM data WHERE max(id)")
	if e != nil {
		return -1, e
	}
	rows.Next()
	var maxInt int64
	rows.Scan(&maxInt)

	//execute
	res, e := stmt.Exec(maxInt, accNam, password)

	if e != nil {
		return -1, e
	}

	id, e := res.LastInsertId()
	if e != nil {
		return -1, e
	}

	return id, nil
}

func (db *DB) RemoveAccount(accNam string) (int64, error) {
	stmt, e := db.client.Prepare("insert into data(id, account, password) values (?, ?, ?)")
	return 0, nil
}
