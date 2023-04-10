package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

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

	// var tabs *container.AppTabs
	tabs := container.NewAppTabs()
	for i := 0; i < 12; i++ {
		tabs.Append(container.NewTabItem(fmt.Sprint(i+1), months[i].con))
	}
	// tabs.Append(container.NewTabItem("january", months[0].con))
	// tabs.Append(container.NewTabItem("february", months[1].con))

	// mainContainer := container.NewBorder(topBar, month, nil, nil)
	// mainWindow.SetContent(months[0].con)
	mainWindow.SetContent(tabs)

	mainWindow.Show()
	app.Run()
}
