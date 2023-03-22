package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Init() {
  app := app.New()
  mainWindow := app.NewWindow("CALGO")
  mainWindow.SetMaster()
  mainWindow.Show()
  
  mainWindow.SetContent(widget.NewLabel("boom baam booommm"))

  app.Run()
}
