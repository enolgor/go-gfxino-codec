package color

import (
	"fmt"
	"strconv"
	"strings"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func FromRGB(rgb uint32) *Color {
	return &Color{R: uint8(rgb >> 16), G: uint8(rgb >> 8), B: uint8(rgb)}
}

func FromString(str string) (*Color, error) {
	if strings.Index(str, ",") != -1 {
		return fromRGBString(str)
	}
	return fromHexString(str)
}

func fromRGBString(rgbStr string) (*Color, error) {
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

func fromHexString(hexStr string) (*Color, error) {
	hex, err := strconv.ParseUint(hexStr, 0, 32)
	if err != nil {
		return nil, err
	}
	r := uint8((hex >> 16) & 0xFF)
	g := uint8((hex >> 8) & 0xFF)
	b := uint8((hex >> 0) & 0xFF)
	return &Color{R: r, G: g, B: b}, nil
}

func From565(c uint16) *Color {
	return &Color{
		R: uint8((c >> 8) & 0xF8),
		G: uint8((c >> 3) & 0xFC),
		B: uint8(c << 3),
	}
}

func (c *Color) ToHexString() string {
	hex := (uint32(c.R) << 16) | (uint32(c.G) << 8) | uint32(c.B)
	return fmt.Sprintf("0x%X", hex)
}

func (c *Color) ToRGBString() string {
	return fmt.Sprintf("%d,%d,%d", c.R, c.G, c.B)
}

func (c *Color) To565() uint16 {
	return ((uint16(c.R) & 0xF8) << 8) | ((uint16(c.G) & 0xFC) << 3) | (uint16(c.B) >> 3)
}
