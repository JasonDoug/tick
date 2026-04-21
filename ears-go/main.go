package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"charm.land/fantasy"
	"charm.land/fantasy/object"
	"charm.land/fantasy/providers/openai"
	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
)

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5")).MarginBottom(1)
	agentStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true)
	userStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	infoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	borderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)
)

func NewAssistantMessage(text string) fantasy.Message {
	return fantasy.Message{
		Role:    fantasy.MessageRoleAssistant,
		Content: []fantasy.MessagePart{fantasy.TextPart{Text: text}},
	}
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY is not set.")
		os.Exit(1)
	}

	ctx := context.Background()
	provider, err := openai.New(openai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	model, err := provider.LanguageModel(ctx, "gpt-4o")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(titleStyle.Render("Welcome to Tick: Agentic EARS Refinement Pipeline"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userStyle.Render("Seed Idea: "))
	seedIdea, _ := reader.ReadString('\n')
	seedIdea = strings.TrimSpace(seedIdea)

	// Phase 1: Charter Agent
	onePager, err := runCharterAgent(ctx, model, seedIdea, reader)
	if err != nil {
		log.Fatal(err)
	}

	// Phase 2: Architect Agent
	add, err := runArchitectAgent(ctx, model, onePager, reader)
	if err != nil {
		log.Fatal(err)
	}

	// Phase 3: Planning Agent
	plan, err := runPlanningAgent(ctx, model, onePager, add, reader)
	if err != nil {
		log.Fatal(err)
	}

	// Output final plan
	planJSON, _ := json.MarshalIndent(plan, "", "  ")
	fmt.Println("\n" + titleStyle.Render("Final Project Plan Generated!"))
	fmt.Println(string(planJSON))

	// Write artifacts to files for the Python layer to pick up
	os.WriteFile("one_pager.json", mustMarshal(onePager), 0644)
	os.WriteFile("add.json", mustMarshal(add), 0644)
	os.WriteFile("project_plan.json", planJSON, 0644)
}

func mustMarshal(v interface{}) []byte {
	b, _ := json.MarshalIndent(v, "", "  ")
	return b
}

func runCharterAgent(ctx context.Context, model fantasy.LanguageModel, seedIdea string, reader *bufio.Reader) (*OnePager, error) {
	fmt.Println("\n" + agentStyle.Render("Charter Agent is analyzing your seed idea..."))

	systemPrompt := `You are the Charter Agent in the Agentic EARS system.
Your sole responsibility is to transform a vague Seed Idea into a clear Project Charter / One-Pager.

Rules:
- Always propose one single NorthStar (one concise sentence).
- Decompose into 3-8 MVP Features (colloquial, user-visible capabilities).
- Include success metrics, scope boundaries, and high-level non-functional notes.
- Ask ONE clarifying question at a time. Never list questions.
- NEVER create EARS Requirements or Use Cases.
- Output the final One-Pager ONLY after the user says “APPROVE”, “FINALIZE”, “LOOKS GOOD”, or equivalent.`

	messages := []fantasy.Message{
		fantasy.NewUserMessage("Seed Idea: " + seedIdea),
	}

	agent := fantasy.NewAgent(model, fantasy.WithSystemPrompt(systemPrompt))

	for {
		result, err := agent.Generate(ctx, fantasy.AgentCall{Messages: messages})
		if err != nil {
			return nil, err
		}

		response := result.Response.Content.Text()
		fmt.Println("\n" + agentStyle.Render("Charter Agent: ") + response)

		// Check for approval
		if isApproved(response) {
			fmt.Println(infoStyle.Render("Generating structured One-Pager..."))
			obj, err := object.Generate[OnePager](ctx, model, fantasy.ObjectCall{
				Prompt: append(messages, NewAssistantMessage(response), fantasy.NewUserMessage("Generate the final One-Pager JSON now.")),
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}

		fmt.Print(userStyle.Render("You: "))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		messages = append(messages, NewAssistantMessage(response), fantasy.NewUserMessage(userInput))
		
		if isApproved(userInput) {
			fmt.Println(infoStyle.Render("Generating structured One-Pager..."))
			obj, err := object.Generate[OnePager](ctx, model, fantasy.ObjectCall{
				Prompt: messages,
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}
	}
}

func runArchitectAgent(ctx context.Context, model fantasy.LanguageModel, onePager *OnePager, reader *bufio.Reader) (*ADD, error) {
	fmt.Println("\n" + agentStyle.Render("Architect Agent is designing the technical foundation..."))

	systemPrompt := `You are the Architect Agent in the Agentic EARS system.
Your job is to transform the approved Project Charter / One-Pager into a complete Architectural & Design Document (ADD).

Rules:
- For each Feature in the One-Pager, create precise EARS-style Requirements.
- Expand those Requirements into detailed Use Cases.
- Define the technical architecture: modules, data models, APIs, tech stack, data flows, and constraints.
- Ask technical or design questions ONE AT A TIME when clarification is needed.
- Only output the full ADD after the user explicitly replies with “APPROVE DESIGN” or “FINALIZE ADD”.`

	onePagerJSON, _ := json.Marshal(onePager)
	messages := []fantasy.Message{
		fantasy.NewUserMessage("Approved One-Pager: " + string(onePagerJSON)),
	}

	agent := fantasy.NewAgent(model, fantasy.WithSystemPrompt(systemPrompt))

	for {
		result, err := agent.Generate(ctx, fantasy.AgentCall{Messages: messages})
		if err != nil {
			return nil, err
		}

		response := result.Response.Content.Text()
		fmt.Println("\n" + agentStyle.Render("Architect Agent: ") + response)

		if isApproved(response) || strings.Contains(strings.ToUpper(response), "APPROVE DESIGN") {
			fmt.Println(infoStyle.Render("Generating structured ADD..."))
			obj, err := object.Generate[ADD](ctx, model, fantasy.ObjectCall{
				Prompt: append(messages, NewAssistantMessage(response), fantasy.NewUserMessage("Generate the final ADD JSON now.")),
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}

		fmt.Print(userStyle.Render("You: "))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		messages = append(messages, NewAssistantMessage(response), fantasy.NewUserMessage(userInput))

		if isApproved(userInput) || strings.Contains(strings.ToUpper(userInput), "APPROVE DESIGN") {
			fmt.Println(infoStyle.Render("Generating structured ADD..."))
			obj, err := object.Generate[ADD](ctx, model, fantasy.ObjectCall{
				Prompt: messages,
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}
	}
}

func runPlanningAgent(ctx context.Context, model fantasy.LanguageModel, onePager *OnePager, add *ADD, reader *bufio.Reader) (*ProjectPlan, error) {
	fmt.Println("\n" + agentStyle.Render("Planning Agent is creating the roadmap..."))

	systemPrompt := `You are the Planning Agent in the Agentic EARS system.
Your job is to create a Project Plan from the approved One-Pager and ADD.

Rules:
- Group work logically into Sprints.
- Break work into Tasks and then into atomic Sub-tasks (focused prompts for AI coding tools).
- Since this is a TUI project, ensure there is a task specifically for "TUI Implementation" with a prompt that Termite can use.
- Only generate the final Project Plan after the user replies with “APPROVE PLAN” or “GENERATE ALL ARTIFACTS”.`

	onePagerJSON, _ := json.Marshal(onePager)
	addJSON, _ := json.Marshal(add)
	messages := []fantasy.Message{
		fantasy.NewUserMessage("Approved One-Pager: " + string(onePagerJSON)),
		fantasy.NewUserMessage("Approved ADD: " + string(addJSON)),
	}

	agent := fantasy.NewAgent(model, fantasy.WithSystemPrompt(systemPrompt))

	for {
		result, err := agent.Generate(ctx, fantasy.AgentCall{Messages: messages})
		if err != nil {
			return nil, err
		}

		response := result.Response.Content.Text()
		fmt.Println("\n" + agentStyle.Render("Planning Agent: ") + response)

		if isApproved(response) || strings.Contains(strings.ToUpper(response), "APPROVE PLAN") {
			fmt.Println(infoStyle.Render("Generating structured Project Plan..."))
			obj, err := object.Generate[ProjectPlan](ctx, model, fantasy.ObjectCall{
				Prompt: append(messages, NewAssistantMessage(response), fantasy.NewUserMessage("Generate the final Project Plan JSON now.")),
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}

		fmt.Print(userStyle.Render("You: "))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		messages = append(messages, NewAssistantMessage(response), fantasy.NewUserMessage(userInput))

		if isApproved(userInput) || strings.Contains(strings.ToUpper(userInput), "APPROVE PLAN") {
			fmt.Println(infoStyle.Render("Generating structured Project Plan..."))
			obj, err := object.Generate[ProjectPlan](ctx, model, fantasy.ObjectCall{
				Prompt: messages,
			})
			if err != nil {
				return nil, err
			}
			return &obj.Object, nil
		}
	}
}

func isApproved(text string) bool {
	upper := strings.ToUpper(text)
	return strings.Contains(upper, "APPROVE") ||
		strings.Contains(upper, "FINALIZE") ||
		strings.Contains(upper, "LOOKS GOOD")
}
