package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	X      int
	Y      int
	RelX   int
	RelY   int
	Width  int
	Height int
	Image  *ebiten.Image
	Label  *Label
	Color  color.Color
    Tag   string // close
}

func NewButton(x, y, relX, relY, width, height int, col color.Color, text, tag string) *Button {
	image := ebiten.NewImage(width, height)
	image.Fill(col)
	label := NewLabel(relX, relY, text, color.Black, 18)

	return &Button{
		X:      x,
		Y:      y,
		RelX:   relX,
		RelY:   relY,
		Width:  width,
		Height: height,
		Image:  image,
		Label:  label,
		Color:  col,
        Tag:    tag,
	}
}

func (b *Button) SetOptions() *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.RelX), float64(b.RelY))
	return options
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.Label.DrawCentered(b.Image) // Draw label on Button image
	// b.Label.Draw(b.Image)
	screen.DrawImage(b.Image, b.SetOptions())
    b.Image.Fill(b.Color)
}

func (b *Button) Clicked() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x := b.X + b.RelX
		y := b.Y + b.RelY
		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX >= x && mouseX <= x+b.Width && mouseY >= y && mouseY <= y+b.Height {
			return true
		}
	}
	return false

}

func (b *Button) Hover() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	x := b.X + b.RelX
	y := b.Y + b.RelY
	if mouseX >= x && mouseX <= x+b.Width && mouseY >= y && mouseY <= y+b.Height {
		return true
	}
	return false
}

func (b *Button) UnHover() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX <= b.X || mouseX >= b.X+b.Width || mouseY <= b.Y || mouseY >= b.Y+b.Height {
		return true
	}
	return false
}

func (b *Button) Update() {
	if b.Hover() {
		b.Image.Fill(DarkGray)
		if b.Clicked() {
			b.Image.Fill(color.Black)
		}
	} else if b.UnHover() {
		b.Image.Fill(b.Color)
	}
}
