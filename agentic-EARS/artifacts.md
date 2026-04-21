
**Agentic EARS – Artifacts Document**  
**Version 1.0**

This document defines every artifact produced by the Agentic EARS pipeline, including its purpose, content, who creates it, and how it fits into the hierarchy.

### Hierarchy Recap (for reference)
**NorthStar** → **Feature** → **Requirements** → **Use Case** → **Sprint** → **Task** → **Sub-task / Prompt**

### 1. Project Charter / One-Pager
**Produced by**: Charter Agent  
**When**: Step 1 (from user’s Seed Idea)  
**Purpose**: High-level anchor document that captures the vision and MVP scope. This is the single approved document that everything else must trace back to and support.

**Content**:
- **NorthStar**: One concise sentence describing the overarching mission/vision.
- **Features**: List of 3–8 MVP Features (user-visible capabilities in colloquial language). Each Feature includes a short description.
- **Success Metrics**: 3–6 measurable outcomes.
- **Scope Boundaries**: In-scope and out-of-scope lists.
- **Non-functional Notes**: High-level constraints (performance, privacy, security, etc.).
- **Version & Status**: Version number and approval status.

**Does NOT contain**: EARS Requirements or Use Cases (these come later).

**Format**: Markdown document + structured JSON version.  
**Approval Required**: Yes – User must reply “APPROVE”, “FINALIZE”, or “LOOKS GOOD”.

---

### 2. Architectural & Design Document (ADD)
**Produced by**: Architect Agent  
**When**: Step 2 (after One-Pager approval)  
**Purpose**: Technical foundation document. This is where **Requirements** and **Use Cases** are formally created and the “how” of the system is designed.

**Content**:
- **Requirements**: EARS-style statements for every Feature (functional + non-functional). Example: “When the user records a voice memo, the system shall transcribe it using the Whisper API...”
- **Use Cases**: Detailed flows for each major Requirement (main flow, alternate flows, preconditions, postconditions, exceptions).
- **Architecture**: Modules, layers, data models, APIs, data flows, diagrams (text or Mermaid).
- **Tech Stack**: Chosen technologies, frameworks, databases, libraries.
- **Constraints & Decisions**: Performance, scalability, security, integration points.
- **Traceability Matrix**: Mapping Features → Requirements → Use Cases.

**Format**: Comprehensive Markdown document (with sections and optional diagrams).  
**Approval Required**: Yes – User must reply “APPROVE DESIGN” or “FINALIZE ADD”.

---

### 3. Project Plan / Roadmap
**Produced by**: Planning Agent  
**When**: Step 3 (after ADD approval)  
**Purpose**: Executable breakdown of all work into time/logic-bound chunks and atomic coding prompts.

**Content**:
- **Sprints / Phases**: List of Sprints with goals and dependencies.
- **Tasks**: Technical milestones inside each Sprint.
- **Sub-tasks**: Atomic prompts for the AI coding tool. Each Sub-task includes:
  - The actual prompt text
  - Definition of Done
  - References to relevant Requirements, ADD sections, and Test Plan
- **Timeline / Dependencies**: High-level ordering or effort estimates (optional).
- **Traceability**: Links back to Features, Requirements, and Use Cases.

**Format**: Markdown + structured sections (can include JSON export).  
**Approval Required**: Yes – User must reply “APPROVE PLAN” or “GENERATE ALL ARTIFACTS”.

---

### 4. AGENTS.md
**Produced by**: Planning Agent (as part of Step 3)  
**Purpose**: Persistent grounding and guardrail file placed at the root of the project repository. Every time a Sub-task prompt is sent to an AI coding tool, **AGENTS.md** must be included in the context to prevent drift.

**Content**:
- **Project Overview**: NorthStar + short summary of Features.
- **Core Principles & Guardrails**: Key rules derived from all higher documents.
- **Architecture References**: Pointers to important sections of the ADD.
- **Requirements Summary**: Brief list or links to key EARS Requirements.
- **Coding Standards**: Naming conventions, style, testing expectations, file structure.
- **Workflow Rules**: How to handle back-propagation, approvals, Definition of Done.
- **Do’s and Don’ts**: Explicit prohibitions to avoid scope creep or inconsistencies.
- **Version Note**: Current versions of One-Pager, ADD, and Project Plan.

**Format**: Clean Markdown file (kept concise, ideally under 2500–3000 tokens).  
**Approval Required**: Implicit (part of Planning Agent output). User reviews with the Project Plan.

---

