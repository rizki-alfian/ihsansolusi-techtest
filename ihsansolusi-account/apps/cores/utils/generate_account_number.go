package utils

import (
	"fmt"
	// "math/rand"
	"strconv"
	"time"
)

func HalfAccountNumber(code_area ...string) string {
	// rand.Seed(time.Now().UnixNano())

	prefix := "62"
	if len(code_area) > 0 && code_area[0] != "" {
		prefix = code_area[0]
	}

	now := time.Now()
	yearPart := strconv.Itoa(now.Year())[2:]
	monthPart := fmt.Sprintf("%02d", now.Month())
	datePart := yearPart + monthPart

	return prefix + datePart
}

func GenerateAccountNumber(code_area string, last_prefix int) string {
	
	half_account_number := HalfAccountNumber(code_area)

	seq_number := fmt.Sprintf("%06d", last_prefix)
	return half_account_number + seq_number 
}