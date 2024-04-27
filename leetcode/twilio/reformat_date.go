package twilio

import (
	"strings"
)

// https://leetcode.com/problems/reformat-date/
func reformatDate(date string) string {
	date_arr := strings.Split(date, " ")

	d, m, y := date_arr[0], date_arr[1], date_arr[2]

	return getYearPart(y) + "-" + getMonthPart(m) + "-" + getDatePart(d)
}

func getDatePart(s string) string {
	i := 0
	for _, r := range s {
		if r == 's' || r == 'n' || r == 'r' || r == 't' {
			break
		}
		i++
	}

	s = s[:i]

	if len(s) == 1 {
		s = "0" + s
	}

	return s
}

func getMonthPart(s string) string {
	res := "1"
	switch strings.ToLower(s) {
	case "jan":
		res = "01"
	case "feb":
		res = "02"
	case "mar":
		res = "03"
	case "apr":
		res = "04"
	case "may":
		res = "05"
	case "jun":
		res = "06"
	case "jul":
		res = "07"
	case "aug":
		res = "08"
	case "sep":
		res = "09"
	case "oct":
		res = "10"
	case "nov":
		res = "11"
	case "dec":
		res = "12"
	}

	return res
}

func getYearPart(s string) string {
	return s
}
