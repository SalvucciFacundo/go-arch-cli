package scaffold

import (
	"go-arch/internal/ui"
	"os"
	"path/filepath"
	"testing"
)

func TestScaffolder_Layouts(t *testing.T) {
	tests := []struct {
		name         string
		architecture string
		expectFiles  []string
	}{
		{
			name:         "Minimalist Architecture",
			architecture: "Minimalist",
			expectFiles:  []string{"main.go", "go.mod", ".go-arch.yaml"},
		},
		{
			name:         "Standard Architecture",
			architecture: "Standard",
			expectFiles:  []string{"cmd/api/main.go", "internal/service", "go.mod"},
		},
		{
			name:         "Hexagonal Architecture",
			architecture: "Hexagonal",
			expectFiles:  []string{"cmd/api/main.go", "internal/domain", "internal/ports", "internal/adapters", "go.mod"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "scaffold-test-*")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tempDir)

			oldWd, _ := os.Getwd()
			os.Chdir(tempDir)
			defer os.Chdir(oldWd)

			config := &ui.ProjectConfig{
				ProjectName:  "TestApp",
				ModuleName:   "github.com/test/app",
				Architecture: tt.architecture,
			}

			scaffolder := NewScaffolder(config)
			err = scaffolder.Execute()
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			for _, f := range tt.expectFiles {
				path := filepath.Join("TestApp", f)
				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Errorf("expected %s to exist in %s layout", f, tt.architecture)
				}
			}
		})
	}
}

func TestScaffolder_CRUD(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "crud-integration-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Mocking a project root
	config := &ui.ProjectConfig{
		ProjectName:  ".",
		ModuleName:   "github.com/test/crud",
		Architecture: "Hexagonal",
	}

	scaffolder := NewScaffolder(config)
	err = scaffolder.GenerateCRUD("User")
	if err != nil {
		t.Fatal(err)
	}

	expectedFiles := []string{
		"internal/domain/User.go",
		"internal/adapters/User_handler.go",
	}

	for _, f := range expectedFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected crud file %s to exist", f)
		}
	}
}
