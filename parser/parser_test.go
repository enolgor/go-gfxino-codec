package parser

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

var input string = `
#this is a comment

DRAWPIXEL 16 35 171,37,243
SETBITSIZE8
DRAWPIXEL 16 35 171,37,243
SETCOLOR 171,37,243
DRAWPIXEL 16 35
CLEARCOLOR
DRAWPIXEL 16 35 171,37,243

`

var output string = `

0c00100023a93e
00
0c1023a93e
09a93e
0c1023
0a
0c1023a93e

`

func TestParse(t *testing.T) {
	output = strings.TrimSpace(output)
	output = strings.ReplaceAll(output, "\n", "")
	reader := strings.NewReader(input)
	p := &Parser{}
	var buf bytes.Buffer
	err := p.Parse(reader, &buf)
	if err != nil {
		t.Error(err)
	}
	got := hex.EncodeToString(buf.Bytes())
	if got != output {
		t.Errorf("Exepcted %s, got %s", output, got)
	}
}
