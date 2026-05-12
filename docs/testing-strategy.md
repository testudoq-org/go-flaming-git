# Testing Strategy

## Test Types

### Unit Tests
- Test individual functions and methods
- Do not depend on external systems
- Use table-driven tests with `t.Run` for clarity
- Keep tests deterministic and fast
- Assert behavior, not implementation details

### Integration Tests
- Test interactions between packages or modules
- Cover boundary behavior and data flow
- Use stubs, fixtures, or test containers when needed
- Keep integration tests reliable and repeatable

### End-to-End Tests
- Validate complete workflows and user-facing behavior
- Run in CI for critical paths only
- Keep end-to-end tests stable and focused

## Tooling
- `go test ./...`
- `go test -coverprofile=coverage.out ./...`
- `go run ./cmd/cover2lcov -in coverage.out -out coverage.lcov`
- `go test -race ./...` for concurrency-sensitive code
- `golangci-lint run`
- `go vet ./...`
- build `crap` from `https://github.com/testudoq-org/go-crap4go` and run `crap --no-run-tests --coverprofile=coverage.out --threshold=15`
- `go mod tidy` and `go mod verify`

## Coverage Rules
- Minimum repository coverage: 85%
- Critical modules should aim for 95%+
- Measure coverage with `go test -coverprofile`
- Use CI enforcement for coverage thresholds

## CRAP Rules
- Tool source: `https://github.com/testudoq-org/go-crap4go`
- Tolerance zone is below 15
- Any function with CRAP score 15 or above fails the quality gate
- Run CRAP after generating `coverage.out`
- Enforce the threshold in local quality checks and CI
- Enforced CRAP scope is runtime application paths: `cmd/fgit`, `internal/alias`, `internal/dispatch`, and `internal/runner`

## LCOV Rules
- Generate `coverage.lcov` from `coverage.out` in local and CI quality runs
- Publish `coverage.lcov` as a quality artifact for downstream reporting

## CRAP Failure Process
- Capture the failing report in `crap-report.txt`
- Identify offending functions and rank by highest CRAP score
- Perform root cause analysis for each offender
- If the report is caused by analyzer coverage-mapping limits in non-runtime tooling packages, keep the runtime scope gate and continue tracking tooling quality with unit tests
- Add failing tests for missing risky paths first (TDD red)
- Apply minimal refactors to reduce complexity and improve testability (green + refactor)
- Re-run quality gates until no function remains at or above 15

## Best Practices
- Prefer table-driven tests for functions with multiple inputs
- Keep test setup minimal and readable
- Use helper functions only for repeated setup logic
- Name tests clearly and describe behavior in `t.Run` names
- Keep unit tests isolated from shared state
- Use mocks or fakes only when the dependency is outside the test boundary

## Mutation Testing (Optional but Recommended)
- Use mutation testing tools to validate test strength
- Good tests should fail when logic is altered
- Treat mutation testing as a way to catch weak assertions and brittle coverage