package database

import (
	"container/list"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	client *sql.DB
}

type Account struct {
	Id       int64
	Account  string
	Password string
}

func Init_db( /*password string*/ ) *DB {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/datadb")
	if err != nil {
		fmt.Println(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{
		client: db,
	}
}

func (db *DB) InsertAccount(accNam string, password string) (int64, error) {
	stmt, err := db.client.Prepare("INSERT INTO data(id, Account, password) VALUES (?, ?, ?);")
	if err != nil {
		return -1, err
	}
	rows, err := db.client.Query("SELECT max(id) FROM data;")
	if err != nil {
		return -1, err
	}
	rows.Next()
	var maxInt int64
	rows.Scan(&maxInt)

	res, err := stmt.Exec(maxInt+1, accNam, password)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *DB) GetAccount() (*list.List, error) {
	res, e := db.client.Query("SELECT id, Account, password FROM data;")
	var acc Account
	if e != nil {
		err := db.createTable()
		if err != nil {
			return nil, err
		}
		fmt.Println(e)
		fmt.Println("we just created a new one just for you my little friend")
		res, e = db.client.Query("SELECT id, Account, password FROM data;")
		if e != nil {
			fmt.Println(e)
		}

	}
	listOfAcc := list.New()
	for res.Next() {
		res.Scan(&acc.Id, &acc.Account, &acc.Password)
		listOfAcc.PushBack(acc)
	}
	if e != nil {
		return nil, e
	}
	return listOfAcc, nil
}

func (db *DB) createTable() error {
	const sqlCreateTable = "CREATE TABLE data (id int,Account varchar(255),password varchar(255));"
	const sqlInsIntTab = "INSERT INTO data(id, Account, password) VALUES (0, start, 1234);"

	_, err := db.client.Exec(sqlCreateTable)
	if err != nil {
		return err
	}
	_, err = db.client.Query(sqlInsIntTab)
	return err
}
