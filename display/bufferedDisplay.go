package display

import (
	"bytes"

	"github.com/enolgor/go-gfxino-codec/color"
	"github.com/enolgor/go-gfxino-codec/commands"
	"github.com/enolgor/go-gfxino-codec/text"
)

type BufferedDisplay struct {
	bitSize8 bool
	colorSet bool
	Buffer   bytes.Buffer
}

func (bd *BufferedDisplay) SetBitSize8() {
	bd.Buffer.WriteByte(byte(commands.SETBITSIZE8))
	bd.bitSize8 = true
}

func (bd *BufferedDisplay) SetRotateON() {
	bd.Buffer.WriteByte(byte(commands.SETROTATEON))
}

func (bd *BufferedDisplay) SetRotateOFF() {
	bd.Buffer.WriteByte(byte(commands.SETROTATEOFF))
}

func (bd *BufferedDisplay) SetFlipON() {
	bd.Buffer.WriteByte(byte(commands.SETFLIPON))
}

func (bd *BufferedDisplay) SetFlipOFF() {
	bd.Buffer.WriteByte(byte(commands.SETFLIPOFF))
}

func (bd *BufferedDisplay) SetBrightness(brightness uint8) {
	bd.Buffer.WriteByte(byte(commands.SETBRIGHTNESS))
	bd.Buffer.WriteByte(byte(brightness))
}

func (bd *BufferedDisplay) Display() {
	bd.Buffer.WriteByte(byte(commands.DISPLAY))
}

func (bd *BufferedDisplay) Delay10MS(t uint8) {
	bd.Buffer.WriteByte(byte(commands.DELAY10MS))
	bd.Buffer.WriteByte(byte(t))
}

func (bd *BufferedDisplay) Delay1S(t uint8) {
	bd.Buffer.WriteByte(byte(commands.DELAY1S))
	bd.Buffer.WriteByte(byte(t))
}

func (bd *BufferedDisplay) Delay1M(t uint8) {
	bd.Buffer.WriteByte(byte(commands.DELAY1M))
	bd.Buffer.WriteByte(byte(t))
}

func (bd *BufferedDisplay) ClearDisplay() {
	bd.Buffer.WriteByte(byte(commands.CLEARDISPLAY))
}

func (bd *BufferedDisplay) SetColor(c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.SETCOLOR))
	bd.colorSet = true
	bd.writeColor(c)
}

func (bd *BufferedDisplay) ClearColor() {
	bd.Buffer.WriteByte(byte(commands.CLEARCOLOR))
	bd.colorSet = false
}

func (bd *BufferedDisplay) FillScreen(c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.FILLSCREEN))
	bd.writeColor(c)
}

func (bd *BufferedDisplay) DrawPixel(x, y uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWPIXEL))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawLine(x1, y1, x2, y2 uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWLINE))
	bd.writeUint(x1)
	bd.writeUint(y1)
	bd.writeUint(x2)
	bd.writeUint(y2)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawFastVLine(x, y, l uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWFASTVLINE))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(l)
	bd.writeSkippableColor(c)
}
func (bd *BufferedDisplay) DrawFastHLine(x, y, l uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWFASTHLINE))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(l)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawRect(x, y, w, h uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWRECT))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(w)
	bd.writeUint(h)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) FillRect(x, y, w, h uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.FILLRECT))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(w)
	bd.writeUint(h)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawCircle(x, y, r uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWCIRCLE))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(r)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) FillCircle(x, y, r uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.FILLCIRCLE))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(r)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawRoundRect(x, y, w, h, r uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWROUNDRECT))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(w)
	bd.writeUint(h)
	bd.writeUint(r)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) FillRoundRect(x, y, w, h, r uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.FILLROUNDRECT))
	bd.writeUint(x)
	bd.writeUint(y)
	bd.writeUint(w)
	bd.writeUint(h)
	bd.writeUint(r)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) DrawTriangle(x1, y1, x2, y2, x3, y3 uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.DRAWTRIANGLE))
	bd.writeUint(x1)
	bd.writeUint(y1)
	bd.writeUint(x2)
	bd.writeUint(y2)
	bd.writeUint(x3)
	bd.writeUint(y3)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) FillTriangle(x1, y1, x2, y2, x3, y3 uint16, c *color.Color) {
	bd.Buffer.WriteByte(byte(commands.FILLTRIANGLE))
	bd.writeUint(x1)
	bd.writeUint(y1)
	bd.writeUint(x2)
	bd.writeUint(y2)
	bd.writeUint(x3)
	bd.writeUint(y3)
	bd.writeSkippableColor(c)
}

func (bd *BufferedDisplay) SetTextWrapON() {
	bd.Buffer.WriteByte(byte(commands.SETTEXTWRAPON))
}

func (bd *BufferedDisplay) SetTextWrapOFF() {
	bd.Buffer.WriteByte(byte(commands.SETTEXTWRAPOFF))
}

func (bd *BufferedDisplay) SetCursor(x, y int16) {
	bd.Buffer.WriteByte(byte(commands.SETCURSOR))
	bd.writeInt16(x)
	bd.writeInt16(y)
}

func (bd *BufferedDisplay) SetTextSize(x uint8) {
	bd.Buffer.WriteByte(byte(commands.SETTEXTSIZE))
	bd.Buffer.WriteByte(byte(x))
}

func (bd *BufferedDisplay) SetTextColor(fg *color.Color, bg *color.Color) {
	bd.Buffer.WriteByte(byte(commands.SETTEXTCOLOR))
	bd.writeColor(fg)
	bd.writeColor(bg)
}

func (bd *BufferedDisplay) Print(text string) {
	bd.Buffer.WriteByte(byte(commands.PRINT))
	bd.writeText(text)
}

func (bd *BufferedDisplay) writeUint(v uint16) {
	if !bd.bitSize8 {
		bd.Buffer.WriteByte(byte(v >> 8))
	}
	bd.Buffer.WriteByte(byte(v & 0xFF))
}

func (bd *BufferedDisplay) writeInt16(v int16) {
	bd.Buffer.WriteByte(byte(v >> 8))
	bd.Buffer.WriteByte(byte(v & 0xFF))
}

func (bd *BufferedDisplay) writeColor(c *color.Color) {
	c565 := c.To565()
	bd.Buffer.WriteByte(byte(c565 >> 8))
	bd.Buffer.WriteByte(byte(c565 & 0xFF))
}

func (bd *BufferedDisplay) writeSkippableColor(c *color.Color) {
	if bd.colorSet {
		return
	}
	bd.writeColor(c)
}

func (bd *BufferedDisplay) writeText(txt string) {
	encoded := text.EncodeCP437(txt)
	if len(encoded) > 255 {
		encoded = encoded[:255]
	}
	bd.Buffer.WriteByte(byte(len(encoded)))
	bd.Buffer.Write(encoded)
}

func (bd *BufferedDisplay) Read(p []byte) (int, error) {
	n := 0
	for {
		cmd, err := bd.Buffer.ReadByte()
		if err != nil {
			return n, err //io.EOF
		}
		size := commands.GetArgSize(cmd, bd.bitSize8, bd.colorSet)
		if size+1+n > len(p) {
			bd.Buffer.UnreadByte()
			return n, nil
		}
		p[n] = cmd
		_, err = bd.Buffer.Read(p[n+1 : n+1+size])
		if err != nil {
			return n, err
		}
		n += size + 1
	}
}
