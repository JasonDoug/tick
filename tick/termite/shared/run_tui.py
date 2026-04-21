# Standard library
import os
import re
import ast
import time
import tempfile
from typing import Tuple
from subprocess import TimeoutExpired, DEVNULL, PIPE, run as run_cmd


# Local
from tick.termite.dtos.Script import Script
from tick.termite.shared.utils.fix_imports import fix_any_import_errors
from tick.termite.shared.utils.python_exe import get_python_executable


#########
# HELPERS
#########


def save_script_to_file(script: Script) -> str:
    with tempfile.NamedTemporaryFile(mode="w", suffix=".py", delete=False) as temp_file:
        temp_file.write(script.code)
        temp_file_path = temp_file.name

    return temp_file_path


def strip_ansi_escape_sequences(data: str) -> str:
    ansi_escape = re.compile(
        r"""
        \x1B
        (?:
            [@-Z\\-_]
        |
            \[
            [0-?]*
            [ -/]*
            [@-~]
        |
            \]
            (?:
                [^\x07]*?
                \x07
            |
                [^\x1B]*?
                \x1B\\
            )
        |
            [PX^_]
            .*?
            \x1B\\
        )
        """,
        re.VERBOSE | re.DOTALL,
    )
    return ansi_escape.sub("", data)


def run_in_pseudo_terminal(script: Script, timeout: int = 5) -> Tuple[str, str]:
    python_exe = get_python_executable()
    tui_file = save_script_to_file(script)
    # The runner_file is in tick/termite/shared/utils/run_pty.py
    runner_file = os.path.join(os.path.dirname(__file__), "utils", "run_pty.py")

    stdout, stderr = "", ""
    try:
        result = run_cmd(
            [
                python_exe,
                runner_file,
                tui_file,
            ],
            stdout=PIPE,
            stderr=PIPE,
            text=True,
            timeout=timeout,
        )
        stdout = result.stdout if result.stdout else ""
        stderr = result.stderr if result.stderr else ""
    except TimeoutExpired as e:
        stdout = e.stdout.decode() if e.stdout else ""
        stderr = e.stderr.decode() if e.stderr else ""
    finally:
        os.remove(tui_file)

    stdout = strip_ansi_escape_sequences(stdout)
    stderr = strip_ansi_escape_sequences(stderr)
    return stdout.strip(), stderr.strip()


def run_in_subprocess(script: Script):
    python_exe = get_python_executable()
    script_file = save_script_to_file(script)
    run_cmd([python_exe, script_file])


######
# MAIN
######


def run_tui(script: Script, pseudo=True):
    if not pseudo:
        run_in_subprocess(script)
        return

    stdout, stderr = "", ""
    try:
        # Check for valid syntax before executing
        ast.parse(script.code)

        # Execute the script, iteratively fixing any import errors
        retry = True
        while retry:
            stdout, stderr = run_in_pseudo_terminal(script)
            
            # Combine stdout and stderr for error detection
            # Most TUI errors go to stderr, but some libraries or wrappers might be messy
            all_output = f"{stdout}\n{stderr}".strip()
            
            try:
                retry = fix_any_import_errors(all_output)
            except ImportError as e:
                retry = False
                stderr = str(e)
            
            if not retry:
                # If there's a traceback in either, treat it as an error
                if "Traceback" in all_output:
                    stderr = all_output
                break
    except SyntaxError as e:
        stderr = str(e)

    script.stdout = stdout
    script.stderr = stderr
