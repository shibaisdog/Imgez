package Imgez

import (
	"image"

	"golang.org/x/image/draw"
)

func (img *Image) Resize(s Size) {
	baseImage := Imgez_To_Image(*img)
	newImg := image.NewRGBA(image.Rect(0, 0, int(s.Width), int(s.Height)))
	draw.BiLinear.Scale(newImg, newImg.Rect, baseImage, baseImage.Bounds(), draw.Over, nil)
	*img = Image_To_Imgez(newImg)
}

func (img Image) Getsize() Size {
	height := len(img)
	if height == 0 {
		return Size{Width: uint(0), Height: uint(0)}
	}
	width := len(img[0])
	return Size{Width: uint(width), Height: uint(height)}
}

func Resize(img Image, s Size) Image {
	baseImage := Imgez_To_Image(img)
	newImg := image.NewRGBA(image.Rect(0, 0, int(s.Width), int(s.Height)))
	draw.BiLinear.Scale(newImg, newImg.Rect, baseImage, baseImage.Bounds(), draw.Over, nil)
	return Image_To_Imgez(newImg)
}
