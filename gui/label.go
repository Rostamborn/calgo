package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Label struct {
    X int
    Y int
    Background *ebiten.Image
    text string
}
