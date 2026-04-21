# Third party
from rich.progress import Progress

# Local
from tick.termite.dtos import Config
from tick.shared.call_llm import call_llm, MAX_TOKENS


#########
# HELPERS
#########


PROGRESS_LIMIT = MAX_TOKENS // 15
PROMPT = """You are an expert TUI designer.
You will be given a prompt that describes a TUI. Your job is to generate a design document for that TUI.
The design document should be written in Markdown and should be as detailed as possible.

You MUST follow these rules at all times:
- Use ONLY the {library} library to build the TUI. Do NOT use any other TUI libraries.
- The design document must include a list of requirements for the TUI.
- Use ONLY Markdown. Do NOT use any other formatting.

Respond with the complete design document without any explanations or other formatting."""


######
# MAIN
######


def design_tui(prompt: str, p_bar: Progress, config: Config) -> str:
    task = p_bar.add_task("design", total=PROGRESS_LIMIT)

    output = call_llm(
        system=PROMPT.format(library=config.library),
        messages=[{"role": "user", "content": prompt}],
        stream=True,
    )
    design = ""
    for token in output:
        design += token
        p_bar.update(task, advance=1)

    p_bar.update(task, completed=PROGRESS_LIMIT)
    return design
