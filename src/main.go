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

func main() {
	db := database.Init_db()
	/*var db *database.DB
	for true {
		fmt.Println("Enter password for your account pls")

		reader := bufio.NewReader(os.Stdin)
		pass, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		pass = strings.Replace(pass, "\n", "", -1)

		_, err = db.GetAccount(0)
		if err == nil {
			break
		}
	}*/
	for true {

		fmt.Println("1.) Get Account")
		fmt.Println("2.) Insert Account")
		fmt.Println("3.) Generate Password")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		switch num {
		case 1:
			listAcc, err := db.GetAllAccount()
			if err != nil {
				fmt.Println(err)
			}
			for e := listAcc.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
			break
		case 2:
			fmt.Println("pls enter id name")
			id, err := reader.ReadString('\n')
			id = strings.Replace(text, "\n", "", -1)
			if err != nil {
				fmt.Println(err)
			}
			idNum, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Println("no number")
				return
			}
			if a, err := db.GetAccountFile(idNum); a != "" {
				if err != nil {
					fmt.Println(err)
				}

				return
			}
			fmt.Println("pls enter account name ")
			text, err = reader.ReadString('\n')
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
			break
		case 3:
			fmt.Println("password length ")
			text, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			text := strings.Replace(text, "\n", "", -1)
			num, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			a := sec.GenPassword(uint8(num))
			println(a)
			break
		default:
			return
		}
	}
}
