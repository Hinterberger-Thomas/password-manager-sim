package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Hinterberger-Thomas/password-manager-sim/database"
	"github.com/Hinterberger-Thomas/password-manager-sim/sec"
)

func main() {

	db := database.Init_db()
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
			fmt.Println("pls enter id")
			text, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			text = strings.Replace(text, "\n", "", -1)
			num, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			acc, err := db.GetAccount(num)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(acc)
			break
		case 2:
			fmt.Println("pls enter account name ")
			text, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			accNam := strings.Replace(text, "\n", "", -1)
			fmt.Println("pls enter account password ")
			text, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			password := strings.Replace(text, "\n", "", -1)
			id, err := db.InsertAccount(accNam, password)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id)
			break
		case 3:
			a := sec.GenPassword()
			println(a)
			break
		default:
			return
		}
	}
}
