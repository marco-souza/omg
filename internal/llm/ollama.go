package llm

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

//go:embed role
var role string

func Completion(prompt string) {
	model, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	readStream(model, prompt)
}

func processPrompt(prompt string) string {
	return fmt.Sprintf(role, prompt)
}

func readStream(llm *ollama.LLM, query string) {
	ctx := context.Background()
  prompt := processPrompt(query)

	_, err := llm.Call(ctx, prompt, llms.WithStreamingFunc(handleChunks))
	if err != nil {
		log.Fatal(err)
	}
}

func handleChunks(ctx context.Context, chunk []byte) error {
	fmt.Printf("%s", chunk)
	// fmt.Printf("chunk len=%d: %s\n", len(chunk), chunk)
	return nil
}
