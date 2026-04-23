package validator

import (
	"fmt"
	"go-arch/internal/ui"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/oops"
)

// Violation representa una infracción a las reglas arquitectónicas.
type Violation struct {
	File     string
	Message  string
	Severity string // "ERROR" o "WARNING"
}

type Validator struct {
	config *ui.ProjectConfig
}

func NewValidator(config *ui.ProjectConfig) *Validator {
	return &Validator{config: config}
}

// Validate verifica la estructura del proyecto y las dependencias según la arquitectura configurada.
func (v *Validator) Validate() ([]Violation, error) {
	var violations []Violation

	// 1. Validar integridad de carpetas
	structureViolations := v.checkStructure()
	violations = append(violations, structureViolations...)

	// 2. Validar reglas de imports (Dependency Rule)
	dependencyViolations, err := v.checkDependencies()
	if err != nil {
		return nil, oops.
			Code("validator_io_error").
			Wrapf(err, "Error analizando dependencias del proyecto")
	}
	violations = append(violations, dependencyViolations...)

	return violations, nil
}

func (v *Validator) checkStructure() []Violation {
	var violations []Violation
	var requiredDirs []string

	switch v.config.Architecture {
	case "Hexagonal":
		requiredDirs = []string{"internal/domain", "internal/ports", "internal/adapters"}
	case "Standard":
		requiredDirs = []string{"internal/handler", "internal/service", "internal/repository", "internal/model"}
	}

	for _, dir := range requiredDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			violations = append(violations, Violation{
				File:     dir,
				Message:  fmt.Sprintf("Falta la carpeta requerida para el layout %s", v.config.Architecture),
				Severity: "ERROR",
			})
		}
	}

	return violations
}

func (v *Validator) checkDependencies() ([]Violation, error) {
	var violations []Violation

	err := filepath.Walk("internal", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return nil // Ignorar archivos con errores de sintaxis por ahora
		}

		fileViolations := v.applyArchitectureRules(path, f.Imports)
		violations = append(violations, fileViolations...)

		return nil
	})

	return violations, err
}

func (v *Validator) applyArchitectureRules(path string, imports []*ast.ImportSpec) []Violation {
	var violations []Violation
	modulePrefix := v.config.ModuleName + "/internal"

	for _, imp := range imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		
		// Solo nos interesan los imports internos del propio proyecto
		if !strings.HasPrefix(importPath, modulePrefix) {
			continue
		}

		relImport := strings.TrimPrefix(importPath, modulePrefix)

		switch v.config.Architecture {
		case "Hexagonal":
			// Regla: domain no puede importar nada de ports o adapters
			if strings.Contains(path, "internal/domain") {
				if strings.Contains(relImport, "/ports") || strings.Contains(relImport, "/adapters") {
					violations = append(violations, Violation{
						File:     path,
						Message:  fmt.Sprintf("Fuga de capas: Dominio no debe importar '%s'", importPath),
						Severity: "ERROR",
					})
				}
			}
			// Regla: ports no puede importar adapters
			if strings.Contains(path, "internal/ports") {
				if strings.Contains(relImport, "/adapters") {
					violations = append(violations, Violation{
						File:     path,
						Message:  fmt.Sprintf("Fuga de capas: Los Puertos (interfaces) no deben importar Adaptadores '%s'", importPath),
						Severity: "ERROR",
					})
				}
			}

		case "Standard":
			// Regla: model no importa nada
			if strings.Contains(path, "internal/model") {
				violations = append(violations, Violation{
					File:     path,
					Message:  fmt.Sprintf("El paquete 'model' debe ser autocontenido, no debe importar '%s'", importPath),
					Severity: "ERROR",
				})
			}
			// Regla: repository no importa service ni handler
			if strings.Contains(path, "internal/repository") {
				if strings.Contains(relImport, "/service") || strings.Contains(relImport, "/handler") {
					violations = append(violations, Violation{
						File:     path,
						Message:  fmt.Sprintf("Inversión de dependencias prohibida: El repositorio no debe depender de '%s'", importPath),
						Severity: "ERROR",
					})
				}
			}
		}
	}

	return violations
}
