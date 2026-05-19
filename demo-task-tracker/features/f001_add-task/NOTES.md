# Notes: Add Task

## Decisions
- ID generation uses placeholder `fmt.Sprintf("tsk-%d", 1)` — will be replaced with real `time.UnixNano()` + random in Task 3

## Deviations from Plan
- `io.Reader`/`io.Writer` interfaces for Store: deferred. Current tests use temp files directly which is simpler and sufficient. Interfaces add complexity without clear benefit for this scale.

## Future Improvements
- Consider `io.Reader`/`io.Writer` if Store needs to support multiple backends later
- ID generation should use `time.UnixNano()` + `rand.Int63()` for uniqueness

## Known Issues
- None
