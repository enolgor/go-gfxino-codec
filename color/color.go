package color

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func ParseRGB(str string) (*Color, error) {
	if strings.Index(str, ",") != -1 {
		return parseRGBString(str)
	}
	return parseRGBHexString(str)
}

func parseRGBString(rgbStr string) (*Color, error) {
	vals := strings.Split(rgbStr, ",")
	if len(vals) != 3 {
		return nil, fmt.Errorf("Invalid format. Should be R,G,B")
	}
	c := Color{}
	for i, v := range vals {
		x, e := strconv.ParseUint(v, 0, 8)
		if e != nil {
			return nil, e
		}
		switch i {
		case 0:
			c.R = uint8(x)
		case 1:
			c.G = uint8(x)
		case 2:
			c.B = uint8(x)
		}
	}
	return &c, nil
}

func parseRGBHexString(hexStr string) (*Color, error) {
	hex, err := strconv.ParseUint(hexStr, 0, 32)
	if err != nil {
		return nil, err
	}
	r := uint8((hex >> 16) & 0xFF)
	g := uint8((hex >> 8) & 0xFF)
	b := uint8((hex >> 0) & 0xFF)
	return &Color{R: r, G: g, B: b}, nil
}

func FromColor(c color.Color) *Color {
	r, g, b, _ := c.RGBA() // investigar...
	return &Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}
}

func From565(c uint16) *Color {
	return &Color{
		R: uint8((c >> 8) & 0xF8),
		G: uint8((c >> 3) & 0xFC),
		B: uint8(c << 3),
	}
}

func FromRGB(rgb uint32) *Color {
	return &Color{R: uint8(rgb >> 16), G: uint8(rgb >> 8), B: uint8(rgb)}
}

func (c *Color) ToRGBHexString() string {
	hex := (uint32(c.R) << 16) | (uint32(c.G) << 8) | uint32(c.B)
	return fmt.Sprintf("0x%X", hex)
}

func (c *Color) ToRGBString() string {
	return fmt.Sprintf("%d,%d,%d", c.R, c.G, c.B)
}

func (c *Color) To565() uint16 {
	return ((uint16(c.R) & 0xF8) << 8) | ((uint16(c.G) & 0xFC) << 3) | (uint16(c.B) >> 3)
}

// HSV

func FromHSV(H, S, V uint) *Color {
	if H > 360 {
		H = 360
	}
	if S > 100 {
		S = 100
	}
	if V > 100 {
		V = 100
	}
	h := float64(H)
	s := float64(S) / 100
	v := float64(V) / 100

	hi := H / 60
	f := h/60 - float64(hi)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	maxRGB := func(r, g, b float64) *Color {
		return &Color{
			R: uint8(math.Round(r * 255)),
			G: uint8(math.Round(g * 255)),
			B: uint8(math.Round(b * 255)),
		}
	}

	switch hi {
	case 0:
		return maxRGB(v, t, p)
	case 1:
		return maxRGB(q, v, p)
	case 2:
		return maxRGB(p, v, t)
	case 3:
		return maxRGB(p, q, v)
	case 4:
		return maxRGB(t, p, v)
	case 5:
		return maxRGB(v, p, q)
	}
	return maxRGB(0, 0, 0)
}
