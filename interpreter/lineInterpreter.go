package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/enolgor/go-gfxino-codec/color"
	"github.com/enolgor/go-gfxino-codec/commands"
)

var errBufferEmtpy error = fmt.Errorf("Buffer is empty")

func (ip *Interpreter) writeLine(br *bufio.Reader, w io.Writer, buf []byte) error {
	var cmd uint8
	var err error
	if cmd, err = br.ReadByte(); err != nil {
		return errBufferEmtpy
	}
	cmdStr, ok := commands.InstructionForwardMap[cmd]
	if !ok {
		return fmt.Errorf("Command with bytecode %X not found", cmd)
	}
	cmdArgs := commands.CommandArgsMap[cmd]
	if ip.colorMode && len(cmdArgs) > 0 && cmdArgs[len(cmdArgs)-1] == commands.COLOR_SKIPPABLE {
		cmdArgs = cmdArgs[:len(cmdArgs)-1]
	}
	argSize := commands.GetArgSize(cmd, ip.mode8Bit, ip.colorMode)
	b := buf[:argSize]
	if _, err = io.ReadFull(br, b); err != nil {
		return err
	}
	words := make([]string, len(cmdArgs)+1)
	words[0] = cmdStr
	var s string
	i := 0
	for argIdx, argType := range cmdArgs {
		s, i, err = ip.readNext(i, b, argType)
		if err != nil {
			return err
		}
		words[argIdx+1] = s
	}
	if cmd == commands.SETBITSIZE8 {
		ip.mode8Bit = true
	} else if cmd == commands.SETCOLOR {
		ip.colorMode = true
	} else if cmd == commands.CLEARCOLOR {
		ip.colorMode = false
	}
	if _, err := fmt.Fprint(w, strings.Join(words, " ")); err != nil {
		return err
	}
	return nil
}

func (ip *Interpreter) readNext(i int, buf []byte, argType uint8) (s string, j int, err error) {
	j = i
	switch argType {
	case commands.UINT8:
		if s, err = readNextUint8(i, buf); err == nil {
			j++
		}
	case commands.UINT16:
		if s, err = readNextUint8(i, buf); err == nil {
			j += 2
		}
	case commands.UINT:
		if ip.mode8Bit {
			if s, err = readNextUint8(i, buf); err == nil {
				j++
			}
		} else {
			if s, err = readNextUint16(i, buf); err == nil {
				j += 2
			}
		}
	case commands.COLOR:
		if s, err = readNextColor(i, buf, ip.WriteColorMode); err == nil {
			j += 2
		}
	case commands.COLOR_SKIPPABLE:
		if !ip.colorMode {
			if s, err = readNextColor(i, buf, ip.WriteColorMode); err == nil {
				j += 2
			}
		}
	}
	return
}

func readNextUint16(i int, buf []byte) (string, error) {
	var v uint16
	if len(buf)-i < 2 {
		return "", fmt.Errorf("Buf rem size smaller than 2")
	}
	v = uint16(buf[i])<<8 | uint16(buf[i+1])
	return fmt.Sprintf("%d", v), nil
}

func readNextUint8(i int, buf []byte) (string, error) {
	var v uint8
	if len(buf)-i < 1 {
		return "", fmt.Errorf("Buf rem size smaller than 1")
	}
	v = uint8(buf[i])
	return fmt.Sprintf("%d", v), nil
}

func readNextColor(i int, buf []byte, writeColorMode int) (string, error) {
	var v uint16
	if len(buf)-i < 2 {
		return "", fmt.Errorf("Buf rem size smaller than 2")
	}
	v = uint16(buf[i])<<8 | uint16(buf[i+1])
	c := color.From565(v)
	switch writeColorMode {
	case COLOR_RGB:
		return c.ToRGBString(), nil
	case COLOR_HEX:
		return c.ToHexString(), nil
	default:
		return c.ToRGBString(), nil
	}
}
