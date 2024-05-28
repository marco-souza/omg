package llm

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/marco-souza/omg/internal/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var settings = config.NewSettings()

func Completion(prompt string, output string) {
	model, err := ollama.New(ollama.WithModel(settings.Model))
	if err != nil {
		log.Fatal("failed to connect to ollama: ", err)
	}

	outputFile := os.Stdout
	if output != "" {
		outputFile, err = os.Create(output)
		if err != nil {
			log.Fatal("failed to create output file: ", output, err)
		}
		defer outputFile.Close()
	}

	m := &completionState{prompt, outputFile}
	m.completeWithModel(model)
}

//go:embed role
var role string

type completionState struct {
	prompt     string
	outputFile *os.File
}

func (m *completionState) completeWithModel(llm *ollama.LLM) {
	ctx := context.Background()
	prompt := fmt.Sprintf(role, m.prompt)

	_, err := llm.Call(ctx, prompt, llms.WithStreamingFunc(m.handleChunks))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *completionState) handleChunks(ctx context.Context, chunk []byte) error {
	fmt.Fprintf(m.outputFile, "%s", chunk)
	return nil
}
