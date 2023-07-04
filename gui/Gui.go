package gui

import (
	"image/color"
	"log"
    "os"
    "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/calgo/parser"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const(
    WindowWidth = 1400
    WindowHeight = 900
)

var FontFace font.Face

var boxImage *ebiten.Image
var containerImage *ebiten.Image
var dialogImage *ebiten.Image

func init() {
    containerImage = ebiten.NewImage(WindowWidth, WindowHeight)
    containerImage.Fill(color.RGBA{0xff, 0, 0, 0xff})

    boxImage = ebiten.NewImage(boxDimension, boxDimension)
    boxImage.Fill(color.RGBA{0, 0xff, 0, 0xff})

    dialogImage = ebiten.NewImage(2*boxDimension, 2*boxDimension)
    dialogImage.Fill(color.RGBA{39, 0x1a, 0xe8, 0xcc})

    fontFile := "JetBrainsMono-ExtraBold.ttf"
	fontData, err := os.ReadFile(fontFile)
	if err != nil {
		log.Fatal("Failed to load font:", err)
	}

    tt, err := opentype.Parse(fontData)
    if err != nil {
        log.Fatal(err)
    }

    FontFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
        Size: 24,
        DPI: 72,
    })
    if err != nil {
        log.Fatal(err)
    }
}

type Game struct {
    container *BoxContainer
    font font.Face
}

func (g *Game) Update() error {
    g.container.Update()
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
   g.container.Draw(screen)
   text.Draw(screen, "cunt", g.font, 50 , 50, color.RGBA{R: 128, G: 0, B: 128, A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return WindowWidth, WindowHeight
}

func Initialize(events []parser.Event) {
    

    container, _ := NewBoxContainer(7, 4, containerImage)
    game := &Game{
        container: container,
        font: FontFace,
    }

    ebiten.SetWindowSize(WindowWidth, WindowHeight)
    ebiten.SetWindowTitle("Calgo")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
	
}
