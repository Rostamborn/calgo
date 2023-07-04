package gui

import (
	"log"
    "image/color"

	"github.com/calgo/parser"
	"github.com/hajimehoshi/ebiten/v2"
)

var boxImage *ebiten.Image
// var containerImage *ebiten.Image
var dialogImage *ebiten.Image

func init() {
    boxImage = ebiten.NewImage(100, 100)
    boxImage.Fill(color.RGBA{0, 0xff, 0, 0xff})

    dialogImage = ebiten.NewImage(200, 200)
    dialogImage.Fill(color.RGBA{39, 0x1a, 0xe8, 0xcc})
}

type Game struct {
    container *BoxContainer
}

func (g *Game) Update() error {
    g.container.Update()
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
   g.container.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 800, 800
}

func Initialize(events []parser.Event) {
    container, _ := NewBoxContainer(7, 4)
    game := &Game{
        container: container,
    }

    ebiten.SetWindowSize(800, 800)
    ebiten.SetWindowTitle("Calgo")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
	
}
