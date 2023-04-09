package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"github.com/calgo/parser"
)

type Date struct {
	day, month, year string
}

type Day struct {
	date       Date
	con        *fyne.Container
	weekday    string // sat sun mon tue wed thu fri
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
			date:       calculateDate(i + 1),
			con:        container.New(layout.NewVBoxLayout()),
			weekday:    weekdays[i%7],
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
