package gui

import (
	"fmt"
	"image/color"

	"github.com/calgo/parser"
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
    Events []parser.Event
}

func NewBoxContainer(xCount, yCount int, image *ebiten.Image, events []parser.Event) (*BoxContainer, error) {
    container := new(BoxContainer)
    boxes := make([]*Box, 0)
    container.xCount = xCount
    container.yCount = yCount
    container.Boxes = boxes
    container.Mode = "default"
    container.Image = image
    container.Events = events

    counter := 1
    for j := 0; j < container.yCount; j++ {
        for i := 0; i < container.xCount; i++ {
            boximage := ebiten.NewImage(boxDimension, boxDimension)
            boximage.Fill(SlateGray)

            dialogimage := ebiten.NewImage(2*boxDimension, 2*boxDimension)
            dialogimage.Fill(Teal)

            labels := make([]*Label, 0)
            x := i * boxDimension + (i+1)*boxMargin
            y := j * boxDimension + (j+1)*boxMargin
            title := NewLabel(boxMargin, boxMargin*2, fmt.Sprint(counter), color.Black, 18)
            label := NewLabel(boxMargin, boxMargin*4, container.Events[counter-1].Summary, color.Black, 14)
            box := NewBox(x, y, boxDimension, boximage, dialogimage, title, labels)
            counter++
            box.AddLabel(label)
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
            b.DialogBox.Update()
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
