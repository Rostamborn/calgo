package gui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Month struct {
	number int
	year   int
	days   []*Day
	con    *fyne.Container
}

func InitMonths(days []*Day) []*Month {
	calendarType := "Solar" // Solar, Gregorian
	startingWeekday := days[0].date.weekday
	startingYear, _ := strconv.Atoi(days[0].date.year)

	var daysInMonth []int
	if calendarType == "Gregorian" {
		daysInMonth = GREGORIAN
	} else if calendarType == "Solar" {
		daysInMonth = SOLAR
	}

	months := make([]*Month, 12)
	prevMonthdays := 0
	for i, v := range daysInMonth {
		months[i] = &Month{
			number: i + 1,
			year:   startingYear,
			days:   days[prevMonthdays : prevMonthdays+v],
		}
		prevMonthdays += v
		monthContainer := container.New(layout.NewGridLayout(7))
		// adding weekdays
		for _, v := range WEEKDAYS {
			text := canvas.NewText(v, color.White)
			text.TextSize = TextSizeForWeekday
			monthContainer.Add(text)
		}
		// adding empty days
		for i, v := range WEEKDAYS {
			if v == startingWeekday {
				for j := i; j > 0; j-- {
					monthContainer.Add(container.New(layout.NewVBoxLayout()))
				}
			}
		}
		for _, v := range months[i].days {
			monthContainer.Add(v.con)
		}
		months[i].con = monthContainer
	}
	return months
}
