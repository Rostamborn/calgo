package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"github.com/calgo/parser"
)

type Date struct {
	day, month, year, weekday string
}

type Day struct {
	date Date
	con  *fyne.Container
	// weekday    string // sat sun mon tue wed thu fri
	events     []parser.Event
	mainWindow fyne.Window
}

// an Event object can be created here and be added to
// the global list of Events and written into ics file
// TODO: make it so that events newly added, are also
// written in the .ics file
func (day *Day) addEvent(event parser.Event) {
	if len(day.events) < 3 {
		if event == (parser.Event{}) {
			newEvent := parser.Event{}
			newText := &tappableLabel{day: day, event: newEvent}
			newText.SetText("new event")
			day.events = append(day.events, newEvent)
			day.con.Add(newText)
			day.con.Refresh()
		} else if day.date.day == getDate(event).day {
			newLabel := &tappableLabel{day: day, event: event}
			newLabel.SetText(event.Summary)
			day.events = append(day.events, event)
			day.con.Add(newLabel)
			day.con.Refresh()
		}
	} else {
		dialog := dialog.NewInformation("Error", "can't add more events", day.mainWindow)
		dialog.Show()
	}
}

func (day *Day) removeEvent(label *tappableLabel) {
	day.con.Remove(label)
	var indexToRemove int
	for i, v := range day.events {
		if label.event == v {
			indexToRemove = i
			break
		}
	}
	day.events = append(day.events[:indexToRemove], day.events[indexToRemove+1:]...)
	day.con.Refresh()
}

// a function to initialize all the days in a year
func InitDays(events []parser.Event, mainWindow fyne.Window) *fyne.Container {
	// these starting variables can be stored and read from a file
	// it does'nt matter which kind of calendar it is
	startingWeekday := "sun"
	startingYear := 2023
	calendarType := "Gregorian" // Solar, Gregorian

	weekdays := []string{"sat", "sun", "mon", "tue", "wed", "thu", "fri"}
	var days [31]*Day // for testing we stick to small numbers for now

	mainContainer := container.New(layout.NewGridLayout(7))
	for _, v := range weekdays {
		text := canvas.NewText(v, color.White)
		text.TextSize = 26
		mainContainer.Add(text)
	}

	for i, v := range weekdays {
		if v == startingWeekday {
			for j := i; j > 0; j-- {
				mainContainer.Add(container.New(layout.NewVBoxLayout()))
			}
		}
	}

	for i := 0; i < 31; i++ {
		days[i] = &Day{
			date:       calculateDate(calendarType, i+1, startingYear, startingWeekday),
			con:        container.New(layout.NewVBoxLayout()),
			mainWindow: mainWindow,
		}
		// newText := &tappableLabel{day: days[i]}
		// newText.SetText(days[i].date.day)
		newText := &titleLabel{day: days[i]}
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
func calculateDate(calendarType string, dayNumber int, year int, startingWeekday string) Date {
	daysInMonthGregorian := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	daysInMonthSolar := []int{31, 31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 29}
	var daysInMonth []int
	if calendarType == "Gregorian" {
		daysInMonth = daysInMonthGregorian
	} else if calendarType == "Solar" {
		daysInMonth = daysInMonthSolar
	}

	weekdays := []string{"sat", "sun", "mon", "tue", "wed", "thu", "fri"}

	/////////////////////
	tmp := (dayNumber - 1) % 7
	var index int
	for i, v := range weekdays {
		if v == startingWeekday {
			index = i
			break
		}
	}
	finalIndex := (index + tmp) % 7
	finalResultForWeekday := weekdays[finalIndex]

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
