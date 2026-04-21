# GEMINI.md - Agentic EARS Project Context

## Project Overview
**Agentic EARS** is a recursive, agent-driven refinement pipeline designed to convert a vague **Seed Idea** (a single sentence or paragraph) into fully executable, atomic coding prompts for AI development tools (like Cursor, Copilot, or the Gemini CLI itself). It solves the problem of "Context Drift" in LLM-driven development by enforcing strict approval gates, hierarchical traceability, and persistent grounding through specialized artifacts.

### The Core Methodology (The SDLC Hierarchy)
The system operates on a 7-level hierarchy of refinement, moving from abstract vision to concrete execution:

1.  **NorthStar**: The single, unifying vision/mission of the project.
2.  **Feature**: MVP capabilities described in colloquial, user-visible terms.
3.  **Requirements**: Precise, verifiable "shall" statements using **EARS** (Easy Approach to Requirements Syntax).
4.  **Use Case**: Detailed behavioral interaction flows (main flow, alternate flows, preconditions).
5.  **Sprint / Phase**: Logical or time-bound groupings of development work.
6.  **Task**: Technical milestones or buildable components.
7.  **Sub-task / Prompt**: The atomic, single-purpose instruction fed directly to an AI coding tool.

---

## Agent Roles & Workflow
The pipeline uses three specialized agents, each with its own system prompt and output responsibility:

| Agent | Input | Output Artifact(s) | Goal |
| :--- | :--- | :--- | :--- |
| **Charter Agent** | Seed Idea | **Project Charter / One-Pager** | Refine the vision into a NorthStar and 3–8 MVP Features. |
| **Architect Agent** | Approved One-Pager | **Architectural & Design Doc (ADD)** | Create EARS Requirements, Use Cases, and technical foundation. |
| **Planning Agent** | One-Pager + ADD | **Project Plan**, **AGENTS.md**, **Test Plan** | Breakdown into Sprints/Tasks/Prompts and create grounding files. |

### The "One Question at a Time" Rule
All agents are configured to ask only **one clarifying question at a time**. This prevents overwhelming the user and ensures high-quality, focused data gathering.

---

## Core Artifacts
Every interaction with this project should respect and reference these primary artifacts:

- **Project Charter / One-Pager**: The anchor document containing the NorthStar, Features, Success Metrics, and Scope.
- **Architectural & Design Document (ADD)**: The technical blueprint (modules, data models, APIs, and EARS Requirements).
- **Project Plan / Roadmap**: The executable breakdown of work into Sprints, Tasks, and Sub-tasks.
- **AGENTS.md**: **CRITICAL.** The persistent grounding file placed at the project root. It acts as the "README for agents," summarizing the NorthStar, guardrails, and standards to be used in every coding session.
- **Test Plan**: Acceptance criteria tied directly to Requirements and Use Cases.

---

## Key Files Reference
- `Workflow-Master-Reference.md`: The definitive "Source of Truth" for the entire system methodology.
- `sldc.md` & `hierarchy.md`: Detailed definitions of the SDLC levels (NorthStar to Sub-task).
- `agents.md`: System prompts and role definitions for Charter, Architect, and Planning agents.
- `artifacts.md`: Detailed templates and content requirements for every document in the pipeline.
- `structured-outputs.md`: JSON schemas for agents to produce consistent artifact data.

---

## Operational Conventions
- **Approval Gates**: No agent may proceed to the next stage without explicit user approval (e.g., "APPROVE", "APPROVE DESIGN", "APPROVE PLAN").
- **Back-Propagation**: If a technical impossibility or scope issue is discovered by a later agent, a `ChangeRequest` flag is raised, routing context back to the Charter Agent for refinement.
- **Grounding Requirement**: Every AI-generated coding prompt (Sub-task) **must** reference `AGENTS.md` and the relevant section of the `ADD`.
- **Atomic Prompts**: Sub-tasks must be small enough to fit within an LLM's context window and produce a single, verifiable output.

---

## Development & Usage
To use this system:
1. Provide a **Seed Idea**.
2. Interact with the **Charter Agent** until a One-Pager is approved.
3. Interact with the **Architect Agent** until the ADD is approved.
4. Interact with the **Planning Agent** until the Project Plan and `AGENTS.md` are generated.
5. Use the generated **Sub-tasks** as prompts for your coding tool, always attaching `AGENTS.md` for context.
