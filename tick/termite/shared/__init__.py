try:
    from tick.termite.shared.run_tui import run_tui
    from tick.termite.shared.call_llm import call_llm, MAX_TOKENS
except ImportError:
    from tick.shared.run_tui import run_tui
    from tick.shared.call_llm import call_llm, MAX_TOKENS
