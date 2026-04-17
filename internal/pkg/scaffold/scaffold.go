package scaffold

import (
	"fmt"
	"go-arch/internal/pkg/template"
	"go-arch/internal/ui"
	"os"
	"path/filepath"
)

type Scaffolder struct {
	engine *template.Engine
	config *ui.ProjectConfig
}

func NewScaffolder(config *ui.ProjectConfig) *Scaffolder {
	return &Scaffolder{
		engine: template.NewEngine(),
		config: config,
	}
}

func (s *Scaffolder) Execute() error {
	fmt.Printf("🏗️ Creando proyecto '%s' con arquitectura %s...\n", s.config.ProjectName, s.config.Architecture)

	// 1. Crear directorio base
	if err := os.MkdirAll(s.config.ProjectName, 0755); err != nil {
		return err
	}

	// 2. Generar estructura según el layout
	switch s.config.Architecture {
	case "Minimalist":
		return s.scaffoldMinimalist()
	case "Standard":
		return s.scaffoldStandard()
	case "Hexagonal":
		return s.scaffoldHexagonal()
	default:
		return fmt.Errorf("arquitectura no soportada: %s", s.config.Architecture)
	}
}

func (s *Scaffolder) createFile(path string, templatePath string) error {
	fullPath := filepath.Join(s.config.ProjectName, path)
	
	// Crear directorios intermedios
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return s.engine.Render(f, templatePath, s.config)
}

func (s *Scaffolder) scaffoldMinimalist() error {
	// Solo main.go y go.mod
	if err := s.createFile("main.go", "minimalist/main.tmpl"); err != nil {
		return err
	}
	return s.createCommonFiles()
}

func (s *Scaffolder) scaffoldStandard() error {
	dirs := []string{
		"cmd/api",
		"internal/handler",
		"internal/service",
		"internal/repository",
	}
	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(s.config.ProjectName, d), 0755); err != nil {
			return err
		}
	}
	
	if err := s.createFile("cmd/api/main.go", "standard/main.tmpl"); err != nil {
		return err
	}
	return s.createCommonFiles()
}

func (s *Scaffolder) scaffoldHexagonal() error {
	dirs := []string{
		"cmd/api",
		"internal/domain",
		"internal/ports",
		"internal/adapters",
	}
	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(s.config.ProjectName, d), 0755); err != nil {
			return err
		}
	}

	if err := s.createFile("cmd/api/main.go", "hexagonal/main.tmpl"); err != nil {
		return err
	}
	return s.createCommonFiles()
}

func (s *Scaffolder) createCommonFiles() error {
	// go.mod, .env.example, .go-arch.yaml
	if err := s.createFile("go.mod", "common/go.mod.tmpl"); err != nil {
		return err
	}
	if err := s.createFile(".go-arch.yaml", "common/config.tmpl"); err != nil {
		return err
	}
	return nil
}

// GenerateComponent genera un componente específico (service, repository, handler)
// en la ubicación correcta según la arquitectura.
func (s *Scaffolder) GenerateComponent(compType, name string) error {
	var targetPath string
	var templatePath string

	data := struct {
		ui.ProjectConfig
		EntityName string
	}{
		ProjectConfig: *s.config,
		EntityName:    name,
	}

	switch compType {
	case "service":
		templatePath = "common/service.tmpl"
		if s.config.Architecture == "Hexagonal" {
			targetPath = filepath.Join("internal/domain", name+"_service.go")
		} else {
			targetPath = filepath.Join("internal/service", name+"_service.go")
		}
	case "repository":
		templatePath = "common/repository.tmpl"
		if s.config.Architecture == "Hexagonal" {
			targetPath = filepath.Join("internal/ports", name+"_repository.go")
		} else {
			targetPath = filepath.Join("internal/repository", name+"_repository.go")
		}
	case "handler":
		templatePath = "common/handler.tmpl"
		if s.config.Architecture == "Hexagonal" {
			targetPath = filepath.Join("internal/adapters", name+"_handler.go")
		} else {
			targetPath = filepath.Join("internal/handler", name+"_handler.go")
		}
	default:
		return fmt.Errorf("tipo de componente no soportado: %s", compType)
	}

	fmt.Printf("🛠️  Generando %s en %s...\n", compType, targetPath)

	// En generación, el ProjectName del config es el directorio raíz (que suele ser ".")
	// o el nombre del proyecto original. Vamos a asumir que generamos en el directorio actual.
	fullPath := targetPath
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return s.engine.Render(f, templatePath, data)
}
