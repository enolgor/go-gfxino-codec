package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/enolgor/go-gfxino-codec/interpreter"
	"github.com/enolgor/go-gfxino-codec/parser"
)

const usageMessage = `
Usage: gfxinocodec <command> [<args>]
Available commands:
  encode      Encode file
  decode      Decode file
`

var encodeCommand = flag.NewFlagSet("encode", flag.ExitOnError)
var decodeCommand = flag.NewFlagSet("decode", flag.ExitOnError)

var inputfile string

const (
	inputfileFlag        = "i"
	inputfileFlagDefault = ""
	inputfileFlagUsage   = "Input file"
)

var outputfile string

const (
	outputfileFlag        = "o"
	outputfileFlagDefault = ""
	outputfileFlagUsage   = "Output file or stdout if missing"
)

var colorModeHex bool

const (
	colorModeHexFlag        = "h"
	colorModeHexFlagDefault = false
	colorModeHexFlagUsage   = "Color output mode as hex"
)

var doubleLine bool

const (
	doubleLineFlag        = "d"
	doubleLineFlagDefault = false
	doubleLineFlagUsage   = "Double line between instructions instead of 1"
)

func init() {
	encodeCommand.StringVar(&inputfile, inputfileFlag, inputfileFlagDefault, inputfileFlagUsage)
	encodeCommand.StringVar(&outputfile, outputfileFlag, outputfileFlagDefault, outputfileFlagUsage)

	decodeCommand.StringVar(&inputfile, inputfileFlag, inputfileFlagDefault, inputfileFlagUsage)
	decodeCommand.StringVar(&outputfile, outputfileFlag, outputfileFlagDefault, outputfileFlagUsage)
	decodeCommand.BoolVar(&colorModeHex, colorModeHexFlag, colorModeHexFlagDefault, colorModeHexFlagUsage)
	decodeCommand.BoolVar(&doubleLine, doubleLineFlag, doubleLineFlagDefault, doubleLineFlagUsage)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, usageMessage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "encode":
		encodeCommand.Parse(os.Args[2:])
		runEncodeCommand()
	case "decode":
		decodeCommand.Parse(os.Args[2:])
		runDecodeCommand()
	default:
		fmt.Fprintf(os.Stderr, "%q is not a valid command\n", os.Args[1])
		fmt.Fprint(os.Stderr, usageMessage)
		os.Exit(1)
	}
}

func printDefaults(flagSet *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "\n%s command usage:\n\n", flagSet.Name())
	flagSet.PrintDefaults()
}

func printErrorAndExit(e error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", e.Error())
	os.Exit(1)
}

func runEncodeCommand() {
	if inputfile == "" {
		fmt.Fprintln(os.Stderr, "Input file must be specified")
		printDefaults(encodeCommand)
		os.Exit(1)
	}
	fi, err := os.Open(inputfile)
	defer fi.Close()
	if err != nil {
		printErrorAndExit(err)
	}
	var w io.Writer
	if outputfile != "" {
		fo, err := os.OpenFile(outputfile, os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer fo.Close()
		if err != nil {
			printErrorAndExit(err)
		}
		w = fo
	} else {
		w = os.Stdout
	}
	p := &parser.Parser{}
	if err = p.Parse(fi, w); err != nil {
		printErrorAndExit(err)
	}
}

func runDecodeCommand() {
	if inputfile == "" {
		fmt.Fprintln(os.Stderr, "Input file must be specified")
		printDefaults(decodeCommand)
		os.Exit(1)
	}
	fi, err := os.Open(inputfile)
	defer fi.Close()
	if err != nil {
		printErrorAndExit(err)
	}
	var w io.Writer
	if outputfile != "" {
		fo, err := os.OpenFile(outputfile, os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer fo.Close()
		if err != nil {
			printErrorAndExit(err)
		}
		w = fo
	} else {
		w = os.Stdout
	}
	ip := &interpreter.Interpreter{}
	if colorModeHex {
		ip.WriteColorMode = interpreter.COLOR_HEX
	}
	if doubleLine {
		ip.DoubleLine = true
	}
	if err = ip.Interpret(fi, w); err != nil {
		printErrorAndExit(err)
	}

}
