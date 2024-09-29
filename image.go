package Imgez

import (
	"image"
	"image/color"
	"image/draw"
	"os"

	clr "github.com/shibaisdog/Imgez/color"
)

type Pixel []clr.RGBA

type Image []Pixel

type Location struct {
	X, Y int
}

type Size struct {
	Width, Height uint
}

func (i *Image) At(l Location) color.Color {
	Img := Imgez_To_Image(*i)
	return Img.At(l.X, l.Y)
}

func (i *Image) Bounds() image.Rectangle {
	Img := Imgez_To_Image(*i)
	return Img.Bounds()
}

func NewImage(p clr.RGBA, s Size) Image {
	New_Image := Image{}
	for i := uint(0); i < s.Height; i++ {
		ImageW := Pixel{}
		for j := uint(0); j < s.Width; j++ {
			ImageW = append(ImageW, p)
		}
		New_Image = append(New_Image, ImageW)
	}
	return New_Image
}

func Open(filename string) (Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	newImage := Image_To_Imgez(img)
	return newImage, nil
}

func (baseImg *Image) Overlay(overlayImg Image, l Location) {
	baseImage := Imgez_To_Image(*baseImg)
	overlayImage := Imgez_To_Image(overlayImg)
	bounds := baseImage.Bounds()
	newImage := image.NewRGBA(bounds)
	draw.Draw(newImage, bounds, baseImage, image.Point{}, draw.Src)
	overlayBounds := overlayImage.Bounds()
	draw.Draw(newImage, overlayBounds.Add(image.Point{X: l.X, Y: l.Y}), overlayImage, image.Point{}, draw.Over)
	*baseImg = Image_To_Imgez(newImage)
}
