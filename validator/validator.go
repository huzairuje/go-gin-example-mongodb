package validator

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-gin-example-mongodb/utils"
)

func IsValidProvince(province string) bool {
	province = strings.ToUpper(province)
	switch province {
	case "DKI JAKARTA", "JAWA BARAT", "JAWA TIMUR", "SUMATERA UTARA":
		return true
	}
	return false
}

func IsValidTenor(tenor string) bool {
	switch tenor {
	case "3", "6", "9", "12", "24":
		return true
	}
	return false
}

func IsValidGender(gender string) bool {
	gender = strings.ToUpper(gender)
	switch gender {
	case "L", "P":
		return true
	}
	return false
}

func IsValidAge(t time.Time) bool {
	dob := utils.GetDOB(t.Year(), int(t.Month()), t.Day())
	age := utils.Age(dob)
	if age >= 17 && age <= 80 {
		return true
	}
	return false
}

func IsValidLoanAmount(amount string) bool {
	amountInt, _ := strconv.Atoi(amount)
	if amountInt >= 1000000 && amountInt <= 10000000 {
		return true
	}
	return false
}
