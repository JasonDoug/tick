# GEMINI.md - Tick Workspace

## Overview
**Tick** is a unified AI-powered development environment that combines high-level project refinement with low-level TUI generation. It uses a Go-based interactive pipeline for requirement engineering and a Python-based engine for implementing the final Terminal User Interface.

## Architecture
The project is a hybrid Go/Python system:
- **`ears-go/`**: A Go-based CLI built with the `fantasy` framework. It implements the **Agentic EARS** methodology (Charter -> Architect -> Planning) to refine a "Seed Idea" into a structured `Project Plan`.
- **`tick/`**: A Python package containing the orchestrator and the **Termite** TUI generation engine.
- **`fantasy/`**: The underlying Go agentic framework (local copy).
- **`termite/`**: The original TUI generation code (preserved for reference, source moved to `tick/termite`).

## Workflow
1.  **Seed Idea**: User provides a vague idea for a TUI.
2.  **Interactive Refinement (Go)**:
    - **Charter Agent**: Proposes NorthStar and MVP Features.
    - **Architect Agent**: Creates EARS Requirements and Use Cases.
    - **Planning Agent**: Generates Sprints, Tasks, and atomic Prompts.
3.  **Handoff**: The Go pipeline outputs validated JSON artifacts.
4.  **TUI Generation (Python)**:
    - Termite takes the specifications and uses its Design -> Build -> Fix -> Refine loop to generate a functional Python TUI script.

## Getting Started

### Prerequisites
- Go 1.26.2+
- Python 3.13+
- API Keys: `OPENAI_API_KEY` or `ANTHROPIC_API_KEY`.

### Installation
```bash
# Build the Go refinement engine
cd ears-go && go build -o tick-ears

# Install the Python orchestrator
cd ..
pip install -e .
```

### Usage
```bash
# Run the full Tick pipeline
python3 -m tick "I want a system monitor with live graphs"
```

## Development Conventions
- **Agents**: EARS agents are defined in `ears-go/main.go` using the `fantasy` API.
- **TUI Engine**: Generation logic resides in `tick/termite/`.
- **Shared LLM**: Both systems use a unified approach to LLM calls (centralized in `tick/shared/call_llm.py` for Python and via `fantasy` for Go).
- **Artifacts**: All intermediary documents (One-Pager, ADD, Project Plan) are stored as JSON files in the workspace during execution.
