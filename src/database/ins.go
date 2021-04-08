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

type Account struct {
	Id       int64
	Account  string
	Password string
}

func InitDB( /*password string*/ ) *DB {
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

func (db *DB) InsertAccount(accNam string, file string) (int64, error) {
	stmt, err := db.client.Prepare("INSERT INTO data(id, Account, file) VALUES (?, ?, ?);")
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

	res, err := stmt.Exec(maxInt+1, accNam, file)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *DB) GetAllAccounts() (*sql.Rows, error) {
	res, e := db.client.Query("SELECT id, Account, file FROM data;")
	if e != nil {
		err := db.createTable()
		if err != nil {
			return nil, err
		}
		fmt.Println(e)
		fmt.Println("we just created a new one just for you my little friend")
		res, e = db.client.Query("SELECT id, Account, file FROM data;")
		if e != nil {
			fmt.Println(e)
		}

	}
	return res, nil

}

func (db *DB) createTable() error {
	const sqlCreateTable = "CREATE TABLE data (id int PRIMARY KEY,Account varchar(255),file blob);"
	const sqlInsIntTab = "INSERT INTO data(id, Account, file) VALUES (0, start, utl_raw.cast_to_raw('This is a blob description'));"

	_, err := db.client.Exec(sqlCreateTable)
	if err != nil {
		return err
	}
	_, err = db.client.Query(sqlInsIntTab)
	return err
}

func (db *DB) GetAccountFile(id int64) (string, error) {
	stmt, err := db.client.Query("SELECT file FROM data WHERE id = ?;", id)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	var file string
	stmt.Next()
	stmt.Scan(&file)
	return file, nil
}

func (db *DB) UpdDBFile(id int64, JSON []byte) error {
	_, err := db.client.Query("UPDATE data SET file = ? WHERE id = ?", id, JSON)
	if err != nil {
		return err
	}
	return nil
}
