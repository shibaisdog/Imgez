package imgez

import (
	"image"
	"image/color"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type PixelX []Pixel

type Image []PixelX

func Imgez_To_Image(img Image) image.Image {
	height := len(img)
	if height == 0 {
		return nil
	}
	width := len(img[0])
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := img[y][x]
			newImg.Set(x, y, color.RGBA{p.R, p.G, p.B, p.A})
		}
	}
	return newImg
}

func Image_To_Imgez(img image.Image) Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := make(Image, height)
	for y := 0; y < height; y++ {
		row := make(PixelX, width)
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			row[x] = Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
		}
		newImg[y] = row
	}
	return newImg
}
