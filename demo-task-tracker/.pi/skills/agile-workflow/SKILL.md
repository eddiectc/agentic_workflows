---
name: agile-workflow
description: Structured agile workflow for agentic coding. Use for project setup, writing BDD feature specs, implementation planning, execution, and retrospectives. Covers both new and existing projects. Invoke when the user mentions features, specs, sprints, planning, or project setup.
---

# Agile Workflow for Agentic Coding

Structured workflow where **the user decides and the agent executes**. Every phase requires user confirmation before proceeding.

## Core Principles (Apply in Every Phase)

1. **Ask before deciding** — Never assume. If anything is ambiguous, ask the user for clarification. Present options with pros and cons, then wait for confirmation.
2. **Follow existing patterns** — Look at how things are already done in the project. Follow that way first. Only suggest deviations if the existing approach is clearly not best practice or over-complicated.
3. **Challenge over-engineering** — If a solution seems complex, ask "what's the simplest thing that could work?" Prefer simple, incremental solutions.
4. **Small increments** — Break work into end-to-end testable pieces. Each task should be independently verifiable.
5. **Document deviations** — If implementation drifts from the plan, explain why. If valid, update the plan. If not, document it in NOTES.md as future work.

## Workflow Phases

```
/setup-project → /write-spec → /review-spec → /plan-impl → implement → /retro
```

Each phase has a **gate** — the user must review and approve before moving to the next phase. Never skip a gate.

---

## Phase 1: Project Setup (`/setup-project`)

### New Project

When setting up a new project:

1. Ask the user for:
   - Project name and purpose
   - Target users / audience
   - High-level goals (what the project should achieve)
   - Non-goals (what it explicitly should NOT do)
   - Any technical constraints (language, platform, hosting, budget)
   - Any existing systems it must integrate with

2. Propose:
   - Technical architecture (with alternatives and pros/cons)
   - Tech stack (justify each choice, follow user preferences)
   - Project directory structure

3. After user confirmation, create:

**docs/PROJECT.md:**
```markdown
# Project: <name>

## Description
<what this project does, for whom>

## Goals
- <goal 1>
- <goal 2>

## Non-Goals
- <explicitly out of scope 1>
- <explicitly out of scope 2>

## Architecture
<high-level architecture diagram or description>

## Tech Stack
| Component | Choice | Reason |
|---|---|---|
| Language | <lang> | <reason> |
| Framework | <fw> | <reason> |
| Database | <db> | <reason> |
| <etc> | | |

## Constraints
- <technical or business constraints>

## Directory Structure
```
<show the project layout>
```

## Conventions
<Coding conventions, naming patterns, file organization>
```

**docs/CONVENTIONS.md:**
```markdown
# Coding Conventions

## Code Style
<language-specific style guide reference>

## File Organization
<how files are structured>

## Naming
<naming conventions>

## Testing
<testing approach, framework, conventions>

## Git
<branch strategy, commit conventions>
```

**features/README.md:**
```markdown
# Features

Each feature has its own directory: `features/<id>_<name>/`

- `SPEC.md` — BDD feature spec (what, implementation-agnostic)
- `PLAN.md` — Implementation plan with tasks (how, with checkboxes for progress)
- `NOTES.md` — Deviations, decisions, future improvements

Feature IDs are sequential: f001, f002, etc.

## Feature Index

| ID | Name | Status |
|---|---|---|
| f001 | <name> | <planned | spec | planning | in-progress | done> |
```

4. Create the directory structure:
```
mkdir -p docs features
```

5. Update or create project `AGENTS.md` (or `CLAUDE.md`) with:
```markdown
# Project Conventions

This project follows an agile workflow for feature development.

## Feature Structure
Features live in `features/<id>_<name>/` with SPEC.md, PLAN.md, NOTES.md.

## Workflow
1. Write feature spec (SPEC.md) — user reviews before proceeding
2. Create implementation plan (PLAN.md) — user reviews before coding
3. Implement tasks incrementally, update checkboxes in PLAN.md
4. Document deviations in NOTES.md
5. Retrospective after feature completion

## Principles
- Follow existing patterns before introducing new approaches
- Each task is end-to-end testable
- Ask for clarification on ambiguous requirements
```

