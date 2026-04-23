# Exploration Report: Refactor Samber Upgrade

## Current State Analysis

- **Cobra Implementation**: Basic. Lacks `SilenceUsage` and `SilenceErrors` on RootCmd, leading to suboptimal UX when errors occur.
- **Configuration**: Uses Viper but without explicit flag binding in many commands. `initConfig` is standard but could be more robust.
- **Error Handling**: Uses standard `fmt.Errorf` and `fmt.Println`. No context-rich error wrapping.
- **UI/UX**: Directly prints to stdout/stderr using `fmt`. No abstraction for different output levels (Success, Warning, Error).
- **Go Version**: 1.24.0. Project is well-positioned to use modern features.

## Samber Standards Gap

| Area | Samber Standard | Current Status |
|------|-----------------|----------------|
| **CLI UX** | SilenceUsage/Errors = true | ❌ Not set |
| **Testing** | Use `testify` | ✅ Already in go.mod (some) |
| **Errors** | Use `oops` wrapping | ❌ Standard only |
| **Config** | BindPFlag mandatory | ❌ Partial implementation |
| **Logic** | Clean separation of UI | ❌ Direct printing in commands |

## Recommendations

1. **Centralize UI**: Create `internal/ui/output.go` for all CLI interactions.
2. **Standardize Commands**: Create a base command pattern or helper to apply the silence flags and common PreRun logic.
3. **Enhance Errors**: Integrate `samber/oops` for all internal logic.
4. **Modernize loops**: Audit `internal/` for legacy loop patterns.
