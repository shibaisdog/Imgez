package Imgez

import (
	"errors"
	"image"

	"image/color"

	clr "github.com/shibaisdog/Imgez/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func (baseImg *Image) Text(Font *sfnt.Font, Color clr.RGBA, x, y int, size float64, text string) error {
	baseImage := Imgez_To_Image(*baseImg)
	rgbaImage, ok := baseImage.(*image.RGBA)
	if !ok {
		return errors.New("failed to convert base image to RGBA")
	}
	fontFace, err := opentype.NewFace(Font, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return err
	}
	d := &font.Drawer{
		Dst:  rgbaImage,
		Src:  image.NewUniform(color.RGBA{Color.R, Color.G, Color.B, Color.A}),
		Face: fontFace,
		Dot:  fixed.P(x, y),
	}

	d.DrawString(text)

	*baseImg = Image_To_Imgez(rgbaImage)
	return nil
}
