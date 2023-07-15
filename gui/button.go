package gui

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
    X int
    Y int
    RelX int
    RelY int
    Width int
    Height int
    Image *ebiten.Image
    Label *Label
    Color color.Color
}

func NewButton(x, y, relX, relY, width, height int, col color.Color, text string) *Button {
    image := ebiten.NewImage(width, height)
    image.Fill(col)

    font, err := CreateFontFace(18, 72)
    if err != nil {
        log.Fatal(err)
    }

    label := NewLabel(width/2 , height/2 , text, font, color.Black)

    return &Button{
        X: x,
        Y: y,
        RelX: relX,
        RelY: relY,
        Width: width,
        Height: height,
        Image: image,
        Label: label,
        Color: col,
    }
}

func (b *Button) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(b.RelX), float64(b.RelY))
    return options
}

func (b *Button) Draw(screen *ebiten.Image) {
    b.Label.Draw(b.Image)
    screen.DrawImage(b.Image, b.SetOptions())
}


func (b *Button) Click() bool {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        x := b.X + b.RelX
        y := b.Y + b.RelY
        mouseX, mouseY := ebiten.CursorPosition()
        if mouseX >= x && mouseX <= x + b.Width && mouseY >= y && mouseY <= y + b.Height {
            return true
        }
    }
    return false

}

func (b *Button) Hover() bool {
    mouseX, mouseY := ebiten.CursorPosition()
    x := b.X + b.RelX
    y := b.Y + b.RelY
    if mouseX >= x && mouseX <= x + b.Width && mouseY >= y && mouseY <= y + b.Height {
        return true
    }
    return false
}

func (b *Button) UnHover() bool {
    mouseX, mouseY := ebiten.CursorPosition()
    if mouseX <= b.X || mouseX >= b.X + b.Width || mouseY <= b.Y || mouseY >= b.Y + b.Height {
        return true
    }
    return false
}

func (b *Button) Update() {
    if b.Hover() {
        b.Image.Fill(color.White)
        if b.Click() {
        b.Image.Fill(color.Black)
        }
    } else if b.UnHover() {
        b.Image.Fill(b.Color)
    }
}