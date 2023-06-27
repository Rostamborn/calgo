package main

import (
    "fmt"
    "github.com/calgo/parser"
)

type Date struct {
	day, month, year, weekday string
}

// give the function a day in a year and it
// returns a Date object with correctly filled fields
func calculateDate(calendarType string, dayNumber int, year int, startingWeekday string) Date {
	var daysInMonth []int
	if calendarType == "Gregorian" {
		daysInMonth = GREGORIAN
	} else if calendarType == "Solar" {
		daysInMonth = SOLAR
	}

	// we modulu daynumber congruent to 7 and then based
	// on the starting weekday, we add the remainder to know
	// which weekday it is
	tmp := (dayNumber - 1) % 7
	var index int
	for i, v := range WEEKDAYS {
		if v == startingWeekday {
			index = i
			break
		}
	}
	finalIndex := (index + tmp) % 7
	finalResultForWeekday := WEEKDAYS[finalIndex]

	num := dayNumber
	var date Date
	for month, dayInMonth := range daysInMonth {
		num -= dayInMonth
		if num <= 0 {
			num += dayInMonth
			date.day = formatDayDate(num)
			date.month = formatMonthDate(month)
			date.year = fmt.Sprint(year)
			date.weekday = finalResultForWeekday
			break
		}
	}
	if date == (Date{}) {
		panic("couldn't calculate dates")
	}
	return date
}

// returns correct format for day in Date structure
func formatDayDate(dayInMonth int) string {
	var result string
	if dayInMonth < 10 {
		result = fmt.Sprintf("0%d", dayInMonth)
	} else {
		result = fmt.Sprint(dayInMonth)
	}
	return result
}

// returns correct format for month in Date structure
func formatMonthDate(month int) string {
	var result string
	if month < 10 {
		result = fmt.Sprintf("0%d", month+1)
	} else {
		result = fmt.Sprint(month + 1)
	}
	return result
}

// parsing the date string in .ics files
// ex: 20231127 -> year:2023 month:11 day:27
func getDate(event parser.Event) Date {
	startDate := event.DTStart
	runes := []rune(startDate)
	date := Date{
		day:   string(runes[6:]),
		month: string(runes[4:6]),
		year:  string(runes[:4]),
	}
	return date
}
