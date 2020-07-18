package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	mode8Bit  bool
	colorMode bool
}

func (p *Parser) Parse(reader io.Reader, writer io.Writer) error {
	var buf bytes.Buffer
	lineBuffer := make([]byte, 16)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		if line[0] == '#' {
			continue
		}
		n, err := p.parseLine(line, lineBuffer)
		if err != nil {
			return err
		}
		w, err := writer.Write(lineBuffer[:n])
		if err != nil {
			return err
		}
		if w != n {
			return fmt.Errorf("All bytes weren't written")
		}
		buf.Write(lineBuffer[:n])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
