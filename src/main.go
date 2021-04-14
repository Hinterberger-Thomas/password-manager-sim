package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Hinterberger-Thomas/password-manager-sim/src/database"
	"github.com/Hinterberger-Thomas/password-manager-sim/src/sec"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var db *database.DB = database.InitDB()

func main() {
	for true {
		num := mainMenu()
		switch num {
		case 1:
			getAccCase()
			break
		case 2:
			getAccCase()
			break
		case 3:
			genPassCase()
			break
		default:
			fmt.Println("see you later alligator")
			return
		}
	}
}

func mainMenu() int64 {
	fmt.Println("1.) Get Account")
	fmt.Println("2.) Insert Account")
	fmt.Println("3.) Generate Password")

	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}
	num, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func getAccCase() {
	res, err := db.GetAllAccounts()

	if err != nil {
		fmt.Println(err)
	}
	var acc database.Account
	var listOfAcc []database.Account
	for res.Next() {
		res.Scan(&acc.Id, &acc.Account, &acc.Password)
		listOfAcc = append(listOfAcc, acc)
	}
	fmt.Println(listOfAcc)

}

func insAccCase() {
	fmt.Println("pls enter id name")
	id, err := reader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	if err != nil {
		fmt.Println(err)
	}
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("no number")
		return
	}
	if res, err := db.GetAccountFile(idNum); res != "" {
		if err != nil {
			fmt.Println(err)
		}

		insAccIntoDB(database.JSONToList(res))
	}
	var empByarr []database.Account
	insAccIntoDB(empByarr)

}

func insAccIntoDB(acc []database.Account) {
	fmt.Println("pls enter account name ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	accNam := strings.Replace(text, "\n", "", -1)
	fmt.Println("pls enter account password ")
	passwordUn, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	idInserted, err := db.InsertAccount(accNam, passwordUn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(idInserted)
}

func genPassCase() {
	fmt.Println("password length ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text = strings.Replace(text, "\n", "", -1)
	num, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	a := sec.GenPassword(uint8(num))
	fmt.Println(a)
}
