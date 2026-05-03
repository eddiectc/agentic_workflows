---
description: Review implementation against spec and plan quality checks
argument-hint: "<feature-dir>"
---
Load the agile-workflow skill and execute Phase 5b: Implementation Review.

Feature directory: $@

Read the SPEC.md, PLAN.md, NOTES.md in that directory, and the actual implementation code. Conduct a structured review checking:

1. **Spec coverage** — Every scenario in SPEC.md is implemented and testable
2. **Plan fidelity** — Tasks completed match what was planned; deviations documented in NOTES.md
3. **Test quality** — Tests cover happy paths, error paths, and edge cases from the spec
4. **Code conventions** — Implementation follows docs/CONVENTIONS.md and existing patterns
5. **Edge cases** — Edge cases listed in the spec are handled
6. **No scope creep** — Nothing outside the spec's scope was added without approval
7. **No TODOs/debt** — No unresolved TODOs, temporary workarounds, or known issues

Present findings as a structured report with pass/fail per category and specific actionable items. Wait for my confirmation before suggesting any fixes.
