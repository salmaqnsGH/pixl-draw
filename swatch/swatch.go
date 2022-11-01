package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/salmaqnsGH/pixl-draw/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(color color.Color) {
	s.Color = color
	s.Refresh()
}

func (s *Swatch) NewSwatch(atate *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}
	swatch.ExtendBaseWidget(swatch)
	return swatch
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
