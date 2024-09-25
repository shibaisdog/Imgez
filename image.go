package Imgez

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/shibaisdog/Imgez/color"
)

type Pixel []color.RGBA

type Image []Pixel

func NewImage(p color.RGBA, w uint, h uint) Image {
	New_Image := Image{}
	for i := uint(0); i < h; i++ {
		ImageW := Pixel{}
		for j := uint(0); j < w; j++ {
			ImageW = append(ImageW, p)
		}
		New_Image = append(New_Image, ImageW)
	}
	return New_Image
}

func ReadImage(filename string) (Image, error) {
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

func (img Image) SavePNG(filename string) error {
	newImg := Imgez_To_Image(img)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, newImg)
	if err != nil {
		return err
	}
	return nil
}

func (baseImg *Image) Overlay(overlayImg Image, x, y int) {
	baseImage := Imgez_To_Image(*baseImg)
	overlayImage := Imgez_To_Image(overlayImg)
	bounds := baseImage.Bounds()
	newImage := image.NewRGBA(bounds)
	draw.Draw(newImage, bounds, baseImage, image.Point{}, draw.Src)
	overlayBounds := overlayImage.Bounds()
	draw.Draw(newImage, overlayBounds.Add(image.Point{X: x, Y: y}), overlayImage, image.Point{}, draw.Over)
	*baseImg = Image_To_Imgez(newImage)
}
