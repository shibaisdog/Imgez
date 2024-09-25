package Imgez

import (
	"image"

	"golang.org/x/image/draw"
)

func (img *Image) Resize(w, h uint) {
	baseImage := Imgez_To_Image(*img)
	newImg := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	draw.BiLinear.Scale(newImg, newImg.Rect, baseImage, baseImage.Bounds(), draw.Over, nil)
	*img = Image_To_Imgez(newImg)
}

func Resize(img Image, w, h uint) Image {
	baseImage := Imgez_To_Image(img)
	newImg := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	draw.BiLinear.Scale(newImg, newImg.Rect, baseImage, baseImage.Bounds(), draw.Over, nil)
	return Image_To_Imgez(newImg)
}
