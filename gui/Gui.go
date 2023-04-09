package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/calgo/parser"
)

func Init(events []parser.Event) {
	app := app.New()
	mainWindow := app.NewWindow("CALGO")
	mainWindow.Resize(fyne.NewSize(700, 850))
	mainWindow.SetMaster()

	month := InitDays(events, mainWindow)
	mainWindow.SetContent(month)

	mainWindow.Show()
	app.Run()
}
