### 1. Level A: Feature (Project Charter / One-Pager)
**Feature Name:** Agentic EARS – Recursive LLM-Agent Pipeline for Turning Vague Ideas into Executable Code

**Description (One-Pager Summary):**  
A user provides a single sentence or paragraph describing a new idea or a feature inside an existing application. A chain of specialized LLM agents (Charter Agent → Architect Agent → Planning Agent) iteratively refines the input through targeted questioning until it produces:  
1. A signed-off Project Charter (One-Pager)  
2. An Architectural & Design Document  
3. A detailed Project Plan containing Sprints, Tasks, and Sub-tasks (where Sub-tasks = atomic prompts for AI coding tools).

**Success Metrics:**  
- User can go from “I want a fitness tracker with voice notes” → production-ready prompts in under 3 agent sessions.  
- Zero context drift between agents.  
- Every Sub-task prompt is self-contained and < context-window limit.  
- Back-propagation supported (later agents can request Charter updates).

**Scope Boundaries:**  
- In-scope: Charter generation, Architecture doc, Sprint planning, Sub-task prompt generation.  
- Out-of-scope: Actual code execution, UI for the final app, deployment.

This Feature is the North Star. Everything below lives inside it.

---

### 2. Level B: Use Cases (Functional Specification)
Each Use Case is a concrete interaction script (the “C” in the original 4-level model). A Use Case describes the flow, actors, and success criteria. One Feature contains **many** Use Cases.

#### Use Case UC-B1: Generate & Approve Project Charter (One-Pager)
**Actors:** End User, Charter Agent (LLM)  
**Goal:** Convert a vague sentence/paragraph into a signed-off One-Pager.  
**Preconditions:** User provides initial idea text.  
**Main Flow:**  
1. User submits idea (1 sentence or paragraph).  
2. Charter Agent asks up to 8 clarifying questions (EARS-style).  
3. User answers; agent iterates until complete.  
4. Agent outputs One-Pager (goals, success metrics, boundaries, non-functional requirements).  
5. User reviews, edits, and explicitly approves.  
**Alternative Flows:** User rejects → loop back to questioning.  
**Post-condition:** Approved One-Pager is stored and passed to Architect Agent.

#### Use Case UC-B2: Produce Architectural & Design Document
**Actors:** Architect Agent (LLM), User (reviewer)  
**Goal:** Turn approved Charter into technical blueprint.  
**Main Flow:**  
1. Architect Agent loads approved One-Pager.  
2. Agent asks targeted technical questions (stack, constraints, integrations).  
3. Outputs Architecture & Design Doc (schema, modules, APIs, data flow, tech choices).  
4. User reviews and approves.  
**Post-condition:** Design Doc is versioned and passed to Planning Agent.

#### Use Case UC-B3: Create Project Plan with Sprints
**Actors:** Planning Agent (LLM)  
**Goal:** Break Use Cases + Design Doc into executable Sprints, Tasks, and Sub-tasks.  
**Main Flow:**  
1. Planning Agent loads Design Doc + all Use Cases.  
2. Groups work into logical Sprints/Phases.  
3. For each Sprint: defines Tasks → Sub-tasks (prompts).  
4. Includes Definition of Done for every Sub-task.  
5. Outputs full Project Plan / Roadmap.  
**Post-condition:** User can feed Sub-task prompts directly into coding tools.

#### Use Case UC-B4: Back-Propagation & Scope Change
**Actors:** Any later Agent, User  
**Goal:** If downstream agent discovers impossibility, signal upstream to update Charter.  
**Main Flow:**  
1. Architect/Planning Agent flags issue.  
2. System notifies user + Charter Agent.  
3. User approves updated One-Pager; pipeline restarts from affected point.

(Additional Use Cases can be added later; these four cover 90 % of the core workflow.)

---

### 3. Level C: Sprints / Phases (Project Plan / Roadmap)
Each Sprint groups work derived from the Use Cases above. One Use Case can span multiple Sprints; one Sprint contains **many** Tasks.

**Sprint C1: Foundation (2–3 days)**  
Goal: Build core agent orchestration and Charter Agent.

**Sprint C2: Architecture & Design Agent (3 days)**  
Goal: Implement Architect Agent + Design Doc generation.

**Sprint C3: Planning & Roadmap Agent (3–4 days)**  
Goal: Implement Planning Agent + full Project Plan output with Sprints/Tasks/Sub-tasks.

**Sprint C4: Back-Propagation & Polish (2 days)**  
Goal: Add change-request loops, UI/approval flows, Definition of Done enforcement.

**Sprint C5: Pilot & Validation (2 days)**  
Goal: End-to-end test with real user ideas.

---

### 4. Level D & E: Tasks and Sub-tasks (Implementation Guide)
Below is a sample breakdown for **Sprint C1** only (to illustrate the pattern). Every Sprint will follow the same structure: many Tasks per Sprint, many Sub-tasks per Task.

#### Sprint C1 – Tasks & Sub-tasks

**Task D1.1: Implement Charter Agent Core Loop**  
- **Sub-task E1.1.1 (Prompt):** “Using the approved system prompt for Charter Agent, write a Python function `run_charter_session(initial_idea: str) -> dict` that: (1) asks up to 8 clarifying questions one at a time, (2) maintains conversation state, (3) outputs structured One-Pager JSON when complete. Use the exact EARS-style question templates from the Design Doc. Return the JSON only when user types ‘APPROVE’.”
- **Sub-task E1.1.2 (Prompt):** “Add validation: ensure One-Pager contains sections for Goals, Success Metrics, Boundaries, and Non-Functional Requirements. If missing, loop back to user.”

**Task D1.2: Build Persistent Storage for Documents**  
- **Sub-task E1.2.1 (Prompt):** “Create a simple SQLite schema with tables `projects`, `charters`, `design_docs`, `plans`. Write CRUD functions that store versioned JSON blobs. Include a `status` field (draft/approved).”

**Task D1.3: Define “Definition of Done” Enforcement**  
- **Sub-task E1.3.1 (Prompt):** “For every Sub-task prompt generated by the Planning Agent, automatically append this block at the end: ‘Definition of Done: Output must be valid Python, pass all unit tests you generate, respect the schema from the Design Doc, and stay under 8000 tokens.’”

(Every other Sprint follows the identical pattern: 3–6 Tasks per Sprint, 2–4 Sub-tasks per Task, each Sub-task = one atomic prompt.)

---

### 5. Additional System Requirements (Cross-Cutting)
- All agents must be stateless except for the current session context passed explicitly.  
- Every output document must be version-controlled and user-approved before proceeding.  
- Sub-task prompts must never exceed the target coding tool’s context window.  
- Back-propagation: any agent can raise a “ChangeRequest” flag that routes back to the Charter Agent.  
- The Planning Agent must always include a “Definition of Done” paragraph in every Sub-task prompt.
