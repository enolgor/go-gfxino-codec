package interpreter

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"testing"
)

func testLineHex(t *testing.T, ip *Interpreter, buf []byte, hexLine string, expected string) {
	t.Helper()
	var w bytes.Buffer
	b, err := hex.DecodeString(hexLine)
	if err != nil {
		t.Fatal(err)
	}
	br := bufio.NewReader(bytes.NewReader(b))
	err = ip.writeLine(br, &w, buf)
	if err != nil {
		t.Fatal(err)
	}
	got := string(w.Bytes())
	if got != expected {
		t.Errorf("Unexpected parsed line result. Got %s, expected %s", got, expected)
	}
}

func TestWriteLine(t *testing.T) {
	buf := make([]byte, 16)
	ip := &Interpreter{}
	testLineHex(t, ip, buf, "0c00100023a93e", "DRAWPIXEL 16 35 168,36,240")
	testLineHex(t, ip, buf, "00", "SETBITSIZE8")
	testLineHex(t, ip, buf, "0c1023a93e", "DRAWPIXEL 16 35 168,36,240")
	testLineHex(t, ip, buf, "09a93e", "SETCOLOR 168,36,240")
	testLineHex(t, ip, buf, "0c1023", "DRAWPIXEL 16 35")
	testLineHex(t, ip, buf, "0a", "CLEARCOLOR")
	testLineHex(t, ip, buf, "0c1023a93e", "DRAWPIXEL 16 35 168,36,240")
}
