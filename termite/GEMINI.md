# GEMINI.md

## Project Overview

**Termite** is an AI-powered CLI tool that generates functional Terminal User Interfaces (TUIs) from simple text prompts. It leverages Large Language Models (LLMs) to design, implement, and iteratively fix and refine Python-based TUI scripts.

### Core Technologies
- **Language:** Python 3.13+
- **LLM Providers:** Anthropic (Claude 3.5 Sonnet), OpenAI (GPT-4o), and Ollama.
- **Supported TUI Libraries:**
    - [urwid](https://urwid.org/) (Default)
    - [rich](https://rich.readthedocs.io/)
    - [textual](https://textual.textualize.io/)
    - [curses](https://docs.python.org/3/library/curses.html)
- **CLI Feedback:** [rich](https://rich.readthedocs.io/) for progress bars and logging.

### Architecture
The project is structured into a main package `termite/` with the following components:
- `termite/termite.py`: Orchestrates the generation pipeline.
- `termite/tools/`: Contains the specialized modules for each step of the pipeline:
    - `design_tui.py`: Generates the design document from the user prompt.
    - `build_tui.py`: Implements the TUI code based on the design.
    - `fix_errors.py`: Iteratively executes the code in a pseudo-terminal and fixes runtime errors.
    - `refine.py`: (Optional) Self-reflection and refinement loop.
- `termite/shared/`: Utility modules for LLM interaction (`call_llm.py`), TUI execution (`run_tui.py`), and system-level helpers.
- `termite/dtos/`: Data Transfer Objects for configuration (`Config.py`) and script state (`Script.py`).

---

## Building and Running

### Installation
Termite uses `uv` for project and dependency management. To install it locally:
```bash
uv pip install -e .
```
Or run it directly with `uv`:
```bash
uv run termite "Make me a system monitor"
```

### Environment Setup
You must provide an API key for your preferred LLM provider:
```bash
export ANTHROPIC_API_KEY="..." # Recommended
export OPENAI_API_KEY="..."
# For Ollama:
export OLLAMA_MODEL="llama3"
```

### Execution
Run the CLI using the `termite` command or via the module:
```bash
termite "Make me a system monitor"
# OR
python3 -m termite
```

### Key Arguments
- `--library`: Choose the TUI library (`urwid`, `rich`, `textual`, `curses`).
- `--refine`: Enable the refinement loop for higher quality output.
- `--fix-iters`: Maximum attempts to fix runtime errors (Default: 10).
- `--run-tool <name>`: Run a previously generated and saved TUI.

---

## Development Conventions

### Code Style
- **Type Hinting:** Use type hints for function signatures and data structures.
- **DTOs:** Use `dataclasses` for structured data (see `termite/dtos/`).
- **Error Handling:** LLM prompts explicitly forbid `try/except` blocks in generated code to ensure all runtime errors are captured by the `fix_errors` tool during development.
- **LLM Interaction:** All LLM calls should go through `termite/shared/call_llm.py` to maintain provider abstraction.

### Testing
- **TODO:** There are currently no automated unit or integration tests in the repository. Testing is primarily done manually by verifying the generated TUIs.

### TUI Generation Lifecycle
When modifying the generation logic, update the corresponding tool in `termite/tools/`. Prompts are defined as constants within these files.
- Design -> Implementation -> Error Correction -> (Optional) Refinement.
- Generated scripts are saved to `~/.termite/` (or `$XDG_CONFIG_HOME/termite/`).
