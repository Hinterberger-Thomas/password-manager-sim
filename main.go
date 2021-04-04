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
	reader := bufio.NewReader(os.Stdin)
	database.Init_db()
	for true {

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
		switch num {
		case 1:
			log.Fatal("received value: " + "hey")
			break
		case 2:
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
