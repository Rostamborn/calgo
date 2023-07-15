package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
    Text string
    Font font.Face
    Color color.Color
}

func NewLabel(x, y int, text string, fontface font.Face, col color.Color) *Label {
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
        Text: text,
        Font: fontface,
        Color: col,
    }
}

func (b *Label) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(b.X), float64(b.Y))
    return options
}

func (l *Label) Draw(screen *ebiten.Image) {
    text.Draw(screen, l.Text, l.Font, l.X, l.Y, l.Color)
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