### 5. Test Plan
**Produced by**: Planning Agent (as part of Step 3)  
**Purpose**: Defines how quality will be ensured. Provides acceptance criteria and testing guidance for the AI coding tool.

**Content**:
- **Testing Strategy**: Unit, integration, and end-to-end testing approach.
- **Acceptance Criteria**: Tied directly to Requirements and Use Cases.
- **Test Types & Coverage Goals**: What must be tested and target coverage.
- **Test Generation Guidance**: Instructions for AI coders on writing tests alongside code.
- **Edge Cases**: Important scenarios from Use Cases.
- **Tools & Commands**: Test running commands and frameworks.

**Format**: Markdown document.  
**Approval Required**: Implicit (part of Planning Agent output).

---

### Artifact Traceability & Relationships

| Artifact                        | Created By       | Contains / Defines                  | References                          | Approval Gate                  |
|---------------------------------|------------------|-------------------------------------|-------------------------------------|--------------------------------|
| Project Charter / One-Pager     | Charter Agent    | NorthStar + Features                | Seed Idea                           | User “APPROVE”                 |
| Architectural & Design Doc (ADD)| Architect Agent  | Requirements + Use Cases + Design   | One-Pager                           | User “APPROVE DESIGN”          |
| Project Plan                    | Planning Agent   | Sprints → Tasks → Sub-tasks         | One-Pager + ADD                     | User “APPROVE PLAN”            |
| AGENTS.md                       | Planning Agent   | Grounding rules + references        | All previous artifacts              | Reviewed with Project Plan     |
| Test Plan                       | Planning Agent   | Testing strategy + acceptance criteria | Requirements + Use Cases         | Reviewed with Project Plan     |

### Key Rules for All Artifacts
- **Traceability**: Every lower-level item (Requirement, Use Case, Task, Sub-task) must trace upward to a Feature and ultimately the NorthStar.
- **Versioning**: All artifacts should include version numbers and date.
- **Back-Propagation**: If issues are found, affected artifacts are updated and downstream work is re-generated only for changed parts.
- **User Approval**: No artifact moves to the next stage without explicit user approval.
- **Grounding**: All coding Sub-tasks must reference **AGENTS.md**.


### Additional Recommended Artifacts
Beyond the One-Pager, ADD, Project Plan (with Tasks/Prompts), and **AGENTS.md**, the following artifacts would be highly useful in an agentic workflow like this. They address common pain points in LLM-driven development (context drift, quality, traceability, and handoff):

- **TEST-PLAN.md** or **Testing Strategy Document** (produced by Planning Agent or a new Testing-focused sub-agent during Sprint planning)  
  Defines unit/integration test strategy, acceptance criteria per Use Case/Sprint, and how the AI Coding tool should generate tests alongside code. Includes test coverage goals and edge-case lists. Useful because Sub-task prompts can then reference "Write tests per TEST-PLAN.md".

- **DECISION-LOG.md** (or RUN-LOG.md) (maintained/updated by Planning Agent or Coding Agent after each Sprint)  
  A living record of key decisions, trade-offs, why certain tech choices were made, and any back-propagation events. Helps prevent "why did we do it this way?" questions in future sessions and provides auditability.

- **GLOSSARY.md** or **TERMINOLOGY.md** (produced early by Architect Agent)  
  Consistent definitions of domain terms, acronyms, and entity names used across the project. Prevents hallucinations where the coding agent invents conflicting names.

- **PROMPTS.md** or **Prompt Library** (optional, produced by Planning Agent)  
  A curated collection of reusable prompt templates (e.g., the Charter questioning template, Definition of Done boilerplate). Useful if your system grows beyond one project.

- **CHANGELOG.md** (standard, but auto-updated by agents)  
  Tracks versions of the generated documents themselves and major code changes tied back to Sprints/Use Cases.
  
  
  ### Recommended Structure for AGENTS.md
- Project Overview (from One-Pager Feature + Requirements)  
- Core Guardrails & Principles (including EARS usage)  
- Architecture References (pointers to ADD)  
- Coding & Testing Standards (links to Test Plan)  
- Workflow Rules (back-propagation, approval gates, Definition of Done)  
- Do’s and Don’ts (to prevent drift)  


### How the Agents Produce the Artifacts (Updated Workflow)
1. **Charter Agent** (starts from user’s Feature idea)  
   - Asks refining questions.  
   - Produces **Project Charter / One-Pager** (includes initial Feature description + high-level Requirements in EARS format).  
   - User reviews and approves.

2. **Architect Agent**  
   - Loads approved One-Pager.  
   - Asks technical questions.  
   - Produces **Architectural & Design Document (ADD)**.  
   - Expands Requirements into detailed functional and non-functional specs.  
   - User approves.

