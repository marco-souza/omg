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
	input := cli.Parse()
	prompt := strings.Join(input.Args, " ")

	if len(input.Args) == 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		prompt = string(stdin)
	}

	llm.Completion(prompt, input.OutputPath)
}
