---
description: Review implementation against spec, plan, and DoD checklist
argument-hint: "<feature-dir> [task-number]"
---
Load the agile-workflow skill and execute Phase 4b: Implementation Review.

Feature directory: $1
$2 = task number (optional) — if provided, review only this specific task; otherwise review the full feature.

Read SPEC.md, PLAN.md, NOTES.md in the feature directory, read DoD.md at the project root, and read the actual implementation code.

If a task number was given:
- Locate the task in PLAN.md by number
- Identify which spec scenario(s) it corresponds to
- Scope spec coverage and test quality checks to that task only
- Still check all DoD.md items

Conduct a structured review checking:

1. **Spec coverage** — Every scenario (or task's scenarios) is implemented and testable
2. **Plan fidelity** — Tasks completed match what was planned; deviations documented in NOTES.md
3. **Test quality** — Tests cover happy paths, error paths, and edge cases from the spec
4. **Code conventions** — Implementation follows docs/CONVENTIONS.md and existing patterns
5. **Edge cases** — Edge cases listed in the spec are handled
6. **No scope creep** — Nothing outside the spec's scope was added without approval
7. **No TODOs/debt** — No unresolved TODOs, temporary workarounds, or known issues
8. **NOTES.md current** — All deviations and decisions are documented
9. **DoD checklist** — All applicable items from DoD.md are satisfied

Present findings as a structured report with pass/fail per category and specific actionable items. Wait for my confirmation before suggesting any fixes.