3. **Planning Agent** (or Documentation sub-agent)  
   - Loads One-Pager + ADD.  
   - Derives full set of Use Cases from Requirements.  
   - Breaks work into Sprints → Tasks → Sub-tasks (prompts).  
   - Produces:  
     - **Project Plan** (with all Sprints, Tasks, and Sub-task prompts)  
     - **AGENTS.md** (synthesizes rules from One-Pager + ADD + Plan)  
     - **Test Plan** (acceptance criteria tied to Requirements/Use Cases)
### Detailed Mapping of Artifacts to Hierarchy Levels

**Level A: Feature**  
→ **Artifact**: Project Charter / One-Pager  
**Content Focus**: User’s initial sentence/paragraph is refined into a clear Feature definition + high-level boundaries.  
**Example Output Section**: “Feature: Agentic EARS – A recursive agent pipeline that turns vague ideas into production-ready code via specialized LLM agents.”

**Level B: Requirements**  
→ Derived and formalized in the **One-Pager** and further detailed in the **ADD**.  
These are written using **EARS syntax** (e.g., “When the user submits an initial idea, the Charter Agent shall ask clarifying questions until a complete One-Pager can be generated.”).  
Requirements are traceable to the Feature and drive Use Cases.

**Level C: Use Case**  
→ Described in the **ADD** (functional flows) and referenced in the **Project Plan**.  
Each Use Case links back to one or more Requirements.

**Level D: Sprint / Phase**  
→ Defined in the **Project Plan**.  
Example: Sprint 1 – Foundation (Charter Agent + storage), Sprint 2 – Architect Agent, etc.

**Level E: Task**  
→ Listed inside each Sprint in the **Project Plan**.

**Level F: Sub-task**  
→ The actual prompts stored in the **Project Plan**. Each Sub-task prompt explicitly references:  
- AGENTS.md (for grounding)  
- Relevant sections of the ADD  
- The Test Plan (for testing expectations)

### Updated Artifacts Produced by the Pipeline
The system now generates these core documents (plus recommended additions):

1. **Project Charter / One-Pager** (produced by Charter Agent at Level A)  
   High-level vision, goals, success metrics, boundaries, and non-functional requirements. This anchors everything.

2. **Architectural & Design Document (ADD)** (produced by Architect Agent after Level A approval)  
   Technical blueprint: system architecture, data models, modules, APIs, tech stack, data flows, constraints.

3. **Project Plan / Roadmap** (produced by Planning Agent)  
   Breaks everything into Sprints (C), Tasks (D), and Sub-tasks (E/prompts). Includes Definition of Done for each Sub-task.

4. **AGENTS.md** (new – produced by Planning Agent or a dedicated Documentation Agent as part of the final Project Plan phase)  
   A single, persistent grounding file placed at the root of the project repository. It enforces consistency across all AI coding sessions by summarizing and referencing rules from the One-Pager, ADD, and Project Plan.  
   **Purpose**: Serves as the "README for agents" — every time you feed a Sub-task prompt to your AI coding tool (Cursor, Copilot, etc.), you also include or reference **AGENTS.md** so the coder stays grounded and doesn't drift from prior decisions.

   **Recommended Structure for AGENTS.md** (tailored to Agentic EARS):
   - **Project Overview** — Short summary from the One-Pager (what the Feature is, who it's for, success criteria).
   - **Core Principles & Guardrails** — Key rules derived from Charter and ADD (e.g., "Always respect the approved tech stack", "No scope creep without back-propagation", "Sub-tasks must stay under context-window limits").
   - **Agent Roles & Responsibilities** — Brief definitions of Charter Agent, Architect Agent, Planning Agent, and Coding Agent behaviors.
   - **Architecture & Design References** — Pointers to key sections of the ADD (e.g., "Use the PostgreSQL schema defined in ADD section 3.2").
   - **Coding Standards & Style** — Language-specific rules, file structure, naming conventions, testing expectations.
   - **Workflow & Processes** — How to handle back-propagation, Definition of Done enforcement, commit/PR guidelines for AI-generated changes.
   - **Commands & Tools** — Build/test commands, how to run the system, any special agent tooling.
   - **Boundaries (Do's and Don'ts)** — Explicit prohibitions (e.g., "Do not add new features without updating the Charter", "Never hallucinate integrations not in the ADD").
   - **Versioning Note** — Link to current versions of One-Pager, ADD, and Project Plan.
se artifacts (in order):

