package interpreter

import (
	"bufio"
	"fmt"
	"io"
)

const (
	COLOR_RGB = iota
	COLOR_HEX
)

type Interpreter struct {
	DoubleLine     bool
	WriteColorMode int
	mode8Bit       bool
	colorMode      bool
}

func (ip *Interpreter) Interpret(r io.Reader, w io.Writer) error {
	br := bufio.NewReader(r)
	lineBuffer := make([]byte, 256)
	for {
		if err := ip.writeLine(br, w, lineBuffer); err != nil {
			if err == errBufferEmtpy {
				break
			} else {
				return err
			}
		}
		fmt.Fprintln(w)
		if ip.DoubleLine {
			fmt.Fprintln(w)
		}
	}
	return nil
}
