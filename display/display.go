package display

import (
	"github.com/enolgor/go-gfxino-codec/color"
)

type Display interface {
	SetRotate(rotate bool)
	SetFlip(flip bool)
	SetBrightness(brightness uint8)
	Display()
	Delay10MS(t uint8)
	Delay1S(t uint8)
	Delay1M(t uint8)
	ClearDisplay()
	SetColor(color *color.Color)
	ClearColor()
	FillScreen(color *color.Color)
	DrawPixel(x, y uint16, color *color.Color)
	DrawLine(x1, y1, x2, y2 uint16, color *color.Color)
	DrawFastVLine(x, y, l uint16, color *color.Color)
	DrawFastHLine(x, y, l uint16, color *color.Color)
	DrawRect(x, y, w, h uint16, color *color.Color)
	FillRect(x, y, w, h uint16, color *color.Color)
	DrawCircle(x, y, r uint16, color *color.Color)
	FillCircle(x, y, r uint16, color *color.Color)
	DrawRoundRect(x, y, w, h, r uint16, color *color.Color)
	FillRoundRect(x, y, w, h, r uint16, color *color.Color)
	DrawTriangle(x1, y1, x2, y2, x3, y3 uint16, color *color.Color)
	FillTriangle(x1, y1, x2, y2, x3, y3 uint16, color *color.Color)
	SetTextWrap(textWrap bool)
	SetCursor(x, y int16)
	SetTextSize(x uint8)
	SetTextColor(fg *color.Color, bg *color.Color)
	Print(text string)
	Read(p []byte) (int, error)
}
