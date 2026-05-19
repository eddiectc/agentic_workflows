# Coding Conventions

## Code Style
- Follow `gofmt` / `goimports` formatting
- Max 100 lines per function; refactor if exceeded
- Error handling: return errors, don't panic (except in `main` setup)

## File Organization
- `cmd/` — entry points
- `internal/` — private application code
- One type/concept per file when possible

## Naming
- Commands: lowercase, hyphenated (`add-task`, `list-tasks`)
- Variables: lowercase, camelCase
- Types: PascalCase, singular noun
- Tests: `Test<Function>_<Scenario>`

## Testing
- Table-driven tests for all command logic
- Test file alongside source: `store.go` → `store_test.go`
- Mock file I/O by using interfaces (`io.Reader`, `io.Writer`)

## Git
- Conventional commits: `feat:`, `fix:`, `chore:`, `test:`
- One feature per branch
