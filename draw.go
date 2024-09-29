package Imgez

import (
	"image"
	"image/color"
	"math"
)

func (img *Image) DrawLine(clr RGBA, l1, l2 Location) {
	newImg := Imgez_To_Image(*img)
	bounds := newImg.Bounds()
	newRGBA := image.NewRGBA(bounds)
	dx := math.Abs(float64(l2.X - l1.X))
	dy := math.Abs(float64(l2.Y - l1.Y))
	var sx, sy int
	if l1.X < l2.X {
		sx = 1
	} else {
		sx = -1
	}
	if l1.Y < l2.Y {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy
	for {
		newRGBA.Set(l1.X, l1.Y, color.RGBA{clr.R, clr.G, clr.B, clr.A})
		if l1.X == l2.X && l1.Y == l2.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			l1.X += sx
		}
		if e2 < dx {
			err += dx
			l1.Y += sy
		}
	}
	*img = Image_To_Imgez(newRGBA)
}
