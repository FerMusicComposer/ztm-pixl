package apptypes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor   color.Color
	BrushType    int
	SwatchSelect int
	FilePath     string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
