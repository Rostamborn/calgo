package gui

import (
	"log"

	"github.com/calgo/parser"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {

}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 0, 0
}

func Initialize(events []parser.Event) {
    game := &Game{
    }

    ebiten.SetWindowSize(800, 800)
    ebiten.SetWindowTitle("Calgo")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
	
}
