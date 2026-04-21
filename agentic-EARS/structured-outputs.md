Finalized Structured Output Templates + Examples

`#### 1\. Project Charter / One-Pager (Charter Agent)
`
**JSON Template** (Recommended for LLM output):

JSON

Copy

```
{
  "document_type": "Project_Charter_OnePager",
  "version": "1.0",
  "created_date": "2026-04-21",
  "northstar": "string",
  "elevator_pitch": "string",
  "problem_solved": "string",
  "target_users": "string",
  "features": [
    {
      "id": "F-001",
      "name": "string",
      "description": "string"
    }
  ],
  "success_metrics": ["string"],
  "scope": {
    "in_scope": ["string"],
    "out_of_scope": ["string"]
  },
  "non_functional_notes": "string",
  "status": "draft | approved"
}
```

**Example** (using Agentic EARS itself):

JSON

Copy

```
{
  "document_type": "Project_Charter_OnePager",
  "version": "1.0",
  "created_date": "2026-04-21",
  "northstar": "Turn vague human ideas into fully traceable, production-ready code through a repeatable chain of specialized agentic LLM agents.",
  "elevator_pitch": "Agentic EARS is a structured agentic pipeline that converts vague ideas into clean documents and atomic coding prompts with zero context drift.",
  "problem_solved": "Current LLM-driven development suffers from severe context drift, scope creep, and inconsistent outputs because there is no enforced refinement and traceability process.",
  "target_users": "Solo indie developers, AI builders, and small teams using Cursor, Copilot, or similar tools who want a repeatable, low-chaos development workflow.",
  "features": [
    {
      "id": "F-001",
      "name": "Charter Agent",
      "description": "Converts a vague Seed Idea into an approved high-level Project Charter / One-Pager"
    },
    {
      "id": "F-002",
      "name": "Architect Agent",
      "description": "Creates EARS Requirements, Use Cases, and full technical design"
    }
  ],
  "success_metrics": [
    "Go from vague idea to working code in under 3 agent sessions",
    "Zero context drift between agents",
    "Full traceability from Sub-task back to NorthStar"
  ],
  "scope": {
    "in_scope": ["Agent pipeline", "Document generation", "Atomic prompts", "Grounding via AGENTS.md"],
    "out_of_scope": ["Actual code execution engine", "Full web UI", "Multi-user collaboration"]
  },
  "non_functional_notes": "Lightweight, Python-native, works with any LLM, strict approval gates and traceability",
  "status": "approved"
}
```

**Structured Output Template** (JSON – always return this exact schema when generating the artifact)
```json
{
  "document_type": "Project_Charter_OnePager",
  "northstar": "string - one concise sentence",
  "features": [
    {
      "id": "F-001",
      "name": "string - short colloquial name",
      "description": "string - 1-2 sentence user-friendly description"
    }
  ],
  "success_metrics": ["string", "string"],
  "scope": {
    "in_scope": ["string", "string"],
    "out_of_scope": ["string", "string"]
  },
  "non_functional_notes": "string - high-level constraints",
  "status": "draft | approved",
  "version": "1.0"
}
```
```json
{
  "document_type": "Project_Charter_OnePager",
  "northstar": "Build an intelligent voice-first fitness companion that turns spoken workout notes into actionable insights and long-term progress tracking.",
  "features": [
    { "id": "F-001", "name": "Voice Memo Recording & Transcription", "description": "..." },
    { "id": "F-002", "name": "AI Summarization of Voice Memos", "description": "..." }
  ],
  "success_metrics": ["End-to-end flow under 30 seconds", "Transcription accuracy ≥ 95%"],
  "scope": { "in_scope": ["Voice recording", "Transcription", "AI summarization"], "out_of_scope": ["Social sharing"] },
  "non_functional_notes": "Low latency, private & encrypted voice data",
  "status": "approved",
  "version": "1.0"
}
```
#### 2\. Architectural & Design Document (ADD) – Core JSON Structure

JSON

Copy

```
{
  "document_type": "Architectural_Design_Document",
  "version": "1.0",
  "requirements": [
    {
      "id": "R-F001-001",
      "feature_id": "F-001",
      "ears_statement": "When the user provides a Seed Idea, the system shall ...",
      "type": "functional",
      "priority": "must_have"
    }
  ],
  "use_cases": [
    {
      "id": "UC-001",
      "name": "Generate Project Charter",
      "realizes_requirements": ["R-F001-001"],
      "preconditions": "...",
      "main_flow": ["Step 1", "Step 2"],
      "alternate_flows": [],
      "postconditions": "...",
      "exceptions": []
    }
  ],
  "architecture": {
    "tech_stack": ["Python", "Any LLM"],
    "modules": ["charter_agent", "architect_agent", "planning_agent"],
    "data_models": ["ProjectCharter", "ADD", "ProjectPlan"]
  },
  "status": "approved"
}
```

#### 3\. Project Plan (Core Structure)

JSON

Copy

```
{
  "document_type": "Project_Plan",
  "version": "1.0",
  "sprints": [
    {
      "id": "S-001",
      "name": "Foundation",
      "tasks": [
        {
          "id": "T-001",
          "name": "Implement Charter Agent",
          "sub_tasks": [
            {
              "id": "ST-001-1",
              "prompt": "Write the Charter Agent function that...",
              "definition_of_done": "Produces valid One-Pager JSON matching schema and passes self-checklist"
            }
          ]
        }
      ]
    }
  ]
}
```

#### 4\. AGENTS.md

This one is **Markdown only** (not JSON). It should contain:

*   NorthStar + Elevator Pitch summary
*   Target Users
*   Core Guardrails
*   Architecture References
*   Coding Standards
*   Do’s and Don’ts
*   Version references

#### 5\. Test Plan

Also **Markdown**, with sections for:

*   Testing Strategy
*   Acceptance Criteria per Requirement / Use Case
*   Test Generation Guidelines

* * *
