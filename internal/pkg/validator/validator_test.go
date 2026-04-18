package validator

import (
	"go-arch/internal/ui"
	"go/ast"
	"os"
	"testing"
)

func TestValidator_checkStructure(t *testing.T) {
	// Crear un directorio temporal para el test
	tmpDir, err := os.MkdirTemp("", "validator_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Cambiar el directorio de trabajo para el test
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	config := &ui.ProjectConfig{
		Architecture: "Hexagonal",
	}
	v := NewValidator(config)

	// Caso 1: Estructura vacía -> Debería haber 3 errores (domain, ports, adapters)
	violations := v.checkStructure()
	if len(violations) != 3 {
		t.Errorf("Esperaba 3 violaciones de estructura, obtuve %d", len(violations))
	}

	// Caso 2: Crear una carpeta -> Debería haber 2 errores
	os.MkdirAll("internal/domain", 0755)
	violations = v.checkStructure()
	if len(violations) != 2 {
		t.Errorf("Esperaba 2 violaciones de estructura, obtuve %d", len(violations))
	}
}

func TestValidator_applyArchitectureRules_Hexagonal(t *testing.T) {
	config := &ui.ProjectConfig{
		ModuleName:   "github.com/test/app",
		Architecture: "Hexagonal",
	}
	v := NewValidator(config)

	tests := []struct {
		name       string
		path       string
		importPath string
		wantError  bool
	}{
		{
			name:       "Domain importing project root (Legal)",
			path:       "internal/domain/user.go",
			importPath: "github.com/test/app",
			wantError:  false,
		},
		{
			name:       "Domain importing adapters (Illegal)",
			path:       "internal/domain/user.go",
			importPath: "github.com/test/app/internal/adapters/db",
			wantError:  true,
		},
		{
			name:       "Adapters importing domain (Legal)",
			path:       "internal/adapters/db/user_repo.go",
			importPath: "github.com/test/app/internal/domain",
			wantError:  false,
		},
		{
			name:       "Ports importing adapters (Illegal)",
			path:       "internal/ports/user_repo.go",
			importPath: "github.com/test/app/internal/adapters/db",
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulamos el objeto ast.ImportSpec de forma minimalista para el test interno
			// Nota: En el código real usamos applyArchitectureRules que recibe []*ast.ImportSpec
			// Para el test, vamos a invocar la lógica interna que extrajimos.
			
			// Como la lógica está en un método privado o semi-privado, podemos testearla 
			// pasando un dummy o ajustando la visibilidad si fuera necesario.
			// Por simplicidad en este entorno, verificamos la lógica del switch.
			
			// Re-implementación rápida de la lógica para el test de unidad puro
			violations := v.applyArchitectureRules(tt.path, createDummyImports(tt.importPath))
			if (len(violations) > 0) != tt.wantError {
				t.Errorf("applyArchitectureRules() error = %v, wantError %v", len(violations) > 0, tt.wantError)
			}
		})
	}
}

// Helper para crear imports falsos para el test
func createDummyImports(path string) []*ast.ImportSpec {
	importSpec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Value: "\"" + path + "\"",
		},
	}
	return []*ast.ImportSpec{importSpec}
}
