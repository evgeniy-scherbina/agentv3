package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	result, err := agent.Stream(ctx, fantasy.AgentStreamCall{
		Prompt: prompt,
		//Messages
		MaxOutputTokens: &maxOutputTokens,
		PrepareStep: func(callContext context.Context, options fantasy.PrepareStepFunctionOptions) (_ context.Context, prepared fantasy.PrepareStepResult, err error) {
			fmt.Printf("PrepareStep BEGIN\n")
			optionsInJSON, err := json.Marshal(options)
			if err != nil {
				return callContext, prepared, err
			}
			fmt.Printf("optionsInJSON: %s\n", optionsInJSON)
			fmt.Printf("PrepareStep END\n\n\n")
			return callContext, prepared, nil
		},
		OnReasoningStart: func(id string, reasoning fantasy.ReasoningContent) error {
			fmt.Printf("OnReasoningStart BEGIN\n")
			reasoningInJSON, err := json.Marshal(reasoning)
			if err != nil {
				return err
			}
			fmt.Printf("id: %s\n", id)
			fmt.Printf("reasoningInJSON: %s\n", reasoningInJSON)
			fmt.Printf("OnReasoningStart END\n\n\n")
			return nil
		},
		OnReasoningDelta: func(id string, text string) error {
			fmt.Printf("OnReasoningDelta BEGIN\n")
			fmt.Printf("id: %s\n", id)
			fmt.Printf("text: %s\n", text)
			fmt.Printf("OnReasoningDelta END\n\n\n")
			return nil
		},
		OnReasoningEnd: func(id string, reasoning fantasy.ReasoningContent) error {
			fmt.Printf("OnReasoningEnd BEGIN\n")
			reasoningInJSON, err := json.Marshal(reasoning)
			if err != nil {
				return err
			}
			fmt.Printf("id: %s\n", id)
			fmt.Printf("reasoningInJSON: %s\n", reasoningInJSON)
			fmt.Printf("OnReasoningEnd END\n\n\n")
			return nil
		},
		OnTextDelta: func(id string, text string) error {
			fmt.Printf("OnTextDelta BEGIN\n")
			fmt.Printf("id: %s\n", id)
			fmt.Printf("text: %s\n", text)
			fmt.Printf("OnTextDelta END\n\n\n")
			return nil
		},
		OnToolInputStart: func(id string, toolName string) error {
			fmt.Printf("OnToolInputStart BEGIN\n")
			fmt.Printf("id: %s\n", id)
			fmt.Printf("toolName: %s\n", toolName)
			fmt.Printf("OnToolInputStart END\n\n\n")
			return nil
		},
		OnToolCall: func(tc fantasy.ToolCallContent) error {
			fmt.Printf("OnToolCall BEGIN\n")
			toolCallContentInJSON, err := json.Marshal(tc)
			if err != nil {
				return err
			}
			fmt.Printf("toolCallContentInJSON: %s\n", toolCallContentInJSON)
			fmt.Printf("OnToolCall END\n\n\n")
			return nil
		},
		OnToolResult: func(result fantasy.ToolResultContent) error {
			fmt.Printf("OnToolResult BEGIN\n")
			resultInJSON, err := json.Marshal(result)
			if err != nil {
				return err
			}
			fmt.Printf("resultInJSON: %s\n", resultInJSON)
			fmt.Printf("OnToolResult END\n\n\n")
			return nil
		},
		OnStepFinish: func(stepResult fantasy.StepResult) error {
			fmt.Printf("OnStepFinish BEGIN\n")
			stepResultInJSON, err := json.Marshal(stepResult)
			if err != nil {
				return err
			}
			fmt.Printf("stepResultInJSON: %s\n", stepResultInJSON)
			fmt.Printf("OnStepFinish END\n\n\n")
			return nil
		},
		StopWhen: []fantasy.StopCondition{},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oof:", err)
		os.Exit(1)
	}
	fmt.Println("text", result.Response.Content.Text())

	//resultInJSON, err := json.Marshal(result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("resultInJSON: %s\n", resultInJSON)
}
