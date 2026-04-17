# Command Reference Guide 📖

This document explains every available command in **Go-Architect CLI**, including usage examples and expected technical behavior.

---

## 1. `setup`
**Purpose**: Prepares the local development environment.
It detects the underlying Operating System (Linux, Windows, macOS) and suggests the correct installation steps for the Go toolchain and essential utilities like `air`.

### Example
```bash
go-arch setup
```
**Output**: 
- OS Detection report.
- Direct links to official installers.
- One-click commands to install the `air` hot-reload engine.

---

## 2. `new`
**Purpose**: Bootstraps a brand-new project from scratch.
It initiates an interactive terminal wizard (powered by `survey.v2`) to gather project metadata.

### Syntax
```bash
go-arch new [project-name]
```

### Options provided by the wizard:
- **Module Name**: The Go namespace (e.g., `github.com/user/repo`).
- **Architecture**: Choice between **Minimalist**, **Standard**, or **Hexagonal**.
- **DB Driver**: Pre-configures specific repository boilerplate (PostgreSQL, MySQL, MongoDB).

### Result:
A fully initialized project with a `.go-arch.yaml` manifest.

---

## 3. `generate`
**Purpose**: Scaffolds business components into an existing project.
Matches the project's architecture (stored in metadata) to determine the target path.

### Syntax
```bash
go-arch generate [component-type] [name]
# Shorthand alias:
go-arch g service Product
```

### Supported Components:
- **`service`**: Business logic layer.
- **`repository`**: Data access layer (generates both interface and implementation).
- **`handler`**: Entry point layer (HTTP/gRPC/CLI handlers).

### Folder Mapping Logic:
| Component Type | Standard Layout | Hexagonal Layout |
|----------------|-----------------|------------------|
| service        | `internal/service/` | `internal/domain/` |
| repository     | `internal/repository/` | `internal/ports/` |
| handler        | `internal/handler/` | `internal/adapters/` |

---

## 4. `serve`
**Purpose**: Starts the application with a production-grade development loop.

### Behavior:
1. **Air Detection**: If `air` is installed on the system, the CLI launches it automatically to provide **Hot-Reload**.
2. **Fallback**: If `air` is missing, it falls back to `go run`, ensuring the project always runs.

---

## 5. Metadata Manifest (`.go-arch.yaml`)
Every project includes this file. It allows the CLI to remain "stateless" but project-aware.

### Sample Content:
```yaml
project_name: MyApp
module_name: github.com/user/myapp
architecture: Hexagonal
db_driver: PostgreSQL
generated_at: 2026-04-17
```
*Note: This file must remain in the project root for `generate` and `serve` to function correctly.*
