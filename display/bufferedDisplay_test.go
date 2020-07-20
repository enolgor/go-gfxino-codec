package display

import (
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/enolgor/go-gfxino-codec/color"
)

func TestBufferedDisplay(t *testing.T) {
	display := BufferedDisplay{Mode8Bit: true}
	display.SetColor(&color.Color{0, 0xF1, 0xA3})
	display.DrawPixel(1, 5, nil)
	display.Display()
	display.Delay10MS(5)
	display.DrawTriangle(1, 1, 10, 10, 10, 1, nil)
	/*r := bytes.NewReader(display.Buffer.Bytes())
	var w io.Writer = os.Stdout
	ip := &interpreter.Interpreter{}
	ip.Interpret(r, w)*/
	buf := make([]byte, 16)
	n, err := display.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(buf[:n]))
	n, err = display.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(buf[:n]))
}
