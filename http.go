package Imgez

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"

	"golang.org/x/image/webp"
)

func UrlImage(url string) (Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get a valid response: " + resp.Status)
	}
	contentType := resp.Header.Get("Content-Type")
	var img image.Image
	if contentType == "image/webp" {
		img, err = webp.Decode(resp.Body)
	} else if contentType == "image/png" {
		img, err = png.Decode(resp.Body)
	} else if contentType == "image/jpeg" || contentType == "image/jpg" {
		img, err = jpeg.Decode(resp.Body)
	} else {
		img, _, err = image.Decode(resp.Body)
	}
	if err != nil {
		return nil, err
	}
	return Image_To_Imgez(img), nil
}
