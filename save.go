package Imgez

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"os"
)

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

func (img Image) SaveJPEG(filename string, quality int) error {
	newImg := Imgez_To_Image(img)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	options := jpeg.Options{
		Quality: quality,
	}
	err = jpeg.Encode(file, newImg, &options)
	if err != nil {
		return err
	}
	return nil
}

func (img Image) Byte() ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, Imgez_To_Image(img))
	if err != nil {
		return nil, err
	}
	imgBytes := buf.Bytes()
	return imgBytes, nil
}
