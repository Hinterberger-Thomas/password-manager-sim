package sec

import (
	"math/rand"
	"time"
)

func GenPassword() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(27) + 1
}
