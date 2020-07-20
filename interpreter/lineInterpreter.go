package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/enolgor/go-gfxino-codec/color"
	"github.com/enolgor/go-gfxino-codec/commands"
	"github.com/enolgor/go-gfxino-codec/text"
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
	if commands.IsTextCommand(cmd) {
		n, err := peekTextLength(cmd, br)
		if err != nil {
			return err
		}
		argSize += n + 1
	}
	b := buf[:argSize]
	if _, err = io.ReadFull(br, b); err != nil {
		return err
	}
	words := make([]string, len(cmdArgs)+1)
	words[0] = cmdStr
	i := 0
	for argIdx, argType := range cmdArgs {
		s, str, err := ip.readNext(i, b, argType)
		if err != nil {
			return err
		}
		words[argIdx+1] = str
		i += s
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

func (ip *Interpreter) readNext(i int, buf []byte, argType uint8) (int, string, error) {
	switch argType {
	case commands.UINT8:
		return readNextUint8(i, buf)
	case commands.UINT16:
		return readNextUint8(i, buf)
	case commands.UINT:
		return readNextUint(i, buf, ip.mode8Bit)
	case commands.COLOR:
		return readNextColor(i, buf, ip.WriteColorMode)
	case commands.COLOR_SKIPPABLE:
		return readNextSkippableColor(i, buf, ip.WriteColorMode, ip.colorMode)
	case commands.TEXT:
		return readNextText(i, buf)
	}
	return -1, "", fmt.Errorf("Command not found")
}

func readNextUint(i int, buf []byte, mode8Bit bool) (int, string, error) {
	if mode8Bit {
		return readNextUint8(i, buf)
	}
	return readNextUint16(i, buf)
}

func readNextUint16(i int, buf []byte) (int, string, error) {
	var v uint16
	if len(buf)-i < 2 {
		return -1, "", fmt.Errorf("Buf rem size smaller than 2")
	}
	v = uint16(buf[i])<<8 | uint16(buf[i+1])
	return 2, fmt.Sprintf("%d", v), nil
}

func readNextUint8(i int, buf []byte) (int, string, error) {
	var v uint8
	if len(buf)-i < 1 {
		return -1, "", fmt.Errorf("Buf rem size smaller than 1")
	}
	v = uint8(buf[i])
	return 1, fmt.Sprintf("%d", v), nil
}

func readNextInt16(i int, buf []byte) (int, string, error) {
	var v int16
	if len(buf)-i < 2 {
		return -1, "", fmt.Errorf("Buf rem size smaller than 2")
	}
	v = int16(buf[i])<<8 | int16(buf[i+1])
	return 2, fmt.Sprintf("%d", v), nil
}

func readNextSkippableColor(i int, buf []byte, writeColorMode int, colorMode bool) (int, string, error) {
	if !colorMode {
		return readNextColor(i, buf, writeColorMode)
	}
	return 0, "", nil
}

func readNextColor(i int, buf []byte, writeColorMode int) (int, string, error) {
	var v uint16
	if len(buf)-i < 2 {
		return -1, "", fmt.Errorf("Buf rem size smaller than 2")
	}
	v = uint16(buf[i])<<8 | uint16(buf[i+1])
	c := color.From565(v)
	switch writeColorMode {
	case COLOR_RGB:
		return 2, c.ToRGBString(), nil
	case COLOR_HEX:
		return 2, c.ToHexString(), nil
	default:
		return 2, c.ToRGBString(), nil
	}
}

func readNextText(i int, buf []byte) (int, string, error) {
	var s int
	if len(buf)-i < 1 {
		return -1, "", fmt.Errorf("Buf rem size smaller than 1")
	}
	s = int(buf[i])
	if len(buf)-i < 1+s {
		return -1, "", fmt.Errorf("Buf rem size smaller than text size")
	}
	return s + 1, text.DecodeCP437(buf[i+1 : i+s+1]), nil
}

func peekTextLength(cmd byte, br *bufio.Reader) (int, error) {
	switch cmd {
	case commands.PRINT:
		b, err := br.Peek(1)
		if err != nil {
			return -1, err
		}
		return int(b[0]), nil
	}
	return -1, fmt.Errorf("Not a text command")
}
