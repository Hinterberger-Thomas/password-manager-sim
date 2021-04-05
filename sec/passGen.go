package sec

import (
	"math/rand"
	"strconv"
	"time"
)

func GenPassword(passlen uint8) string {
	rand.Seed(time.Now().UnixNano())
	var password string
	for i := 0; uint8(i) < passlen; i++ {
		password += getPassData(rand.Intn(36))
	}
	return password
}

func getPassData(num int) string {
	pass := ""
	if num <= 10 {
		pass = strconv.Itoa(num)
	}
	switch num {
	case 11:
		pass = "A"
	case 12:
		pass = "B"
	case 13:
		pass = "C"
	case 14:
		pass = "D"
	case 15:
		pass = "E"
	case 16:
		pass = "F"
	case 17:
		pass = "G"
	case 18:
		pass = "H"
	case 19:
		pass = "I"
	case 20:
		pass = "J"
	case 21:
		pass = "K"
	case 22:
		pass = "L"
	case 23:
		pass = "M"
	case 24:
		pass = "N"
	case 25:
		pass = "O"
	case 26:
		pass = "P"
	case 27:
		pass = "Q"
	case 28:
		pass = "R"
	case 29:
		pass = "S"
	case 30:
		pass = "T"
	case 31:
		pass = "U"
	case 32:
		pass = "V"
	case 33:
		pass = "W"
	case 34:
		pass = "X"
	case 35:
		pass = "Y"
	case 36:
		pass = "Z"
	}
	return pass
}
