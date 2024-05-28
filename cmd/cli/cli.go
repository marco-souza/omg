package cli

import (
	"flag"
	"fmt"
	"os"
)

type CliInput struct {
	OutputPath *string
	ShowHelp   *bool
	Args       []string
}

func NewCli() *CliInput {
	// define cli flags
	outputPath := flag.String("o", "", "")
	flag.StringVar(outputPath, "output", "", "output file")

	help := flag.Bool("h", false, "")
	flag.BoolVar(help, "help", false, "show help")

	return &CliInput{
		outputPath, help, []string{},
	}
}

func (c *CliInput) Parse() {
	flag.Parse()
	c.Args = flag.Args()
}

func (c *CliInput) Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	flag.PrintDefaults()
}
