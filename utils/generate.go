package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReferralCode(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//return random int in the range min...max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//return random char string from charset
func randomChar(cs []byte) string {
	return string(cs[randomInt(0, len(cs)-1)])
}

func GetYearMonth() uint64 {
	yearmonth, _ := strconv.Atoi(time.Now().Format("200601"))
	return uint64(yearmonth)
}
