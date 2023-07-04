package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const(
    boxDimension = 150
    boxMargin = 10
)

type BoxContainer struct {
    xCount int
    yCount int
    Boxes []*Box
    Mode string
    DialogBox *DialogBox
    Image *ebiten.Image
}

func NewBoxContainer(xCount, yCount int, image *ebiten.Image) (*BoxContainer, error) {
    container := new(BoxContainer)
    boxes := make([]*Box, 0)
    container.xCount = xCount
    container.yCount = yCount
    container.Boxes = boxes
    container.Mode = "default"
    container.Image = image

    for j := 0; j < container.yCount; j++ {
        for i := 0; i < container.xCount; i++ {
            x := i * boxDimension + (i+1)*boxMargin
            y := j * boxDimension + (j+1)*boxMargin
            box := NewBox(x, y, boxDimension, boxImage, dialogImage)
            container.AddBox(box)
        }
    }

    return container, nil
}

func (b *BoxContainer) AddBox(box *Box) {
    b.Boxes = append(b.Boxes, box)
}

func (b *BoxContainer) Draw(screen *ebiten.Image) {
    screen.DrawImage(b.Image, nil)
    if b.Mode == "dialogmode" {
        for _, box := range b.Boxes {
            box.Draw(screen)
        }
        if b.DialogBox != nil && b.DialogBox.Visible {
            b.DialogBox.Draw(screen)
        }
    } else {
       for _, box := range b.Boxes {
            box.Draw(screen)
        } 
    }
}

func (b *BoxContainer) Update() {
    if b.Mode == "dialogmode" {
        if b.DialogBox != nil {
            ok := b.DialogBox.ClickToExit()
            if ok {
                b.Mode = "default"
            }
        }
    } else if b.Mode == "default" {
        for _, box := range b.Boxes {
            ok := box.Clicked()
            if ok {
                b.Mode = "dialogmode"
                b.DialogBox = box.Dialog
            }
        }
    }
}
