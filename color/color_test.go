package color

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	c1 := Color{R: 171, G: 37, B: 243}
	c2 := "0xAB25F3"
	c3 := "171,37,243"
	if c1.ToHexString() != c2 {
		t.Error("c1 to hexstring != c2")
	}
	c, e := FromString(c2)
	if e != nil {
		t.Error(e)
	}
	if !reflect.DeepEqual(*c, c1) {
		t.Error("c2 from string != c1")
	}
	c, e = FromString(c3)
	if e != nil {
		t.Error(e)
	}
	if !reflect.DeepEqual(*c, c1) {
		t.Error("c3 from string != c1")
	}
}

func Test565(t *testing.T) {
	c1 := Color{R: 171, G: 37, B: 243}
	var c2 uint16 = 0xA93E
	c3 := Color{R: c1.R & 0xF8, G: c1.G & 0xFC, B: c1.B & 0xF8}
	if c1.To565() != c2 {
		t.Errorf("c1 to 565 != c2")
	}
	fmt.Println(*From565(c2))
	if !reflect.DeepEqual(*From565(c2), c3) {
		t.Errorf("c2 from 565 != c3")
	}
}
