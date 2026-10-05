# agile-workflow

Agile workflow package for [maki](https://maki.sh) — structured feature development with BDD specs, implementation plans, progress tracking, and retrospectives. Designed for agentic coding where you decide and the agent executes.

## Install

### Agile workflow (commands & skills)

```bash
# Personal install — symlink into your maki config
ln -sfn /path/to/agile-workflow/commands ~/.config/maki/commands
ln -sfn /path/to/agile-workflow/skills/agile-workflow ~/.config/maki/skills/agile-workflow

# Or project-local — copy into your project
cp -r /path/to/agile-workflow/skills/. .maki/skills/
cp -r /path/to/agile-workflow/commands/. .maki/commands/
```

### Provider plugin (OpenCode Go Responses)

OpenCode Go serves some models only on `/responses`, but maki's built-in
`opencode-go` provider posts every model to `/chat/completions`, so those
fail with `ModelProtocolUnsupported`. This package ships a companion
provider for them.

Install the repo as a maki package — this loads the provider plugin:

```lua
-- ~/.config/maki/init.lua
maki.pack.add({
  { src = "https://github.com/eddiectc/agentic_workflows", version = "main" },
})
```

Then authenticate and pick a model:

```bash
export OPENCODE_API_KEY=...            # or: maki auth login opencode-go-responses
maki -p -m opencode-go-responses/grok-4.7 hi
```

The hand-linked and global-config installs, the model list, and the
maintenance notes are in
[OpenCode Go (Responses) provider plugin](#opencode-go-responses-provider-plugin).

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
├── .maki/
│   ├── skills/
│   │   └── agile-workflow/
│   │       └── SKILL.md
│   └── commands/
│       ├── setup-project.md
│       ├── write-spec.md
│       ├── review-spec.md
│       ├── plan-impl.md
│       ├── continue.md
│       ├── review-impl.md
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

## OpenCode Go (Responses) provider plugin

Some OpenCode Go models are served only on the OpenAI **Responses** API
(`/v1/responses`), not `/v1/chat/completions`:

- `grok-4.7`, `grok-4.6`
- `gpt-6-luna`, `gpt-5.6-luna`

maki derives one wire format per models.dev provider, and models.dev marks
`opencode-go` as `@ai-sdk/openai-compatible`, so it posts these to
`/chat/completions` and the gateway rejects them with
`ModelProtocolUnsupported`. A Lua plugin cannot claim the `opencode-go` slug
(models.dev owns it), so this package registers a companion provider that
speaks the responses codec against the same gateway.

The plugin lives at `plugin/opencode_go_responses.lua`; `plugin.toml` grants
it `net` (pinned to `opencode.ai`) and `env` (to read `OPENCODE_API_KEY`).

### Install as a maki package

```lua
-- ~/.config/maki/init.lua
maki.pack.add({
  { src = "https://github.com/eddiectc/agentic_workflows", version = "main" },
})
```

Or link the checkout into maki's package dir by hand, which also loads it
(`<maki-data>` is `~/.local/share/maki` on Linux):

```bash
ln -sfn /path/to/agentic_workflows \
  <maki-data>/site/pack/mine/start/agentic_workflows
```

### Install from the global config (iterative editing)

maki's global config runs only `init.lua`, so symlink the module and require
it, and make sure the permissions and `net_hosts` are present:

```bash
mkdir -p ~/.config/maki/lua
ln -sfn /path/to/agentic_workflows/plugin/opencode_go_responses.lua \
  ~/.config/maki/lua/opencode_go_responses.lua
```

```lua
-- ~/.config/maki/init.lua
require("opencode_go_responses")
```

```toml
# ~/.config/maki/plugin.toml
[permissions]
net = true
env = true
net_hosts = ["opencode.ai"]
```

Run `/reload`, then check `~/.local/logs/maki/maki.log` if it does not load.

### Authenticate and use

The plugin reads the same `OPENCODE_API_KEY` the built-in provider uses. If
you authenticated with `maki auth login opencode-go` instead, run
`maki auth login opencode-go-responses` with the same key (a plugin cannot
read a built-in slug's stored credentials).

```bash
export OPENCODE_API_KEY=...            # or: maki auth login opencode-go-responses
maki -p -m opencode-go-responses/grok-4.7 hi
```

Select `opencode-go-responses/<model>` in `/model` or pass it to `--model`.
`opencode-go/grok-4.7` keeps failing: maki routes that slug by models.dev and
there is no per-model protocol override.

Note: `muse-spark-1.2/1.3-contributor` also use `/responses`. They are not
included here because they fail on a workspace privacy setting, not the
protocol — add their ids to `RESPONSES_MODELS` if that is enabled for you.

### Session header (`x-opencode-session`)

OpenCode Go requires a stable `x-opencode-session` on every request and now
rejects requests without it:

```
API error (400): {"type":"error","error":{"type":"MissingSessionID","message":"Request is missing x-opencode-session ..."}}
```

maki sends that header only for its built-in `opencode`/`opencode-go` slugs,
where it is derived from the conversation's session id, and the responses codec
ignores the `openai.session_id` option. This plugin therefore adds the header
itself, through its `auth` hook — the only hook whose headers reach a responses
request. The hook copies the resolved `Authorization` header and adds the
session id; no extra configuration is needed.

maki hands the auth hook no session id and runs it once, so the id is **one
stable value per maki process**, not per conversation. It survives `/reload`
(the credential cell is reused) and changes on restart. That satisfies the
gateway; the only thing weaker than the built-in provider is prompt-cache
affinity across concurrent sessions. True per-conversation ids would need an
upstream maki change (apply `request_headers(session_id)` on the responses
codec path).

To change the id, or to persist one across restarts, edit `SESSION_ID` in
`plugin/opencode_go_responses.lua`.
