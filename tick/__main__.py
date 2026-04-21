import os
import json
import subprocess
import argparse
from pathlib import Path
from rich.console import Console

try:
    from tick.termite.termite import termite
    from tick.termite.dtos import Config
    from tick.termite.shared import run_tui
except ImportError:
    from tick.termite.termite import termite
    from tick.termite.dtos import Config
    from tick.termite.shared import run_tui

console = Console()

def build_ears_go():
    ears_go_dir = Path(__file__).parent.parent / "ears-go"
    binary_path = ears_go_dir / "tick-ears"
    
    if not binary_path.exists():
        console.log("[yellow]Building ears-go binary...[/yellow]")
        subprocess.run(["go", "build", "-o", "tick-ears"], cwd=ears_go_dir, check=True)
    
    return binary_path

def run_ears_go(binary_path):
    # This will be interactive, so we use subprocess.run without capturing stdout
    # except for the final output if we want, but the agents write to files.
    subprocess.run([str(binary_path)], check=True)

def load_artifacts():
    ears_go_dir = Path(__file__).parent.parent / "ears-go"
    plan_path = ears_go_dir / "project_plan.json"
    add_path = ears_go_dir / "add.json"
    
    if not plan_path.exists() or not add_path.exists():
        raise FileNotFoundError("Artifacts not found. EARS pipeline may have failed.")
    
    with open(plan_path, "r") as f:
        plan = json.load(f)
    
    with open(add_path, "r") as f:
        add = json.load(f)
        
    return plan, add

def extract_tui_prompt(plan, add):
    # Search for a sub-task that looks like a TUI implementation prompt
    for sprint in plan.get("sprints", []):
        for task in sprint.get("tasks", []):
            for sub_task in task.get("sub_tasks", []):
                prompt = sub_task.get("prompt", "")
                if "TUI" in prompt.upper() or "TERMINAL" in prompt.upper():
                    return prompt
    
    # Fallback: combine NorthStar and Features into a prompt
    return f"Build a TUI based on the following ADD: {json.dumps(add)}"

def main():
    parser = argparse.ArgumentParser(description="Tick: Agentic EARS + Termite")
    parser.add_argument("--library", type=str, default="urwid", help="TUI library to use")
    parser.add_argument("--refine", action="store_true", help="Enable refinement loop")
    args = parser.parse_args()

    try:
        binary_path = build_ears_go()
        run_ears_go(binary_path)
        
        plan, add = load_artifacts()
        prompt = extract_tui_prompt(plan, add)
        
        console.log(f"[bold green]Specs received! Handing off to Termite...[/bold green]")
        
        config = Config(
            library=args.library,
            should_refine=args.refine,
            refine_iters=1,
            fix_iters=10
        )
        
        # We pass the full ADD as the "design" to termite to ensure high fidelity
        tui = termite(prompt, config, design=json.dumps(add))
        
        console.log("[bold green]TUI Generated Successfully![/bold green]")
        run_tui(tui, pseudo=False)
        
    except Exception as e:
        console.print(f"[bold red]Error: {e}[/bold red]")

if __name__ == "__main__":
    main()
