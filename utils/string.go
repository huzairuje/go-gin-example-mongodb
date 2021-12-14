package utils

import (
	"math/rand"
	"time"
)

const (
	MaxLimitLoanProcessPerDay        = 50
	InterestRate                     = 0.01
	DefaultServerPort                = ":4010"
	DateFormat                       = "2006-01-02"
	CollectionName                   = "loans"
	OK                               = "OK"
	LoanProcessIsSuccess             = "Operation Loan Process Is Successfully Executed"
	NotMatchingAnyRoute              = "Not Matching of Any Routes"
	SomethingWentWrong               = "Oops, Something Went Wrong"
	BadRequest                       = "Bad Request"
	NotFound                         = "Not Found"
	ProvinceIsNotValid               = "Province Is Not Valid, the suggesting value is (DKI JAKARTA, JAWA BARAT, JAWA TIMUR OR SUMATERA UTARA)"
	GenderIsNotValid                 = "Gender Is Not Valid, the suggesting value is (L or P)"
	TenorIsNotValid                  = "Tenor Is Not Valid, the suggesting value is (3, 6, 9, 12 or 24)"
	AgeIsNotValid                    = "Age Is Not Valid, you must be at least 17 years old or not older than 80 years old"
	AmountLoanIsNotValid             = "Amount Loan Is Not Valid, the suggesting value is (1000000 - 10000000)"
	AgeIsNotValidFormat              = "Age Is Not Valid, the suggesting value format is (2006-01-02)"
	KTPNumberIsExist                 = "Can't Process With the KTP Number (Is Already being Used)"
	MaxLimitRequestLoanProcessPerDay = "Sorry, Cannot Process the Loan because Max Limit Request Reached Loan Process Per Day"
	FailedPingDb                     = "failed to ping database, you can't use this connection if you still ignore this error : %v"
	FailedOpenDb                     = "failed to open database connection, you can't use this connection if you still ignore this error : %v"
	DatabaseConfigNotSet             = "database config is not set, so we are using default config for mongodb connection"
	DatabaseConfigSet                = "database config is set,we are using your config for mongodb connection"
	DatabaseNameIsNotSet             = "database name in config is not set, we are using the default database name"
	ServerPortIsNotSet               = "Server Port is not set in config.yml, so we use default port: " + DefaultServerPort
	ServerPortIsSet                  = "Server Port is set in config.yml, so we use from config "
	Accepted                         = "Accepted"
	Rejected                         = "Rejected"
)

func RandomizeStatus() string {
	rand.Seed(time.Now().Unix())
	status := []string{
		Accepted,
		Rejected,
	}
	return status[rand.Intn(len(status))]
}

// AgeAt gets the age of an entity at a certain time.
func AgeAt(birthDate time.Time, now time.Time) int {
	// Get the year number change since the player's birth.
	years := now.Year() - birthDate.Year()

	// If the date is before the date of birth, then not that many years have elapsed.
	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years -= 1
	}

	return years
}

// Age is shorthand for AgeAt(birthDate, time.Now()), and carries the same usage and limitations.
func Age(birthDate time.Time) int {
	return AgeAt(birthDate, time.Now())
}

// Gets the adjusted date of birth to work around leap year differences.
func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

// Works out if a time.Time is in a leap year.
func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func GetDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
