package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"charm.land/fantasy"
	"charm.land/fantasy/providers/anthropic"
	"github.com/evgeniy-scherbina/agentv3/internal/agent/tools"
)

func main() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	modelName := "claude-sonnet-4-5"

	// Choose your fave provider.
	//provider, err := openrouter.New(openrouter.WithAPIKey(apiKey))
	provider, err := anthropic.New(anthropic.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Whoops:", err)
		os.Exit(1)
	}

	ctx := context.Background()

	// Pick your fave model.
	model, err := provider.LanguageModel(ctx, modelName)
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

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	bashTool := tools.NewBashTool(wd, modelName)

	// Equip your agent.
	agent := fantasy.NewAgent(
		model,
		fantasy.WithSystemPrompt("You are a moderately helpful, code-centric assistant."),
		fantasy.WithTools(bashTool),
	)

	// Put that agent to work!
	const prompt = "could you develop coder website for me"
	maxOutputTokens := int64(10_000)
	result, err := agent.Generate(ctx, fantasy.AgentCall{
		Prompt:          prompt,
		MaxOutputTokens: &maxOutputTokens,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oof:", err)
		os.Exit(1)
	}
	fmt.Println("text", result.Response.Content.Text())

	resultInJSON, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resultInJSON: %s\n", resultInJSON)
}
