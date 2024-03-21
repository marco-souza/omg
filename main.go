package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/marco-souza/omg/cmd/cli"
	"github.com/marco-souza/omg/internal/llm"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		prompt := strings.Join(argsWithoutProg, " ")
		llm.Completion(prompt)
		return
	}

	p := tea.NewProgram(cli.MakeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
