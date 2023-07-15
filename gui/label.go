package gui

import (
	"image/color"
    "log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
    labelWidth = 100
    labelHeight = 50
    DPI = 72
)

type Label struct {
    X int
    Y int
    Text string
    Font font.Face
    Color color.Color
    Size float64
    DPI float64
}

// will create a font based on the size. DPI is constant = 72
func NewLabel(x, y int, text string, col color.Color, size float64) *Label {
    font, err := CreateFontFace(size, DPI)
    if err != nil {
        log.Fatal(err)
    }

    return &Label{
        X: x,
        Y: y,
        Text: text,
        Font: font,
        Color: col,
        Size: size,
    }
}

func (b *Label) SetOptions() *ebiten.DrawImageOptions {
    options := &ebiten.DrawImageOptions{}
    options.GeoM.Translate(float64(b.X), float64(b.Y))
    return options
}

// we always draw the label on Button, Box, etc. and not the main screen
func (l *Label) Draw(screen *ebiten.Image) {
    text.Draw(screen, l.Text, l.Font, l.X, l.Y, l.Color)
}

// to draw on buttons in a centered way
func (l *Label) DrawCentered(screen *ebiten.Image) {
    x, y := l.GetCenterCord(screen)
    text.Draw(screen, l.Text, l.Font, x, y, l.Color)
}

func (l *Label) TextDimension() (int, int) {
    width := font.MeasureString(l.Font, l.Text).Floor()
	height := (l.Font.Metrics().Ascent + l.Font.Metrics().Descent).Floor()

	return width, height
}

// coordinations to draw the label in the center of the Button, etc.
func (l *Label) GetCenterCord(screen *ebiten.Image) (int, int) {
    screenWidth, screenHeight := screen.Size()
    TextWidth, TextHeight := l.TextDimension()
    x := (screenWidth - TextWidth) / 2
    y := (screenHeight - TextHeight) / 2 + l.Font.Metrics().Ascent.Floor()
    return x, y
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