### Existing Project

When working on an existing project:

1. Analyze the existing codebase:
   - Read README.md, existing docs, config files
   - Understand the directory structure and patterns
   - Identify the tech stack and conventions
   - Look for existing feature/spec documentation

2. Present findings to the user:
   - Summary of what the project does
   - Detected tech stack and architecture
   - Observed coding conventions
   - Any inconsistencies or areas for improvement

3. Ask the user:
   - What features they want to add or change
   - Whether the detected conventions match their intent
   - Any corrections to the analysis

4. Create `docs/PROJECT.md` if it doesn't exist (populate from analysis)
5. Create `features/` directory and `features/README.md` if they don't exist
6. Proceed to the next phase based on user's feature requests

---

## Phase 2: Feature Spec (`/write-spec <feature-name>`)

Write an **implementation-agnostic** BDD feature spec. The spec describes **what** the system should do, not **how** to build it.

### Process

1. Ask the user for:
   - Feature name and brief description
   - User stories or use cases
   - Any constraints or requirements
   - What is explicitly NOT in scope

2. Draft SPEC.md using the template below

3. Present the spec to the user for review — **wait for approval** before proceeding

### SPEC.md Template

```markdown
# Feature: <name>

## Description
<what this feature does, for whom, why it matters>

## User Stories
<as a <role>, I want <action>, so that <benefit>>

## Scenarios

### Scenario: <descriptive name>
**Given** <initial context / preconditions>
**When** <action / event>
**Then** <expected outcome>
**And** <additional outcomes if any>

### Scenario: <another scenario>
**Given** ...
**When** ...
**Then** ...

## Edge Cases
- <edge case 1: what happens when ...>
- <edge case 2: what happens when ...>

## Constraints
- <technical or business constraints>

## Non-Goals
- <explicitly out of scope for this feature>

## Dependencies
- <other features or systems this depends on, if any>
```

### Spec Quality Checklist

Before presenting the spec, verify:
- [ ] Each scenario is testable (has clear Given/When/Then)
- [ ] Scenarios cover happy path and error paths
- [ ] Edge cases are identified
- [ ] Non-goals are explicit (prevents scope creep)
- [ ] No implementation details leaked into the spec (no "use Redis" or "create a REST endpoint")
- [ ] Terminology is consistent throughout

---

## Phase 3: Spec Review (`/review-spec <feature-name>`)

Review an existing SPEC.md for quality and completeness.

### Process

1. Read the SPEC.md
2. Check against the quality checklist above
3. Identify:
   - Ambiguous scenarios (unclear Given/When/Then)
   - Missing scenarios (obvious cases not covered)
   - Implementation details that leaked into the spec
   - Scope issues (too broad or too narrow)
   - Missing edge cases
   - Missing constraints or non-goals
4. Present findings as a structured list
5. Suggest specific improvements
6. **Wait for user confirmation** before suggesting edits

---

## Phase 4: Implementation Plan (`/plan-impl <feature-name>`)

Create an implementation plan that breaks the feature into small, end-to-end testable tasks.

### Process

1. Read the approved SPEC.md
2. Read existing codebase to understand current patterns
3. Break implementation into tasks where each task:
   - Is independently testable
   - Delivers a complete slice of functionality
   - Can be verified against the spec
   - Is small enough to complete in one focused session
4. Order tasks by dependency (foundations first)
5. For each technical decision, follow existing project patterns. If a new approach is needed, present options with pros/cons and **get user confirmation**.
6. Draft PLAN.md using the template below
7. Present to user — **wait for approval** before implementation

### PLAN.md Template

```markdown
# Implementation Plan: <feature name>

## Overview
<Brief summary of what will be built>

## Task Dependencies
<Visual or textual dependency graph>
e.g. Task 1 → Task 2 → Task 3 (Task 2 and 4 can be parallel)

## Tasks

### Task 1: <name> [PRIORITY: HIGH|MEDIUM|LOW]
**Corresponds to:** Scenario(s) <names from spec>
**Description:** <what this task builds>

- [ ] <sub-step 1>
- [ ] <sub-step 2>
- [ ] <sub-step 3>
- [ ] Write tests

**Verification:** <how to verify this task is done>

### Task 2: <name> [PRIORITY: HIGH|MEDIUM|LOW]
...

## Technical Decisions
| Decision | Choice | Reason |
|---|---|---|
| <decision> | <choice> | <reason, referencing existing patterns> |

## Risks
- <potential risk and mitigation>
```

