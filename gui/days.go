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
	mainContainer := container.New(layout.NewGridLayout(6))
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