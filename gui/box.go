package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Box struct {
	X         int
	Y         int
	Dimension int
	Image     *ebiten.Image
	Dialog    *DialogBox
	Title     *Label
	Labels    []*Label
}

func NewBox(x, y, dimension int, boximage, dialogimage *ebiten.Image, title *Label, labels []*Label) *Box {

	return &Box{
		X:         x,
		Y:         y,
		Dimension: dimension,
		Image:     boximage,
		Dialog:    NewDialogBox(x, y, 2*dimension, dialogimage),
		Title:     title,
		Labels:    labels,
	}
}

func (b *Box) AddLabel(label *Label) {
	b.Labels = append(b.Labels, label)
}

func (b *Box) SetOptions() *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.X), float64(b.Y))
	return options
}

func (b *Box) Draw(screen *ebiten.Image) {
	b.Title.Draw(b.Image)
	for _, label := range b.Labels {
		label.Draw(b.Image)
	}
	screen.DrawImage(b.Image, b.SetOptions())
}

func (b *Box) Clicked() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX >= b.X && mouseX <= b.X+b.Dimension && mouseY >= b.Y && mouseY <= b.Y+b.Dimension {
			b.Dialog.Visible = true
			return true
		}
	}
	return false
}

type DialogBox struct {
	X         int
	Y         int
	Dimension int
	Image     *ebiten.Image
	Visible   bool
	Buttons   map[string]*Button
}

func NewDialogBox(x, y, dimension int, image *ebiten.Image) *DialogBox {
	width := dimension / 3
	height := dimension / 8
	closeButton := NewButton(x, y, dimension-(width+10), dimension-(height+10), width, height, SlateGray, "CLOSE", "close")

	buttons := make(map[string]*Button)
    buttons["close"] = closeButton
	return &DialogBox{
		X:         x,
		Y:         y,
		Dimension: dimension,
		Image:     image,
		Visible:   false,
		Buttons:   buttons,
	}
}

func (d *DialogBox) SetOptions() *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(d.X), float64(d.Y))
	return options
}

func (d *DialogBox) Draw(screen *ebiten.Image) {
	if d.Visible {
		for _, button := range d.Buttons {
			button.Draw(d.Image)

		}
		screen.DrawImage(d.Image, d.SetOptions())
	}
}

func (d *DialogBox) ClickToExit() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX <= d.X || mouseX >= d.X+d.Dimension || mouseY <= d.Y || mouseY >= d.Y+d.Dimension {
			d.Visible = false
			return true
		}
        closeButton, ok := d.Buttons["close"]
        if ok && closeButton.Clicked() {
            d.Visible = false
            return true
        }
	}
	return false
}

func (d *DialogBox) Update() {
	for _, button := range d.Buttons {
		button.Update()
	}
}
