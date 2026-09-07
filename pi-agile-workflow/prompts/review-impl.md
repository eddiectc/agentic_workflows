---
description: Review implementation against spec, plan, and DoD checklist
argument-hint: "<feature-dir> <task-number>"
---
Load the agile-workflow skill and execute Phase 4b: Implementation Review.

Feature directory: $1
Task number: $2 — if "all", review the full feature; otherwise review only this specific task.

Read SPEC.md, PLAN.md, NOTES.md in the feature directory, read DoD.md at the project root, and read the actual implementation code.

If a task number was given:
- Locate the task in PLAN.md by number
- Identify which spec scenario(s) it corresponds to
- Scope spec coverage and test quality checks to that task only
- Still check all DoD.md items

Conduct a structured review checking:

**Workflow checks** (from the skill):
1. **Spec coverage** — Every scenario (or task's scenarios) is implemented and testable
2. **Plan fidelity** — Tasks completed match what was planned; deviations documented in NOTES.md
3. **NOTES.md current** — All deviations and decisions are documented

**Quality checks** (from DoD.md):
4. Check every item in the project's DoD.md checklist

Present findings as a structured report with pass/fail per category and specific actionable items. Wait for my confirmation before suggesting any fixes.
