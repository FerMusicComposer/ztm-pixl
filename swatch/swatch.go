package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/FerMusicComposer/ztm-pixl.git/apptypes"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (swatch *Swatch) SetColor(color color.Color) {
	swatch.Color = color
	swatch.Refresh()
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}

func NewSwatch(state *apptypes.State, color color.Color, swatchindex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchindex,
	}

	swatch.ExtendBaseWidget(swatch)

	return swatch
}
