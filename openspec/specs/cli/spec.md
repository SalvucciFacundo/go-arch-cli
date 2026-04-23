# CLI Specification: go-arch

## Purpose
A command-line tool to scaffold Go projects with clean architecture and validate project health.

## Core Standards (Samber Upgrade)

### 1. Error Handling
- **Library**: `github.com/samber/oops`.
- **Pattern**: Every error returned from internal logic MUST be wrapped with context and a machine-readable code.
- **Root Handling**: The `RootCmd` executor handles final error reporting and exiting.

### 2. User Interface
- **Library**: `internal/ui`.
- **Helpers**: 
  - `Success`, `Warning`, `Error`, `Info` for standard lines.
  - `Analyzing` for long-running checks.
  - `*Msg` versions (e.g., `InfoMsg`) for inline colored text.
- **Aesthetics**: Bold colors, icons, and structured output.

### 3. CLI Framework
- **Framework**: Cobra.
- **Configuration**: Viper (YAML based).
- **Interactive Prompts**: Survey.
- **UX Rules**: 
  - `SilenceUsage: true` and `SilenceErrors: true` in `RootCmd`.
  - Manual error reporting via `ui.Fatal`.

## Commands
- `check`: Validates architectural rules.
- `generate`: Scaffolds components (services, repos, handlers).
- `new`: Interactive project initialization wizard.
- `serve`: Runs the project with hot-reload (Air support).
- `setup`: Environment preparation and tool installation.
