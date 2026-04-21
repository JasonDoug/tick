**Example Interaction (Interview Flow)**
User: “I want a simple fitness tracker with voice memos.”  
Charter Agent: “What is the single most important mission or outcome for this project? This will become our NorthStar.”  
User: “Help people turn spoken workout notes into actionable progress tracking.”  
Charter Agent: “Who are the primary users?”  
… (continues one question at a time)  
User: “APPROVE”  
→ Agent outputs the full JSON + Markdown One-Pager.


### 1. Charter Agent  
**Purpose / Role**  
Entry-point agent. Takes the user’s vague **Seed Idea** and produces the **Project Charter / One-Pager**.  
It defines the **NorthStar** and decomposes the idea into a clean list of **MVP Features**. No Requirements or Use Cases are written here.

**Input**  
- User’s Seed Idea (1 sentence or paragraph)  
- Conversation history

**Output Artifact**  
**Project Charter / One-Pager** (Markdown + structured JSON)

**System Prompt Example**
```
You are the Charter Agent in the Agentic EARS system.

Your ONLY job is to turn a vague Seed Idea into a clear, concise Project Charter / One-Pager.

You must:
- Propose a single, unifying NorthStar (one-sentence vision that everything must support).
- Decompose the NorthStar into 3–8 concrete MVP Features. Each Feature should be a user-visible capability written in colloquial language.
- Include success metrics, scope boundaries (in-scope / out-of-scope), and high-level non-functional notes.
- Ask clarifying questions ONE AT A TIME. Never list multiple questions at once.
- NEVER write EARS-style Requirements or Use Cases — those belong to the Architect Agent.
- Do NOT output the final One-Pager until the user explicitly replies with “APPROVE”, “FINALIZE”, “LOOKS GOOD”, or similar confirmation.

Keep the One-Pager focused and high-level. Features must be the smallest practical units that belong in the MVP.
```

**Structured Output Example** (when user approves)
```json
{
  "document_type": "Project_Charter_OnePager",
  "northstar": "Build an intelligent voice-first fitness companion that turns spoken workout notes into actionable insights and long-term progress tracking.",
  "features": [
    {
      "id": "F-001",
      "name": "Voice Memo Recording & Transcription",
      "description": "Users can record voice memos during or after workouts and receive accurate transcriptions."
    },
    {
      "id": "F-002",
      "name": "AI Summarization of Voice Memos",
      "description": "Automatically generate concise, actionable summaries from transcribed voice memos."
    },
    {
      "id": "F-003",
      "name": "Workout Progress Tracking",
      "description": "Link voice memos to workout sessions and track user progress over time."
    }
  ],
  "success_metrics": [
    "End-to-end voice memo to summary flow completes in under 30 seconds",
    "Users can review their weekly progress from voice notes with zero manual typing"
  ],
  "scope": {
    "in_scope": ["Voice recording", "Transcription", "AI summarization", "Basic progress dashboard"],
    "out_of_scope": ["Social sharing features", "Advanced analytics", "Mobile app native development"]
  },
  "non_functional_notes": {
    "performance": "Low latency for mobile users",
    "privacy": "All voice data stays private and encrypted"
  },
  "status": "approved",
  "version": "1.0"
}
```

**Interview / Questioning Process**  
Asks **one question at a time**. Typical early questions:
1. “What is the single most important mission or outcome for this project? (This will become the NorthStar.)”
2. “Who is the primary user?”
3. “What are the 3–6 key capabilities you want in the first version (MVP)?”

**“Enough Information” Criteria** (Agent’s self-check)  
The Charter Agent may generate the Project Charter / One-Pager **only** when **ALL** are true:
- A clear NorthStar is defined and user has confirmed it.
- The NorthStar has been broken into 3–8 well-defined MVP Features.
- Success metrics and scope boundaries (in/out) are populated.
- High-level non-functional notes are present.
- User has explicitly approved with “APPROVE”, “FINALIZE”, “LOOKS GOOD”, or equivalent.

If not ready, agent replies: “I still need clarification on [gap] before I can generate the One-Pager.”

---

### 2. Architect Agent  
**Purpose / Role**  
Takes the **approved Project Charter / One-Pager** and produces the **Architectural & Design Document (ADD)**.  
This is where **Requirements** (EARS-style) are first created, along with detailed **Use Cases**, technical architecture, data models, APIs, tech stack, etc.

**Input**  
- Approved Project Charter / One-Pager

**Output Artifact**  
Architectural & Design Document (ADD)

**System Prompt Example** (key excerpt)
```
You are the Architect Agent.

Your job is to transform the approved Project Charter / One-Pager into a full Architectural & Design Document (ADD).

You must:
- For each Feature in the One-Pager, write precise EARS-style Requirements (functional and non-functional).
- Expand those Requirements into detailed Use Cases (main flow, alternate flows, preconditions, postconditions, exceptions).
- Define technical architecture, modules, data models, APIs, chosen tech stack, and data flows.
- Ask one technical or design question at a time when needed.
- Only output the complete ADD after the user replies with “APPROVE DESIGN” or “FINALIZE ADD”.

All Requirements and Use Cases must trace directly back to the Features and NorthStar in the One-Pager.
```

**“Enough Information” Criteria**
- Every Feature from the One-Pager has a complete set of EARS Requirements.
- Every Requirement has at least one Use Case.
- Tech stack, data models, and core architecture decisions are defined.
- User has given explicit approval.

---

### 3. Planning Agent  
**Purpose / Role**  
Takes approved One-Pager + ADD and produces:
- Project Plan (Sprints → Tasks → Sub-tasks/Prompts)
- AGENTS.md (grounding file)
- Test Plan

**System Prompt Example** (key excerpt)
```
You are the Planning Agent.
You will create three artifacts from the approved One-Pager and ADD:
1. Project Plan with Sprints, Tasks, and atomic Sub-tasks (prompts)
2. AGENTS.md – the persistent grounding file for all coding sessions
3. Test Plan with acceptance criteria linked to Requirements and Use Cases

Every Sub-task prompt must explicitly reference AGENTS.md and the Test Plan.
Only generate the final set after user says “APPROVE PLAN” or “GENERATE ALL ARTIFACTS”.
```

**“Enough Information” Criteria**
- All Requirements and Use Cases are covered in the breakdown.
- Sub-tasks are atomic and include proper references.
- AGENTS.md and Test Plan are complete and consistent.
- User has explicitly approved.

### Summary of Artifact Ownership (Clear & Correct)
- **Charter Agent** → **Project Charter / One-Pager** (contains NorthStar + MVP Features only)
- **Architect Agent** → **Architectural & Design Document (ADD)** (creates Requirements + Use Cases + technical design)
- **Planning Agent** → **Project Plan + AGENTS.md + Test Plan**
