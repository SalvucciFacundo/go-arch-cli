# Go-Architect CLI (go-arch) 🚀

**Go-Architect CLI** is a professional, agnostic, and multi-platform scaffolding tool designed to standardize Go project initialization. Inspired by the performance and modularity of the Angular CLI, it empowers developers to bootstrap production-ready applications with clean architecture patterns in seconds.

## ✨ Key Features

- 🏗️ **Architecture Layouts**: Native support for **Minimalist**, **Standard**, and **Hexagonal** (Ports & Adapters) architectures.
- 🔌 **Agnostic & Decoupled**: Data-driver independent (PostgreSQL, MySQL, MongoDB, No-SQL) and IDE-agnostic.
- ⚡ **Built-in Hot-Reload**: Seamless integration with `Air` for a high-performance development loop.
- 🛠️ **Component Generators**: Scaffold Services, Repositories, and Handlers that automatically map to your chosen architecture.
- 🧪 **QA & TDD Oriented**: Automatic test file generation with manual mocking patterns to foster high-quality codebases.
- 💻 **Cross-Platform Compatibility**: Full support for Windows (ANSI-enabled), Linux, and macOS.

## 🚀 Installation

Install the CLI globally using the Go toolchain:

```bash
go install github.com/SalvucciFacundo/go-arch-cli@latest
```

## 🛠️ Usage Guide

### 1. Environment Setup
Detects your Operating System and guides you through installing the Go toolchain and essential development tools.
```bash
go-arch setup
```

### 2. Scaffold a New Project
Launches an interactive wizard to configure specific project requirements like Module Name, Layout, and Database Drivers.
```bash
go-arch new my-project
```

### 3. Development Server
Runs the application with optimized settings. Automatically detects `Air` for hot-reload capabilities.
```bash
go-arch serve
```

### 4. Component Generation
Generates boilerplate code adhering to the project's metadata (detects layout and module namespace).
```bash
go-arch generate service Product
go-arch generate repository User
```

## 📐 Supported Architectures

- **Minimalist**: Thin structure for microservices, lambda functions, or single-file scripts.
- **Standard**: Conventional Go layout for mid-sized projects and CLI tools.
- **Hexagonal**: Domain-Centric design for enterprise-grade applications requiring high decoupling from external infrastructure.

## 📚 Documentation

For deep technical insights and usage guides, please refer to:
- [**Architecture Guide**](./docs/ARCHITECTURE.md): Detailed explanation oflayouts and internal design patterns.
- [**Command Reference**](./docs/COMMANDS.md): Comprehensive guide on how to use every CLI command with examples.

## 🤝 Contributing

Contributions are what make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---
Built with ❤️ for the Go Community.
