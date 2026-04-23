# Tasks: Samber Upgrade Refactor

## Phase 1: Infrastructure
- [x] Add `github.com/samber/oops` dependency.
- [x] Create `internal/ui/output.go` for centralized CLI output handling.

## Phase 2: Core Refactor
- [x] Refactor `cmd/root.go` with silence flags and centralized error handling.
- [x] Refactor Viper config binding for robustness.

## Phase 3: Command Migration
- [x] Refactor `cmd/check.go`.
- [x] Refactor `cmd/generate.go`.
- [x] Refactor `cmd/new.go`.
- [x] Refactor `cmd/serve.go`.
- [x] Refactor `cmd/setup.go`.

## Phase 4: Internal Logic
- [x] Update `internal/pkg/validator` with error wrapping.

## Phase 5: Verification
- [x] Run full test suite.
- [x] Manual CLI verification.