### Plan Quality Checklist

- [ ] Each task maps to one or more spec scenarios
- [ ] Tasks are ordered by dependency
- [ ] Each task has a clear verification criterion
- [ ] Tasks include testing
- [ ] Technical decisions reference existing project patterns
- [ ] No task is so large it can't be done in one focused session

---

## Phase 5: Implementation

Execute the implementation plan task by task.

### Process for Each Task

1. Read the PLAN.md and identify the next unchecked task
2. Read relevant existing code to understand context
3. Implement the task:
   - Write code following existing conventions
   - Write tests alongside or before code
   - Run tests to verify
4. Check off the task in PLAN.md (`- [ ]` → `- [x]`)
5. If you deviate from the plan:
   - Explain why to the user
   - If the deviation is valid, update PLAN.md
   - If it should be addressed later, document in NOTES.md
6. Move to the next task

### NOTES.md Template

Create or append to NOTES.md as needed:

```markdown
# Notes: <feature name>

## Decisions
- <date/turn>: <decision made and rationale>

## Deviations from Plan
- <what deviated, why, resolved or pending>

## Future Improvements
- <improvement noticed during implementation>

## Known Issues
- <any unresolved issues>
```

### During Implementation — Ongoing Rules

- **If existing code is not best practice:** Note it in NOTES.md as a future improvement. Don't refactor unless the user asks.
- **If the plan is wrong:** Stop, explain to the user, propose a revised plan. Don't silently diverge.
- **If you discover a better approach:** Present it with pros/cons. Get confirmation before changing direction.
- **If a task is bigger than expected:** Split it, update the plan, get confirmation.

---

## Phase 6: Retrospective (`/retro <feature-name>`)

After a feature is complete, conduct a retrospective.

### Process

1. Review:
   - SPEC.md (what was planned)
   - PLAN.md (how it was planned)
   - NOTES.md (what happened during implementation)
   - The actual code (what was built)

2. Draft RETRO.md using the template below

3. Present to user for review

### RETRO.md Template

```markdown
# Retrospective: <feature name>

## What Went Well
- <positive observation>

## What Could Be Improved
- <area for improvement>

## Spec vs Reality
- <did the spec match the implementation? gaps?>
- <were there scenarios the spec missed?>

## Plan vs Reality
- <was the task breakdown effective?>
- <were tasks sized appropriately?>
- <were dependencies accurate?>

## Learnings
- <what to do differently next time>

## Action Items
- [ ] <concrete improvement for next feature>
```

### After Retrospective

1. Update `features/README.md` feature index (mark feature as "done")
2. Suggest next feature or ask user what to work on next
3. If retro reveals process improvements, suggest updates to CONVENTIONS.md

---

## Resuming Work (`/continue`)

When resuming work on a feature:

1. Read `features/README.md` to find current feature
2. Read the feature's PLAN.md to find first unchecked task
3. Read NOTES.md for any context from previous sessions
4. Confirm with user: "Picking up Task N: <name>. Correct?"
5. Proceed with implementation from that task

---

## Quick Reference: File Purposes

| File | Purpose | Who Writes |
|---|---|---|
| `docs/PROJECT.md` | Project goals, architecture, tech stack | Agent drafts, user approves |
| `docs/CONVENTIONS.md` | Coding conventions | Agent drafts, user approves |
| `features/README.md` | Feature index with status | Agent maintains |
| `features/<id>/SPEC.md` | BDD feature spec (what) | Agent drafts, user approves |
| `features/<id>/PLAN.md` | Implementation tasks (how + progress) | Agent drafts, user approves, agent updates checkboxes |
| `features/<id>/NOTES.md` | Deviations, decisions, improvements | Agent maintains, user reviews |
| `features/<id>/RETRO.md` | Sprint retrospective | Agent drafts, user reviews |
