package gui

import (
	"image/color"
	"strings"

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
    b.Image.Fill(SlateGray)
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
	Buttons   map[string]*Button // close
    TextBoxes map[string]*TextBox // summary
}

func NewDialogBox(x, y, dimension int, image *ebiten.Image) *DialogBox {
	width := dimension / 3
	height := dimension / 8
	closeButton := NewButton(x, y, dimension-(width+10), dimension-(height+10), width, height, SlateGray, "CLOSE", "close")
    textBox := NewTextBox(x, y, 10, 15, 50, 20, 18)

    textBoxes := make(map[string]*TextBox)
    textBoxes["summary"] = textBox

	buttons := make(map[string]*Button)
    buttons["close"] = closeButton
	return &DialogBox{
		X:         x,
		Y:         y,
		Dimension: dimension,
		Image:     image,
		Visible:   false,
		Buttons:   buttons,
        TextBoxes: textBoxes,
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
        for _, textBox := range d.TextBoxes {
            textBox.Draw(d.Image)
        }
		screen.DrawImage(d.Image, d.SetOptions())
        d.Image.Fill(Teal)
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
    for _, textBox := range d.TextBoxes {
        textBox.Update()
    }
}


type TextBox struct {
    X     int
    Y     int
    RelX  int
    RelY  int
    Width int
    Height int
    Size  float64
    Label *Label
    Counter int
}

func NewTextBox(x, y, relX, relY, width, height int, size float64) *TextBox {
    label := NewLabel(relX, relY, "", color.Black, size)
    return &TextBox{
        X:         x,
        Y:         y,
        RelX:      relX,
        RelY:      relY,
        Width:     width,
        Height:    height,
        Size:      size,
        Label:    label,
        Counter: 0,
    }
}

func (t *TextBox) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(t.X), float64(t.Y))
    return options
}

func (t *TextBox) Draw(screen *ebiten.Image) {
    t.Label.Draw(screen)
}

func (t *TextBox) Clicked() bool {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        mouseX, mouseY := ebiten.CursorPosition()
        if mouseX >= t.X && mouseX <= t.X+t.Width && mouseY >= t.Y && mouseY <= t.Y+t.Height {
            return true
        }
    }
    return false
}

func (t *TextBox) Update() {
    t.Label.Runes = ebiten.AppendInputChars(t.Label.Runes[:0])
	t.Label.Text += string(t.Label.Runes)

	ss := strings.Split(t.Label.Text, "\n")
	if len(ss) > 3 {
		// t.Label.Text = strings.Join(ss[len(ss)-10:], "\n")
        t.Label.Text = ""
	}

	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) || len(t.Label.Text)%26 == 0 {
		t.Label.Text += "\n"
	}

	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(t.Label.Text) >= 1 {
			t.Label.Text = t.Label.Text[:len(t.Label.Text)-1]
		}
	}
    t.Counter++
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

