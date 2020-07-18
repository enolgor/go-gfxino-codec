package interpreter

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

var hexInput string = `

0c00100023a93e
00
0c1023a93e
09a93e
0c1023
0a
0c1023a93e

`

var output string = `

DRAWPIXEL 16 35 168,36,240
SETBITSIZE8
DRAWPIXEL 16 35 168,36,240
SETCOLOR 168,36,240
DRAWPIXEL 16 35
CLEARCOLOR
DRAWPIXEL 16 35 168,36,240

`

func TestInterpret(t *testing.T) {
	output = strings.TrimSpace(output)
	hexInput = strings.TrimSpace(hexInput)
	hexInput = strings.ReplaceAll(hexInput, "\n", "")
	input, err := hex.DecodeString(hexInput)
	if err != nil {
		t.Error(err)
	}
	ip := &Interpreter{}
	r := bytes.NewReader(input)
	var w bytes.Buffer
	err = ip.Interpret(r, &w)
	if err != nil {
		t.Error(err)
	}
	got := string(w.Bytes())
	got = strings.TrimSpace(got)
	if got != output {
		t.Errorf("Expected \n%s\n\ngot \n%s", output, got)
	}
}
