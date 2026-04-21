### 3. Planning Agent

**Purpose / Role**
Takes approved One-Pager + ADD → **Project Plan + AGENTS.md + Test Plan**.

**System Prompt Example** (full)
```
You are the Planning Agent in the Agentic EARS system.

Your job is to create three artifacts:
1. Project Plan (Sprints → Tasks → atomic Sub-tasks)
2. AGENTS.md (grounding file)
3. Test Plan

Every Sub-task prompt must reference AGENTS.md and the Test Plan.
Only generate the final artifacts after user says “APPROVE PLAN” or “GENERATE ALL ARTIFACTS”.
```

**Structured Output Template** (Project Plan JSON – core)
```json
{
  "document_type": "Project_Plan",
  "sprints": [
    {
      "id": "S-001",
      "name": "Foundation",
      "tasks": [
        {
          "id": "T-001-1",
          "name": "Implement transcription endpoint",
          "sub_tasks": [
            {
              "id": "ST-001-1-1",
              "prompt": "Using AGENTS.md and ADD section 3.2, write a FastAPI endpoint that ...",
              "definition_of_done": "Passes all tests from Test Plan, follows coding standards"
            }
          ]
        }
      ]
    }
  ],
  "status": "approved",
  "version": "1.0"
}
```

**AGENTS.md Template** (produced as Markdown)
```markdown
# AGENTS.md – Grounding File for Agentic EARS

**NorthStar**: [exact NorthStar text]

**MVP Features**:
- F-001: ...

**Key Guardrails**:
- All code must trace to a Requirement in the ADD.
- Sub-tasks must stay under 8000 tokens.
- Never add new Features without updating the One-Pager.

**References**:
- Project Charter: version 1.0
- ADD: version 1.0
- Test Plan: version 1.0
```

**Test Plan Template** (produced as Markdown)
```markdown
# Test Plan

**Acceptance Criteria per Requirement**:
- R-F001-001: ...

**Test Strategy**:
- Unit tests for every Sub-task
- Integration tests for Use Cases
```
