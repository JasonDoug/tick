**Agentic EARS Workflow Document**  
**Version 1.0**  
**NorthStar → Feature → Requirements → Use Case → Sprint → Task → Sub-task Pipeline**

### 1. Overall Workflow Overview

The Agentic EARS system is a recursive, agent-driven refinement pipeline that converts a vague **Seed Idea** (one sentence or paragraph) into fully executable code through specialized LLM agents and clear approval gates.

**Core Principle**:  
Every step produces a specific artifact. The user must explicitly approve each artifact before the next agent runs. This prevents context drift and enables clean back-propagation.

**Full Hierarchy (for reference)**  
- **NorthStar** (single unifying vision)  
- **Feature** (MVP capabilities – colloquial units)  
- **Requirements** (EARS-style “shall” statements – created in ADD)  
- **Use Case** (behavioral interaction flows)  
- **Sprint / Phase** (logical grouping of work)  
- **Task** (technical milestones)  
- **Sub-task / Prompt** (atomic prompts for AI coding tools – most granular)

### 2. End-to-End Workflow Steps

**Step 0: User Provides Seed Idea**  
User enters a single sentence or short paragraph describing the project or feature.  
Example: “I want a simple fitness tracker where I can record voice memos during workouts and get AI summaries.”

**Step 1: Charter Agent Session**  
- **Agent**: Charter Agent  
- **Goal**: Refine the Seed Idea into a high-level vision and break it into MVP Features.  
- **Process**:  
  - Asks clarifying questions **one at a time**.  
  - Proposes a single **NorthStar**.  
  - Decomposes into 3–8 **MVP Features**.  
  - Adds success metrics, scope boundaries (in/out), and high-level non-functional notes.  
- **Output Artifact**: **Project Charter / One-Pager** (Markdown + JSON)  
  - Contains: NorthStar + List of Features + Success Metrics + Scope + Non-functional notes  
  - **Does NOT** contain Requirements or Use Cases.  
- **Approval Gate**: User reviews and replies with “APPROVE”, “FINALIZE”, or “LOOKS GOOD”.  
- **Back-propagation**: If later agents find issues, they can request updates to the One-Pager.

**Step 2: Architect Agent Session**  
- **Agent**: Architect Agent  
- **Goal**: Turn the approved One-Pager into technical foundation.  
- **Process**:  
  - Loads the approved Project Charter / One-Pager.  
  - For each Feature, creates **EARS-style Requirements** (functional + non-functional).  
  - Expands Requirements into detailed **Use Cases** (main flow, alternate flows, preconditions, etc.).  
  - Defines system architecture, data models, APIs, tech stack, data flows, and constraints.  
  - Asks technical/design questions one at a time if needed.  
- **Output Artifact**: **Architectural & Design Document (ADD)**  
- **Approval Gate**: User replies with “APPROVE DESIGN” or “FINALIZE ADD”.  

**Step 3: Planning Agent Session**  
- **Agent**: Planning Agent  
- **Goal**: Convert the design into an executable plan and grounding materials.  
- **Process**:  
  - Loads approved One-Pager + ADD.  
  - Groups work into logical **Sprints**.  
  - Breaks Sprints into **Tasks** and then into atomic **Sub-tasks** (each Sub-task = one focused prompt for the AI coding tool).  
  - Generates **AGENTS.md** (single grounding file that references NorthStar, Features, Requirements, ADD, coding standards, and guardrails).  
  - Generates **Test Plan** (acceptance criteria tied to Requirements and Use Cases).  
- **Output Artifacts** (three files):  
  1. **Project Plan / Roadmap** (Sprints → Tasks → Sub-tasks/Prompts)  
  2. **AGENTS.md**  
  3. **Test Plan**  
- **Approval Gate**: User replies with “APPROVE PLAN” or “GENERATE ALL ARTIFACTS”.

**Step 4: Execution Phase (Coding)**  
- Developer / User takes Sub-tasks from the Project Plan.  
- For each Sub-task:  
  - Feed the prompt to your AI coding tool (Cursor, Copilot, etc.).  
  - **Always include or attach AGENTS.md** in the context for grounding.  
  - Reference relevant sections of the ADD and Test Plan as instructed in the prompt.  
- After completing a Sprint’s Sub-tasks, move to the next Sprint.  
- Run tests per the Test Plan.

**Step 5: Back-Propagation & Iteration**  
- At any point, if an agent or the user discovers a problem (technical impossibility, missing Requirement, scope issue, etc.):  
  - Raise a **ChangeRequest** flag.  
  - Route back to the Charter Agent (or appropriate earlier agent) with the specific issue.  
  - Update the One-Pager / ADD as needed.  
  - Re-run downstream agents only for affected parts.

### 3. Workflow Diagram (Text Version)

```
Seed Idea (User)
      ↓
Charter Agent → Project Charter / One-Pager (NorthStar + Features)
      ↓ (User APPROVE)
Architect Agent → Architectural & Design Document (ADD)
      ↓ (Requirements created here + Use Cases)
      ↓ (User APPROVE DESIGN)
Planning Agent → Project Plan + AGENTS.md + Test Plan
      ↓ (User APPROVE PLAN)
Execution: Sub-task Prompts → AI Coding Tool (with AGENTS.md)
      ↓
Back-propagation loop (if needed) → return to Charter/Architect Agent
```

### 4. Key Rules Across the Workflow
- **One question at a time**: No agent ever asks multiple questions in a single response.
- **Explicit approval gates**: No automatic progression. User must approve each major artifact.
- **Traceability**: Every Requirement, Use Case, Task, and Sub-task must trace upward to a Feature and ultimately the NorthStar.
- **Grounding**: All coding prompts reference **AGENTS.md**.
- **Atomic prompts**: Sub-tasks must be small enough to fit comfortably in the coding tool’s context window.
- **Definition of Done**: Every Sub-task prompt includes a clear Definition of Done (including testing per Test Plan).

### 5. Agent Hand-off Summary
| Step | Agent            | Input                          | Output Artifact(s)                          | Approval Phrase Examples          |
|------|------------------|--------------------------------|---------------------------------------------|-----------------------------------|
| 1    | Charter Agent    | Seed Idea                      | Project Charter / One-Pager                 | APPROVE, FINALIZE, LOOKS GOOD     |
| 2    | Architect Agent  | Approved One-Pager             | Architectural & Design Document (ADD)       | APPROVE DESIGN, FINALIZE ADD      |
| 3    | Planning Agent   | Approved One-Pager + ADD       | Project Plan + AGENTS.md + Test Plan        | APPROVE PLAN, GENERATE ALL        |

This Workflow Document serves as the master reference for how the entire Agentic EARS system operates.
