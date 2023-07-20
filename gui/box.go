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

func NewBox(x, y, dimension int, boximage *ebiten.Image, title *Label, labels []*Label) *Box {
    dialogWidth, dialogHeight := dimension*4, 2*dimension/3
    dialogimage := ebiten.NewImage(dialogWidth, dialogHeight)
	dialogimage.Fill(Teal)
	return &Box{
		X:         x,
		Y:         y,
		Dimension: dimension,
		Image:     boximage,
		Dialog:    NewDialogBox(dialogWidth, dialogHeight, dialogimage),
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
    Width int
    Height int
	Image     *ebiten.Image
	Visible   bool
	Buttons   map[string]*Button // close
    TextBoxes map[string]*TextBox // summary
    Labels    []*Label
}

func NewDialogBox(width, height int, image *ebiten.Image) *DialogBox {
    screenWidth, screenHeight := WindowWidth, WindowHeight
    x, y := (screenWidth-width)/2, (screenHeight-height)/2
	buttonWidth := width / 5
	buttonHeight := height / 4
	closeButton := NewButton(x, y, width-(buttonWidth+10), height-(buttonHeight+10), buttonWidth, buttonHeight, SlateGray, "CLOSE", "close")
    summaryTextBox := NewTextBox(x, y, 10, 20, 50, 20, 18, "Summary")

    textBoxes := make(map[string]*TextBox)
    textBoxes["summary"] = summaryTextBox

	buttons := make(map[string]*Button)
    buttons["close"] = closeButton
	return &DialogBox{
		X:         x,
		Y:         y,
        Width:     width,
        Height:    height,
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

func (d *DialogBox) AddLabel(label *Label) {
    d.Labels = append(d.Labels, label)
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
		if mouseX <= d.X || mouseX >= d.X+d.Width || mouseY <= d.Y || mouseY >= d.Y+d.Height {
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
    Title *Label
    Counter int
}

func NewTextBox(x, y, relX, relY, width, height int, size float64, title string) *TextBox {
    title = title + ":"
    titleLabel := NewLabel(relX, relY, title, color.Black, size+2)
    _, heightTitle := titleLabel.TextDimension()
    nextLine := (relY + heightTitle)
    label := NewLabel(relX, nextLine, "", color.Black, size)
    return &TextBox{
        X:         x,
        Y:         y,
        RelX:      relX,
        RelY:      relY,
        Width:     width,
        Height:    height,
        Size:      size,
        Label:    label,
        Title:    titleLabel,
        Counter: 0,
    }
}

func (t *TextBox) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(t.X), float64(t.Y))
    return options
}

func (t *TextBox) Draw(screen *ebiten.Image) {
    t.Title.Draw(screen)
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
	if len(t.Label.Text) <= 60 {
        t.Label.Text += string(t.Label.Runes)
	} 

 //    if len(t.Label.Text)%28 == 0 && len(t.Label.Text) != 0 {
	// 	t.Label.Text += "\n"
	// }
    t.Label.Text = strings.Replace(t.Label.Text, "|", "", -1) 

	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(t.Label.Text) >= 1 {
			t.Label.Text = t.Label.Text[:len(t.Label.Text)-1]
		}
	}


    if t.Counter%60 < 30 && !strings.HasSuffix(t.Label.Text, "|") {
        t.Label.Text += "|"
    } else if t.Counter%60 >= 30 {
        t.Label.Text = strings.TrimSuffix(t.Label.Text, "|")
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

