package Imgez

import (
	"errors"
	"image"
	"os"

	"image/color"

	clr "github.com/shibaisdog/Imgez/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type Font struct {
	F *sfnt.Font
}

func LoadFont(file string) (Font, error) {
	fontBytes, err := os.ReadFile(file)
	if err != nil {
		return Font{nil}, err
	}
	ttfFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return Font{nil}, err
	}
	return Font{ttfFont}, nil
}

func (f Font) Measure(text string, size float64) (float64, float64, error) {
	var buf sfnt.Buffer
	ppem := fixed.Int26_6(size * 64)
	scale := float64(ppem) / 64.0
	var totalWidth fixed.Int26_6
	var maxHeight fixed.Int26_6
	for _, r := range text {
		glyphIndex, err := f.F.GlyphIndex(&buf, r)
		if err != nil {
			return 0, 0, err
		}
		advanceWidth, err := f.F.GlyphAdvance(&buf, glyphIndex, ppem, font.HintingNone)
		if err != nil {
			return 0, 0, err
		}
		bounds, _, err := f.F.GlyphBounds(&buf, glyphIndex, ppem, font.HintingNone)
		if err != nil {
			return 0, 0, err
		}
		totalWidth += advanceWidth
		if bounds.Max.Y-bounds.Min.Y > maxHeight {
			maxHeight = bounds.Max.Y - bounds.Min.Y
		}
	}
	width := float64(totalWidth) / 64.0
	height := float64(maxHeight) / 64.0
	return width * scale, height * scale, nil
}

func (baseImg *Image) Text(Font Font, Color clr.RGBA, text string, size float64, l Location) error {
	baseImage := Imgez_To_Image(*baseImg)
	rgbaImage, ok := baseImage.(*image.RGBA)
	if !ok {
		return errors.New("failed to convert base image to RGBA")
	}
	fontFace, err := opentype.NewFace(Font.F, &opentype.FaceOptions{
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
		Dot:  fixed.P(l.X, l.Y),
	}
	d.DrawString(text)
	*baseImg = Image_To_Imgez(rgbaImage)
	return nil
}
