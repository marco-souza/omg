package cli

import "flag"

type CliInput struct {
	OutputPath string
	Args       []string
}

func Parse() *CliInput {
	// setup input
	outputPath := flag.String("o", "", "output file")
	flag.Parse()

	return &CliInput{
		*outputPath, flag.Args(),
	}
}
