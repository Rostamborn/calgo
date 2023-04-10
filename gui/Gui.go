package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	// "fyne.io/fyne/v2/widget"
	"github.com/calgo/parser"
)

func Init(events []parser.Event) {
	app := app.New()
	mainWindow := app.NewWindow("CALGO")
	mainWindow.Resize(fyne.NewSize(700, 850))
	mainWindow.SetMaster()

	// weekdays := []string{"sat", "sun", "mon", "tue", "wed", "thu", "fri"}
	// topBar := container.NewHBox()
	// for _, v := range weekdays {
	// 	topBar.Add(widget.NewLabel(v))
	// }

	days := InitDays(events, mainWindow)
	months := InitMonths(days)

	// mainContainer := container.NewBorder(topBar, month, nil, nil)
	mainWindow.SetContent(months[0].con)

	mainWindow.Show()
	app.Run()
}
