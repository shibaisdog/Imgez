package Imgez

import (
	"errors"
	"image"
	"net/http"
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
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return Image_To_Imgez(img), nil
}
