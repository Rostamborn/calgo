package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/calgo/parser"
)

type Date struct {
	day, month, year string
}

type tappableLabel struct {
	widget.Label
	day *Day
}

func (t *tappableLabel) Tapped(_ *fyne.PointEvent) {
	t.day.addEvent(parser.Event{})
}

type Day struct {
	date    Date
	con     *fyne.Container
	weekday string // sat sun mon tue wed thu fri
}

// an Event object can be created here and be added to
// the global list of Events and written into ics file
// TODO: make it so that events newly added, are also
// written in the .ics file
func (day *Day) addEvent(event parser.Event) {
	if event == (parser.Event{}) {
		newEvent := &tappableLabel{day: day}
		newEvent.SetText("new event")
		day.con.Add(newEvent)
		day.con.Refresh()
	} else if day.date.day == getDate(event).day {
		newLabel := &tappableLabel{day: day}
		newLabel.SetText(event.Summary)
		day.con.Add(newLabel)
		day.con.Refresh()
	}
}

// this function is for test purposes only
// constructing []Day array can be done in various ways
// Also there must be a function that adds the events present
// in the ics files into the []Day objects. (through the global Events list)
func InitMonth(events []parser.Event) *fyne.Container {
	var days [30]*Day
	mainContainer := container.New(layout.NewGridLayout(7))
	for i := 1; i < 31; i++ {
		if i < 10 {
			days[i-1] = &Day{
				date: Date{day: fmt.Sprintf("0%d", i)},
				con:  container.New(layout.NewVBoxLayout()),
			}
		} else {
			days[i-1] = &Day{
				date: Date{day: fmt.Sprint(i)},
				con:  container.New(layout.NewVBoxLayout()),
			}
		}
		newText := &tappableLabel{day: days[i-1]}
		newText.SetText(fmt.Sprint(i))
		days[i-1].con.Add(newText)
		mainContainer.Add(days[i-1].con)

	}

	for _, event := range events {
		for _, day := range days {
			day.addEvent(event)
		}

	}

	return mainContainer
}

// a function to initialize all the days in a year
func InitDays(events []parser.Event) *fyne.Container {
	weekdays := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	var days [31]*Day // for testing we stick to small numbers for now

	// there must be a grid container like this for each month which will
	// contain containers for each day. the container for each day is of type
	// verticalContainer or VBox
	// TODO: make it so that each month has a grid container and includes
	// relavent days
	mainContainer := container.New(layout.NewGridLayout(7))

	for i := 0; i < 31; i++ {
		days[i] = &Day{
			date:    calculateDate(i + 1),
			con:     container.New(layout.NewVBoxLayout()),
			weekday: weekdays[i%7],
		}
		newText := &tappableLabel{day: days[i]}
		newText.SetText(days[i].date.day)
		days[i].con.Add(newText)
		mainContainer.Add(days[i].con)
	}
	for _, event := range events {
		for _, day := range days {
			day.addEvent(event)
		}
	}
	return mainContainer
}

// give the function a day in a year and it
// returns a Date object with correctly filled fields
func calculateDate(dayNumber int) Date {
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	num := dayNumber
	var date Date
	for month, dayInMonth := range daysInMonth {
		num -= dayInMonth
		if num <= 0 {
			num += dayInMonth
			date.day = formatDayDate(num)
			date.month = formatMonthDate(month)
			date.year = "2023"
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
	// switch month {
	// case 0:
	// 	result = "01"
	// case 1:
	// 	result = "02"
	// case 2:
	// 	result = "03"
	// case 3:
	// 	result = "04"
	// case 4:
	// 	result = "05"
	// case 5:
	// 	result = "06"
	// case 6:
	// 	result = "07"
	// case 7:
	// 	result = "08"
	// case 8:
	// 	result = "09"
	// case 9:
	// 	result = "10"
	// case 10:
	// 	result = "11"
	// case 11:
	// 	result = "12"
	// }
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
