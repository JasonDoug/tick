**NorthStar**  
→ **Feature**  
→ **Requirements** (EARS)  
→ **Use Case**  
→ **Sprint**  
→ **Task**  
→ **Sub-task / Prompt**
### Agentic EARS SDLC Hierarchy  
**(Least Granular / Most Abstract → Most Granular / Most Concrete)**

- **NorthStar** (Top of the pyramid – Least Granular / Most Abstract)  
  The single, unifying vision or mission for the entire project.  
  Everything below must support and flow up into this one idea.  
  *Example*: “Build an intelligent voice-first fitness companion that turns spoken workout notes into actionable insights and long-term progress tracking.”

- **Feature** (A – MVP-level capabilities)  
  Distinct, user-visible slices of functionality that must be present in the MVP.  
  These are written in colloquial, natural language – how users actually describe what they want.  
  One NorthStar contains **many Features**.  
  *Example*: “Voice Memo Recording & Transcription”, “AI Summarization of Voice Memos”, “Workout Progress Tracking”.

- **Requirements** (B)  
  Precise, verifiable EARS-style statements that define **what** the system must do (functional and non-functional).  
  These are created after the One-Pager is approved.  
  One Feature contains **many Requirements**.  
  *Example*: “When the user records a voice memo, the system shall transcribe it using the Whisper API with at least 95% accuracy for clear speech.”

- **Use Case** (C)  
  Detailed behavioral flows describing **how** an actor interacts with the system to achieve a goal (main flow, alternate flows, preconditions, postconditions, exceptions).  
  Use Cases elaborate the Requirements.  
  One Requirement is typically realized by **one or more Use Cases**.

- **Sprint / Phase** (D)  
  Logical or time-bound grouping of development work.  
  One Sprint contains **many Tasks**.

- **Task** (E)  
  Technical milestone or buildable component.  
  One Task contains **many Sub-tasks**.

- **Sub-task / Prompt** (F – Most Granular / Most Concrete)  
  The atomic, single-purpose prompt that is fed directly to an AI coding tool.  
  This is the actual executable unit of work.

### Containment Summary (Many-to-One)
- Many **Features** live inside **1 NorthStar**  
- Many **Requirements** live inside **1 Feature**  
- Many **Use Cases** live inside **1 Requirement** (or across related Requirements)  
- Many **Sprints** are derived from Use Cases  
- Many **Tasks** live inside **1 Sprint**  
- Many **Sub-tasks/Prompts** live inside **1 Task**

### How Artifacts Map to the Hierarchy
- **Project Charter / One-Pager** (produced by Charter Agent) → Contains **NorthStar** + list of **Features** (plus success metrics, scope, high-level notes).  
  → **No Requirements** here (per your latest clarification).

- **Architectural & Design Document (ADD)** (produced by Architect Agent) → Creates **Requirements** (EARS) + **Use Cases** + technical design.

- **Project Plan** (produced by Planning Agent) → Breaks down into **Sprints → Tasks → Sub-tasks/Prompts**.

- **AGENTS.md** and **Test Plan** → Grounding and quality enforcement across all levels.
