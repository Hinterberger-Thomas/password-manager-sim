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
				fmt.Println(err)
			}
			fmt.Println("pls enter key ")
			keyUn, err := reader.ReadString('\n')
			keyUnW := strings.Replace(keyUn, "\n", "", -1)
			key := []byte(keyUnW)
			acc.Password = sec.Decrypt(key, acc.Password)
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
			passwordUn, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("pls enter key ")
			keyUn, err := reader.ReadString('\n')
			keyUnW := strings.Replace(keyUn, "\n", "", -1)
			passwordUnW := strings.Replace(passwordUn, "\n", "", -1)
			key := []byte(keyUnW)
			cryptoText := sec.Encrypt(key, passwordUnW)

			if err != nil {
				log.Fatal(err)
			}

			id, err := db.InsertAccount(accNam, cryptoText)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id)
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
