package gui

import (
	"image/color"

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

func NewLabel(x, y int, text string, fontface font.Face) *Label {
    // FontFace, err := opentype.NewFace(FontTT, &opentype.FaceOptions{
    //     Size: 18,
    //     DPI: 72,
    // })
    // if err != nil {
    //     log.Fatal(err)
    // }
    // FontFace, err := CreateFontFace(18, 72)
    // if err != nil {
    //     log.Fatal(err)
    // }

    background := ebiten.NewImage(labelWidth, labelHeight)
    background.Fill(color.White)
    return &Label{
        X: x,
        Y: y,
        Background: background,
        text: text,
        font: fontface,
    }
}

func CreateFontFace(size, dpi float64) (font.Face, error) {
    fontFace, err := opentype.NewFace(FontTT, &opentype.FaceOptions{
        Size: size,
        DPI: dpi,
    })
    if err != nil {
        return nil, err
    }
    return fontFace, nil
}
