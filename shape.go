package Imgez

import (
	"image"
	"image/color"
	"math"
)

// percentage (0 ~ 100)
func (img *Image) Shape(percentage int) {
	rnewImg := Imgez_To_Image(*img)
	bounds := rnewImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	newImg := image.NewRGBA(bounds)
	radius := float64(width) / 2.0
	roundness := float64(percentage) / 100.0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			centerX := float64(width) / 2.0
			centerY := float64(height) / 2.0
			distance := math.Sqrt((centerX-float64(x))*(centerX-float64(x)) + (centerY-float64(y))*(centerY-float64(y)))
			if distance <= radius*roundness {
				newImg.Set(x, y, img.At(x, y))
			} else {
				newImg.Set(x, y, color.Transparent)
			}
		}
	}

	*img = Image_To_Imgez(newImg)
}

// opacity (0.0 ~ 1.0)
func (img *Image) Opacity(opacity float64) {
	newImg := Imgez_To_Image(*img)
	bounds := newImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	newRGBA := image.NewRGBA(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldColor := newImg.At(x, y)
			r, g, b, a := oldColor.RGBA()
			newA := uint8(float64(a>>8) * opacity)
			newColor := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), newA}
			newRGBA.Set(x, y, newColor)
		}
	}

	*img = Image_To_Imgez(newRGBA)
}
