package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/enolgor/go-gfxino-codec/color"
	"github.com/enolgor/go-gfxino-codec/commands"
	"github.com/enolgor/go-gfxino-codec/text"
)

func (p *Parser) parseLine(line string, buf []byte) (int, error) {
	parts := strings.Split(line, " ")
	cmd := parts[0]
	args := parts[1:]
	cmdByte, ok := commands.InstructionReverseMap[cmd]
	if !ok {
		return -1, fmt.Errorf("Command %s not found", cmd)
	}
	cmdArgs := commands.CommandArgsMap[cmdByte]
	if p.colorMode && len(cmdArgs) > 0 && cmdArgs[len(cmdArgs)-1] == commands.COLOR_SKIPPABLE {
		cmdArgs = cmdArgs[:len(cmdArgs)-1]
	}
	if commands.IsTextCommand(cmdByte) {
		args = append(args[:len(cmdArgs)-1], strings.Join(args[len(cmdArgs)-1:], " "))
	}
	if len(args) != len(cmdArgs) {
		return -1, fmt.Errorf("Incorrect numbers of args for command %s. Expected %d, got %d", cmd, len(cmdArgs), len(args))
	}
	if commands.IsTextCommand(cmdByte) {
		args = append(args[:len(cmdArgs)-1], strings.Join(args[len(cmdArgs)-1:], " "))
	}
	buf[0] = byte(cmdByte)
	i := 1
	for argIdx, argType := range cmdArgs {
		s, err := p.writeArg(i, buf, args[argIdx], argType)
		if err != nil {
			return -1, err
		}
		i += s
	}
	if cmdByte == commands.SETBITSIZE8 {
		p.mode8Bit = true
	} else if cmdByte == commands.SETCOLOR {
		p.colorMode = true
	} else if cmdByte == commands.CLEARCOLOR {
		p.colorMode = false
	}
	return i, nil
}

func (p *Parser) writeArg(i int, buf []byte, val string, arg uint8) (int, error) {
	switch arg {
	case commands.UINT8:
		return writeUint8(i, buf, val)
	case commands.UINT16:
		return writeUint16(i, buf, val)
	case commands.UINT:
		return writeUint(i, buf, val, p.mode8Bit)
	case commands.COLOR:
		return writeColor(i, buf, val)
	case commands.COLOR_SKIPPABLE:
		return writeSkippableColor(i, buf, val, p.colorMode)
	case commands.TEXT:
		return writeText(i, buf, val)
	}
	return -1, fmt.Errorf("Command not found")
}

func writeUint(i int, buf []byte, val string, mode8Bit bool) (int, error) {
	if mode8Bit {
		return writeUint8(i, buf, val)
	}
	return writeUint16(i, buf, val)
}

func writeUint8(i int, buf []byte, val string) (int, error) {
	v, err := strconv.ParseUint(val, 0, 8)
	if err != nil {
		return -1, err
	}
	buf[i] = byte(v)
	return 1, nil
}

func writeUint16(i int, buf []byte, val string) (int, error) {
	v, err := strconv.ParseUint(val, 0, 16)
	if err != nil {
		return -1, err
	}
	buf[i] = byte((v >> 8) & 0xFF)
	buf[i+1] = byte(v & 0xFF)
	return 2, nil
}

func writeSkippableColor(i int, buf []byte, val string, colorMode bool) (int, error) {
	if !colorMode {
		return writeColor(i, buf, val)
	}
	return 0, nil
}

func writeColor(i int, buf []byte, val string) (int, error) {
	c, err := color.FromString(val)
	if err != nil {
		return -1, err
	}
	u := c.To565()
	buf[i] = byte((u >> 8) & 0xFF)
	buf[i+1] = byte(u & 0xFF)
	return 2, nil
}

func writeText(i int, buf []byte, val string) (int, error) {
	encoded := text.EncodeCP437(val)
	if len(encoded) > 255 {
		encoded = encoded[:255]
	}
	buf[i] = byte(len(encoded))
	copy(buf[i+1:len(encoded)+1+i], encoded)
	return len(encoded) + 1, nil
}
