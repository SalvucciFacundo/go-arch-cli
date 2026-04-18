package scaffold

import (
	"go-arch/internal/ui"
	"os"
	"path/filepath"
	"testing"
)

func TestScaffolder_Execute(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "scaffold-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	config := &ui.ProjectConfig{
		ProjectName:  filepath.Join(tempDir, "MyApp"),
		ModuleName:   "github.com/test/myapp",
		Architecture: "Standard",
		DBDriver:     "PostgreSQL",
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verificar archivos clave
	files := []string{
		filepath.Join(config.ProjectName, "go.mod"),
		filepath.Join(config.ProjectName, ".go-arch.yaml"),
		filepath.Join(config.ProjectName, "cmd/api/main.go"),
		filepath.Join(config.ProjectName, "internal/handler"),
	}

	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", f)
		}
	}
}

func TestScaffolder_GenerateComponent(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gen-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Cambiar al directorio temporal para simular estar dentro del proyecto
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	config := &ui.ProjectConfig{
		ProjectName:  ".",
		ModuleName:   "github.com/test/gen",
		Architecture: "Hexagonal",
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.GenerateComponent("service", "User")
	if err != nil {
		t.Fatalf("GenerateComponent failed: %v", err)
	}

	// En Hexagonal, el servicio va a internal/domain
	expectedFile := "internal/domain/User_service.go"
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Errorf("expected file %s to exist", expectedFile)
	}
}

func TestScaffolder_WithDocker(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "docker-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	config := &ui.ProjectConfig{
		ProjectName:  filepath.Join(tempDir, "DockerApp"),
		ModuleName:   "github.com/test/docker",
		Architecture: "Standard",
		DBDriver:     "PostgreSQL",
		UseDocker:    true,
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verificar archivos de Docker
	dockerFiles := []string{
		filepath.Join(config.ProjectName, "Dockerfile"),
		filepath.Join(config.ProjectName, "docker-compose.yaml"),
		filepath.Join(config.ProjectName, ".env"),
	}

	for _, f := range dockerFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected docker file %s to exist", f)
		}
	}
}

func TestScaffolder_GenerateCRUD(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "crud-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	config := &ui.ProjectConfig{
		ProjectName:  ".",
		ModuleName:   "github.com/test/crud",
		Architecture: "Hexagonal",
		DBDriver:     "PostgreSQL",
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.GenerateCRUD("Product")
	if err != nil {
		t.Fatalf("GenerateCRUD failed: %v", err)
	}

	// Verificar archivos generados en arquitectura Hexagonal
	expectedFiles := []string{
		"internal/domain/Product.go",
		"internal/domain/Product_service.go",
		"internal/ports/Product_repository.go",
		"internal/adapters/Product_repository.go",
		"internal/adapters/Product_handler.go",
	}

	for _, f := range expectedFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected crud file %s to exist", f)
		}
	}
}
