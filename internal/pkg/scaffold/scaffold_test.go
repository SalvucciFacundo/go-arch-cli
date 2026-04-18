package scaffold

import (
	"go-arch/internal/ui"
	"os"
	"path/filepath"
	"testing"
)

func TestScaffolder_Execute(t *testing.T) {
	// Crear un directorio temporal para el test
	tempDir, err := os.MkdirTemp("", "scaffolder-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Cambiar al directorio temporal
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	config := &ui.ProjectConfig{
		ProjectName:  "MyApp",
		ModuleName:   "github.com/test/app",
		Architecture: "Standard",
		DBDriver:     "PostgreSQL",
		UseDocker:    false,
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.Execute()

	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verificar que el directorio del proyecto se creó
	if _, err := os.Stat("MyApp"); os.IsNotExist(err) {
		t.Error("expected project directory 'MyApp' to exist")
	}

	// Verificar archivos clave
	expectedFiles := []string{
		"MyApp/go.mod",
		"MyApp/.go-arch.yaml",
		"MyApp/cmd/api/main.go",
	}

	for _, f := range expectedFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", f)
		}
	}
}

func TestScaffolder_DockerSupport(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "docker-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	config := &ui.ProjectConfig{
		ProjectName:  "DockerApp",
		ModuleName:   "github.com/test/docker",
		Architecture: "Standard",
		DBDriver:     "PostgreSQL",
		UseDocker:    true,
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.Execute()
	if err != nil {
		t.Fatal(err)
	}

	dockerFiles := []string{
		"DockerApp/Dockerfile",
		"DockerApp/docker-compose.yaml",
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
