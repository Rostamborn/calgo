package gui

import (
	"image/color"
    "log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
    X int
    Y int
    Width int
    Height int
    Image *ebiten.Image
    Label *Label
}

func NewButton(x, y, width, height int, color color.Color, text string) *Button {
    image := ebiten.NewImage(width, height)
    image.Fill(color)

    font, err := CreateFontFace(18, 72)
    if err != nil {
        log.Fatal(err)
    }

    label := NewLabel(x + width / 2, y + height / 2, text, font)

    return &Button{
        X: x,
        Y: y,
        Width: width,
        Height: height,
        Image: image,
        Label: label,
    }
}

func (b *Button) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(b.X), float64(b.Y))
    return options
}

func (b *Button) Draw(screen *ebiten.Image) {

}


func (b *Button) Click() {

}

func (b *Button) Hover() {

}

func (b *Button) UnHover() {

}
