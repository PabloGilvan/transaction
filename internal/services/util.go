package services

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	MAX = 9
	MIN = 1
)

func GenerateAccountNumber(systemTrace string, today time.Time) string {
	dayOfTheYear := strconv.Itoa(time.Time.YearDay(today))
	if len(dayOfTheYear) < 3 {
		for i := 0; i <= (3 - len(dayOfTheYear)); i++ {
			dayOfTheYear = "0" + dayOfTheYear
		}
	}
	year := strconv.Itoa(time.Time.Year(today))
	hour := strconv.Itoa(time.Time.Hour(today))
	if len(hour) < 2 {
		hour = "0" + hour
	}

	sec := strconv.Itoa(time.Time.Second(today))
	if len(sec) < 2 {
		sec = "0" + sec
	}

	return year + dayOfTheYear + hour + sec + systemTrace
}

func GenerateTraceNumber() string {
	return strconv.Itoa(rand.Intn(MAX-MIN)+MIN) +
		strconv.Itoa(rand.Intn(MAX-MIN)+MIN) +
		strconv.Itoa(rand.Intn(MAX-MIN)+MIN) +
		strconv.Itoa(rand.Intn(MAX-MIN)+MIN) +
		strconv.Itoa(rand.Intn(MAX-MIN)+MIN)
}

func Nvl(value string, fallback string) string {
	if len(value) == 0 {
		return fallback
	}
	return value
}
