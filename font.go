package Imgez

import (
	"errors"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func (baseImg *Image) Text(x, y int, Font *sfnt.Font, size float64, text string) error {
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
		Src:  image.NewUniform(color.Black),
		Face: fontFace,
		Dot:  fixed.P(x+int(size), y),
	}

	d.DrawString(text)

	*baseImg = Image_To_Imgez(rgbaImage)
	return nil
}
