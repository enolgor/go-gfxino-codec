package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/enolgor/go-gfxino-codec/color"
	"github.com/enolgor/go-gfxino-codec/commands"
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
	if len(args) != len(cmdArgs) {
		return -1, fmt.Errorf("Incorrect numbers of args for command %s. Expected %d, got %d", cmd, len(cmdArgs), len(args))
	}
	size := commands.GetArgSize(cmdByte, p.mode8Bit, p.colorMode) + 1
	buf[0] = byte(cmdByte)
	i := 1
	var err error
	for argIdx, argType := range cmdArgs {
		i, err = p.writeArg(i, buf, args[argIdx], argType)
		if err != nil {
			return -1, err
		}
	}
	if cmdByte == commands.SETBITSIZE8 {
		p.mode8Bit = true
	} else if cmdByte == commands.SETCOLOR {
		p.colorMode = true
	} else if cmdByte == commands.CLEARCOLOR {
		p.colorMode = false
	}
	return size, nil
}

func (p *Parser) writeArg(i int, buf []byte, val string, arg uint8) (j int, err error) {
	j = i
	switch arg {
	case commands.UINT8:
		if err = writeUint8(i, buf, val); err == nil {
			j++
		}
	case commands.UINT16:
		if err = writeUint16(i, buf, val); err == nil {
			j += 2
		}
	case commands.UINT:
		if p.mode8Bit {
			if err = writeUint8(i, buf, val); err == nil {
				j++
			}
		} else {
			if err = writeUint16(i, buf, val); err == nil {
				j += 2
			}
		}
	case commands.COLOR:
		if err = writeColor(i, buf, val); err == nil {
			j += 2
		}
	case commands.COLOR_SKIPPABLE:
		if !p.colorMode {
			if err = writeColor(i, buf, val); err == nil {
				j += 2
			}
		}
	}
	return
}

func writeUint8(i int, buf []byte, val string) error {
	v, err := strconv.ParseUint(val, 0, 8)
	if err != nil {
		return err
	}
	buf[i] = byte(v)
	return nil
}

func writeUint16(i int, buf []byte, val string) error {
	v, err := strconv.ParseUint(val, 0, 16)
	if err != nil {
		return err
	}
	buf[i] = byte((v >> 8) & 0xFF)
	buf[i+1] = byte(v & 0xFF)
	return nil
}

func writeColor(i int, buf []byte, val string) error {
	c, err := color.FromString(val)
	if err != nil {
		return err
	}
	u := c.To565()
	buf[i] = byte((u >> 8) & 0xFF)
	buf[i+1] = byte(u & 0xFF)
	return nil
}

func (p *Parser) getArgSize(args []uint8) int {
	s := 0
	for _, a := range args {
		switch a {
		case commands.UINT8:
			s++
		case commands.UINT16:
			s += 2
		case commands.UINT:
			if p.mode8Bit {
				s++
			} else {
				s += 2
			}
		case commands.COLOR:
			s += 2
		case commands.COLOR_SKIPPABLE:
			if p.colorMode {
				s += 0
			} else {
				s += 2
			}

		}
	}
	return s
}
