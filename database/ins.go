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

func (db *DB) InsertAccount(accNam string, ) (int64, error) {
	stmt, e := db.client.Prepare("insert into data(id, account, password) values (?, ?, ?)")
	SELECT * FROM permlog WHERE max(id)
	if e != nil {
		return -1, e
	}

	//execute
	res, e := stmt.Exec("5", "Post five", "Contents of post 5")
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
	return 0, nil
}
