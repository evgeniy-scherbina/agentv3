package main

import (
	"context"
	"fmt"
	"os"

	"charm.land/fantasy"
	"charm.land/fantasy/providers/anthropic"
)

func main() {
	apiKey := os.Getenv("API_KEY")

	// Choose your fave provider.
	//provider, err := openrouter.New(openrouter.WithAPIKey(apiKey))
	provider, err := anthropic.New(anthropic.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Whoops:", err)
		os.Exit(1)
	}

	ctx := context.Background()

	// Pick your fave model.
	model, err := provider.LanguageModel(ctx, "claude-sonnet-4-5")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dang:", err)
		os.Exit(1)
	}

	// Make your own tools.
	//cuteDogTool := fantasy.NewAgentTool(
	//	"cute_dog_tool",
	//	"Provide up-to-date info on cute dogs.",
	//	fetchCuteDogInfoFunc,
	//)

	// Equip your agent.
	agent := fantasy.NewAgent(
		model,
		fantasy.WithSystemPrompt("You are a moderately helpful, dog-centric assistant."),
		//fantasy.WithTools(cuteDogTool),
	)

	// Put that agent to work!
	const prompt = "Find all the cute dogs in Silver Lake, Los Angeles."
	result, err := agent.Generate(ctx, fantasy.AgentCall{Prompt: prompt})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oof:", err)
		os.Exit(1)
	}
	fmt.Println(result.Response.Content.Text())
}
