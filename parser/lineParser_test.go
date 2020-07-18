package parser

import (
	"encoding/hex"
	"testing"
)

func testLineHex(t *testing.T, p *Parser, buf []byte, line string, expected string) {
	t.Helper()
	n, err := p.parseLine(line, buf)
	if err != nil {
		t.Fatal(err)
	}
	got := hex.EncodeToString(buf[:n])
	if got != expected {
		t.Errorf("Unexpected parsed line result. Got %s, expected %s", got, expected)
	}
}

func TestParseLine(t *testing.T) {
	buf := make([]byte, 16)
	p := &Parser{}
	testLineHex(t, p, buf, "DRAWPIXEL 16 35 171,37,243", "0c00100023a93e")
	testLineHex(t, p, buf, "SETBITSIZE8", "00")
	testLineHex(t, p, buf, "DRAWPIXEL 16 35 171,37,243", "0c1023a93e")
	testLineHex(t, p, buf, "SETCOLOR 171,37,243", "09a93e")
	testLineHex(t, p, buf, "DRAWPIXEL 16 35", "0c1023")
	testLineHex(t, p, buf, "CLEARCOLOR", "0a")
	testLineHex(t, p, buf, "DRAWPIXEL 16 35 171,37,243", "0c1023a93e")
}
