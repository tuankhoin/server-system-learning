package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle{
	// Bound size to your preference
	return image.Rect(0, 0, 255, 255)
}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	// You can set as you like
	return color.RGBA{uint8(x*y), uint8(x^y), 120, 160}
}



func main() {
	m := Image{}
	pic.ShowImage(m)
}
