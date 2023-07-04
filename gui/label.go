package gui

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
    labelWidth = 100
    labelHeight = 50
)

type Label struct {
    X int
    Y int
    Background *ebiten.Image
    text string
    font font.Face
}

func NewLabel(x, y int, text string) *Label {
    FontFace, err := opentype.NewFace(FontTT, &opentype.FaceOptions{
        Size: 18,
        DPI: 72,
    })
    if err != nil {
        log.Fatal(err)
    }

    background := ebiten.NewImage(labelWidth, labelHeight)
    background.Fill(color.White)
    return &Label{
        X: x,
        Y: y,
        Background: background,
        text: text,
        font: FontFace,
    }
}
