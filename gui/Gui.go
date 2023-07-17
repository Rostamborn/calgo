package gui

import (
	"image/color"
	"log"
	"os"

	"github.com/calgo/parser"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

const (
	WindowWidth  = 1400
	WindowHeight = 900
)

var FontTT *sfnt.Font

// var boxImage *ebiten.Image
var containerImage *ebiten.Image

// var dialogImage *ebiten.Image
var SlateGray = color.RGBA{0x70, 0x80, 0x90, 0xff}
var DarkSlateGrey = color.RGBA{0x2f, 0x4f, 0x4f, 0xff}
var LightGreen = color.RGBA{0x90, 0xee, 0x90, 0xff}
var DarkGray = color.RGBA{0xa9, 0xa9, 0xa9, 0xff}
var Teal = color.RGBA{0x00, 0x80, 0x80, 0xff}

func init() {
	// Background Image
	containerImage = ebiten.NewImage(WindowWidth, WindowHeight)
	containerImage.Fill(DarkSlateGrey)

	fontFile := "JetBrainsMono-ExtraBold.ttf"
	fontData, err := os.ReadFile(fontFile)
	if err != nil {
		log.Fatal("Failed to load font:", err)
	}

	FontTT, err = opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	container *BoxContainer
	font      font.Face
}

func (g *Game) Update() error {
	g.container.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.container.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func Initialize(events []parser.Event) {
	FontFace, err := opentype.NewFace(FontTT, &opentype.FaceOptions{
		Size: 24,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}
	container, _ := NewBoxContainer(7, 4, containerImage, events)
	game := &Game{
		container: container,
		font:      FontFace,
	}

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Calgo")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
