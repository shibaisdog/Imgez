package font

import (
	"os"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

func LoadFont(file string, size, dpi float64) (*sfnt.Font, error) {
	fontBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ttfFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	return ttfFont, nil
}
