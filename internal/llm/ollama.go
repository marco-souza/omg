package llm

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

func Completion(prompt string) {
	model, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	role := "ROLE: You will act as an AI prompt generator. I'll pass you some instructions, and " + "you will asnwer me with a good prompt to feed an AI assistent"
	prompt = fmt.Sprintf("%s\n\n%s", role, prompt)

	ctx := context.Background()
	completion, err := model.Call(ctx, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}
