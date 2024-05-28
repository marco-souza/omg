package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/marco-souza/omg/cmd/cli"
	"github.com/marco-souza/omg/internal/llm"
)

func main() {
	c := cli.NewCli()
	c.Parse()

	prompt := strings.Join(c.Args, " ")

	if *c.ShowHelp {
		c.Usage()
		os.Exit(0)
	}

	if len(c.Args) == 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		prompt = string(stdin)
	}

	llm.Completion(prompt, *c.OutputPath)
}
