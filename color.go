package Imgez

type RGBA struct {
	R, G, B, A uint8
}

func (c *RGBA) RGBA(r, g, b, a uint8) *RGBA {
	return &RGBA{
		R: r, G: g, B: b, A: a,
	}
}

type RGB struct {
	R, G, B uint8
}

func (c *RGB) RGB(r, g, b uint8) *RGB {
	return &RGB{
		R: r, G: g, B: b,
	}
}
