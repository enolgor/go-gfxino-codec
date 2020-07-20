package text

import (
	"bytes"

	"golang.org/x/text/encoding/charmap"
)

var cp437 = charmap.CodePage437

func EncodeCP437(s string) []byte {
	buf := bytes.Buffer{}
	buf.Grow(len(s))
	for _, r := range s {
		c, _ := cp437.EncodeRune(r) // add custom replacement if !ok ??
		buf.WriteByte(c)
	}
	return buf.Bytes()
}

func DecodeCP437(bslice []byte) string {
	buf := bytes.Buffer{}
	for _, b := range bslice {
		buf.WriteRune(cp437.DecodeByte(b))
	}
	return buf.String()
}
