package utils

import (
	"log"
	"testing"
)

func TestGenerateReferral(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := ReferralCode(5)
		log.Println(res)
	}
}

func TestGetYearMonth(t *testing.T) {
	s := GetYearMonth()
	log.Println(s)
}
