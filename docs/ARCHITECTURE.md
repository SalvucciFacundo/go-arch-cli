# Project Architecture: {{ .Architecture }} 🏛️

This document describes the architectural patterns and layout of the **{{ .ProjectName }}** project.

## Project Structure Overview

Based on your selection, this project follows the **{{ .Architecture }}** pattern.

{{ if eq .Architecture "Hexagonal" }}
### ⬢ Hexagonal Architecture (Ports & Adapters)
- **internal/domain/**: Contains business logic, entities, and services. It is the core of the application and has no dependencies on external layers.
- **internal/ports/**: Defines interfaces (contracts) for external dependencies (repositories, external APIs).
- **internal/adapters/**: Contains concrete implementations (SQL repositories, HTTP handlers).
{{ else if eq .Architecture "Standard" }}
### 📦 Standard Layout
- **cmd/api/**: Entry point of the application.
- **internal/handler/**: Request/Response handling.
- **internal/service/**: Core business logic.
- **internal/repository/**: Data access implementation.
{{ else }}
### ⚡ Minimalist Layout
A flat and lean structure for microservices or single-file tools.
{{ end }}

## Database
- **Driver**: {{ .DBDriver }}

## Infrastructure
- **Docker Support**: {{ if .UseDocker }}Enabled (Dockerfile & Compose generated){{ else }}Disabled{{ end }}
