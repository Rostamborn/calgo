package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/calgo/parser"
)

// Title of each day
type titleLabel struct {
	widget.Label
	day *Day
}

func (t *titleLabel) Tapped(_ *fyne.PointEvent) {
	button := widget.NewButton("create", func() {
		t.day.addEvent(parser.Event{})
	})
	dialog := dialog.NewCustom("Add Event", "cancel", button, t.day.mainWindow)
	dialog.Show()

}

// labels which can be tapped. new pane is created for
// more options
type tappableLabel struct {
	widget.Label
	event parser.Event
	day   *Day
}

func (t *tappableLabel) Tapped(_ *fyne.PointEvent) {
	button := widget.NewButton("delete", func() {
		t.day.removeEvent(t)
	})
	dialog := dialog.NewCustom("Add Event", "ok", button, t.day.mainWindow)
	dialog.Show()
}
