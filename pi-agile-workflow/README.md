# pi-agile-workflow

Agile workflow package for [pi](https://pi.dev) — structured feature development with BDD specs, implementation plans, progress tracking, and retrospectives. Designed for agentic coding where you decide and the agent executes.

## Install

```bash
# From this directory (project-local)
cd /path/to/your/project
pi install git:/home/tc/sr.ht/pi_setup/pi-agile-workflow -l

# Or copy the package files into your project
cp -r /home/tc/sr.ht/pi_setup/pi-agile-workflow/skills/. .pi/skills/
cp -r /home/tc/sr.ht/pi_setup/pi-agile-workflow/prompts/. .pi/prompts/
```

## Workflow

```
/write-spec → /review-spec → /plan-impl → implement → /review-impl → /retro
```

Each phase requires your approval before moving forward. The agent drafts, you decide.

Use `/setup-project` as a prerequisite to bootstrap project structure, conventions, and Definition of Done.

## Commands

| Command | Purpose |
|---|---|
| `/setup-project` | Define goals, non-goals, architecture, tech stack, and DoD.md (prerequisite, not a workflow phase) |
| `/write-spec <feature>` | Draft a BDD feature spec (implementation-agnostic) |
| `/review-spec <feature>` | Review a feature spec for gaps, ambiguity, completeness |
| `/plan-impl <feature>` | Create implementation plan with end-to-end testable tasks |
| `/continue` | Resume current sprint — pick up first unfinished task |
| `/review-impl <feature> [task]` | Review full feature, or a single task by number, against spec/plan/DoD |
| `/retro <feature>` | Sprint retrospective — what went well, what to improve |

## Project Structure

```
project/
├── .pi/
│   ├── skills/
│   │   └── agile-workflow/
│   │       └── SKILL.md
│   └── prompts/
│       ├── setup-project.md
│       ├── write-spec.md
│       ├── review-spec.md
│       ├── plan-impl.md
│       ├── continue.md
│       └── retro.md
├── DoD.md                  # Definition of Done checklist
├── docs/
│   ├── PROJECT.md          # Goals, non-goals, architecture, tech stack
│   └── CONVENTIONS.md      # Coding conventions, patterns
├── features/
│   ├── f001_user-auth/
│   │   ├── SPEC.md         # BDD feature spec
│   │   ├── PLAN.md         # Implementation tasks with checkboxes
│   │   └── NOTES.md        # Deviations, decisions, future improvements
│   └── ...
├── CHANGELOG.md
└── README.md
```

## Definition of Done (DoD.md)

Created during `/setup-project`, `DoD.md` is a project-level checklist that defines what "done" means for every feature. The agent proposes a default based on the project type; you customize it to match your quality bar. It is checked during:

- **Per-task self-check** — during implementation, each task is verified against applicable DoD items
- **`/review-impl`** — the full DoD checklist is a review criterion alongside spec coverage, plan fidelity, etc.

## Principles

- **You decide, agent executes** — agent drafts specs/plans, you review and approve
- **Small increments** — each task is end-to-end testable and verifiable
- **Agent self-checks per task** — lightweight quality checklist runs silently during implementation
- **Explicit gate reviews** — `/review-impl` gives you a comprehensive check before closing a feature (or a single task)
- **Follow existing patterns** — only introduce new approaches with your confirmation
- **Challenge over-engineering** — simplest solution that works is preferred
- **Files are state** — progress, decisions, plans, and quality criteria live in version-controlled files
